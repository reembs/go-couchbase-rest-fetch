FROM golang:1.7-alpine

COPY cb-rest-fetch.go /app/

RUN apk --no-cache add git && \
    go get github.com/ant0ine/go-json-rest/rest  && \
    go get github.com/couchbase/gocb && \
    cd /app && go build ./cb-rest-fetch.go && \
    rm -rf $GOPATH/ant0ine/go-json-rest/rest $GOPATH/github.com/couchbase/gocb && \
    apk del git

ENTRYPOINT ["/app/cb-rest-fetch"]