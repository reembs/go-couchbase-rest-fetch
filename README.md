# go-couchbase-rest-fetch
HTTP Multi-fetch facade for Couchbase. Useful for services that do reads but don't want to get all mixed up with the SDK code. Written in Go. I use the binary directly to connect to Couchbase instances for debugging purposes since Couchbase has no HTTP key/value endpoints of its own.

HTTP server has two endpoints:

* `/get/<key>`
  * Fetch a single key.
* `/mget/<keys>`
  * Fetch keys in parallel. List is comma separated.


```
Usage of cb-rest-fetch:
  -bucket string
        Bucket to connect to (required)
  -host string
        Couchbase connection string (required)
  -password string
        Bucket password
  -port int
        Port to listen on (default 8080)
```

Deploy using docker on the Couchbase server localhost: `docker run -p 8080:8080 --network=host -ti reembs/go-couchbase-rest-fetch -host couchbase://127.0.0.1 -password <bucketpass> -bucket <bucket>`
