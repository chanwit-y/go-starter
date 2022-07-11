[Custom JSON Encoder/Decoder](https://docs.gofiber.io/guide/faster-fiber)

[go json](https://github.com/goccy/go-json)

[mockery](https://github.com/vektra/mockery)


[gomodifytags](https://github.com/fatih/gomodifytags)

```
gomodifytags -file demo.go -struct Server -add-tags json
```

[gotests](https://github.com/cweill/gotests)

[swag](https://github.com/swaggo/swag)

```
swag init
```

add below code to route

```
app.Get("/docs/*", swagger.HandlerDefault)
```