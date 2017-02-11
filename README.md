# go-couchbase-rest-fetch
Multi-fetch HTTP gateway for Couchbase. Useful for services that do reads but don't want to get all mixed up with the SDK code. Written in Go.

HTTP server has two endpoints:

* `/get/<KEY>`
  * Fetch a single key.
* `/mget/<KEYS>`
  * Fetch keys in parallel. List is comma delimited.


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