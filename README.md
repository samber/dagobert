
# DaGoBert

Simple Go client for the awesome [clip-as-service](https://clip-as-service.jina.ai/) server.

## Why this name ?

🤴 France had a king named Dagobert.

## Gettings start

```sh
cd example
docker-compose up -d
```

```go
client, err := dagobert.NewClient("grpc://localhost:51000")
if err != nil {
    log.Fatal(err)
}

docs, err := client.Encode(
    []*dagobert.Document{
        dagobert.NewTextDocument("hello world"),
        dagobert.NewTextDocument("hi there"),
    },
)
if err != nil {
    log.Fatal(err)
}

for _, doc := range docs {
    fmt.Println(doc.GetEmbedding().GetDense().GetBuffer())
}
```

## Todo

- [x] Encoding
- [ ] Ranking
- [ ] Indexing
- [ ] Searching
- [ ] Profiling
