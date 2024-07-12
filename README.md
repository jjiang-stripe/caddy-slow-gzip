# Instructions

```
git clone git@github.com:jjiang-stripe/caddy-slow-gzip.git
cd caddy-slow-gzip
caddy run
go run server/server.go
go run client/client.go
```

When the `minimum_length` in the `gzip` config of the `Caddyfile` is set to 600 (or when the `gzip` module is removed entirely), responses are read by the client as the server writes them.

When the `minimum_length` is set to 512, the client doesn't read the response until the server finishes writing the full response.
