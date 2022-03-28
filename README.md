# Algoritmos golang


```
❯ mkdir -vp algoritmos-golang/{calcular,iterar}   

❯ cd algoritmos-golang 

❯ go mod init github.com/treinalinux/algoritmo-golang
```

## Test


```
go test
```

```
go test ./...
```

```
go test ./... -v
```

```
go test ./... -v --cover
```

```
go test ./... -v --coverprofile cobertura.txt

go tool cover --func=cobertura.txt

go tool cover --html=cobertura.txt

firefox /tmp/cover2956667008/coverage.html
```
