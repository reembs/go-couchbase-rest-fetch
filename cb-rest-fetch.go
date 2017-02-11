package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"github.com/couchbase/gocb"
	"flag"
	"os"
	"strconv"
	"strings"
)

func GetFuture(bucket gocb.Bucket, key string, out chan <- map[string]interface{}) {
	go func() {
		var wrappedValue = make(map[string]interface{})

		var valueOut interface{}
		cas, err := bucket.Get(key, &valueOut)

		wrappedValue["key"] = key

		if err != nil {
			wrappedValue["error"] = err.Error()
		} else {
			wrappedValue["cas"] = cas
			wrappedValue["value"] = &valueOut
		}

		out <- wrappedValue
	}()
}

func main() {
	connectionStringPtr := flag.String("host", "", "Couchbase connection string (required)")
	buckerNamePtr := flag.String("bucket", "", "Bucket to connect to (required)")
	buckerPasswordPtr := flag.String("password", "", "Bucket password")
	portPtr := flag.Int("port", 8080, "Port to listen on")

	flag.Parse()

	if *connectionStringPtr == "" || *buckerNamePtr == "" {
		flag.Usage()
		os.Exit(2)
	}

	cluster, err := gocb.Connect(*connectionStringPtr)
	log.Printf("Created couchbase cluster object for '%s'\n", *connectionStringPtr)

	var bucket *gocb.Bucket
	bucket, err = cluster.OpenBucket(*buckerNamePtr, *buckerPasswordPtr)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
		os.Exit(2)
	}
	log.Printf("Opened bucket '%s' successfully\n", *buckerNamePtr)

	api := rest.NewApi()
	api.Use(rest.DefaultProdStack...)
	router, err := rest.MakeRouter(
		rest.Get("/get/#key", func(w rest.ResponseWriter, req *rest.Request) {
			result := make(chan map[string]interface{})
			GetFuture(*bucket, req.PathParams["key"], result)
			w.WriteJson(<- result)
		}),
		rest.Get("/mget/#keys", func(w rest.ResponseWriter, req *rest.Request) {
			keySlice := strings.Split(req.PathParams["keys"], ",")
			results := make(chan map[string]interface{}, len(keySlice))
			for _, key := range keySlice {
				GetFuture(*bucket, key, results)
			}

			resultDocs := make([]map[string]interface{}, len(keySlice))
			for i:=0 ; i<len(keySlice) ; i++  {
				resultDocs[i] = <-results
			}

			w.WriteJson(resultDocs)
		}),
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)

	log.Printf("Starting to listen on port: %d\n", *portPtr)

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(*portPtr), api.MakeHandler()))
}
