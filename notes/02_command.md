#### code format

```shell
go fmt 
```

```shell

gofmt -w main.go

gofmt -w -l .
```

#### run or build

If you want to quickly test a small program or code snippet, you can use the go run command. This will compile and run the program in a single step, without the need to create an executable file. You can use the go build command to build an executable file that you can deploy or distribute across different platforms.

```shell

go run main.go

```shell

go build main.go

```

#### add linter

```shell

go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

golangci-lint run

```

#### run test

```shell
go test -v ./...
```


#### ApacheBench
```shell
ab -n 100 -c 100 -T 'application/json' -p payload.json http://localhost:8000/api/endpoint

ab -n 100 -c 100 -H "Authorization: Bearer your_token" http://localhost:8000/api/endpoint
```

- `-n 100`: Specifies the number of requests to send. Adjust this value according to your needs.
- `-c 100`: Sets the concurrency level, indicating how many requests to send concurrently. Adjust this value as required.
- `-T 'application/json'`: Specifies the content type of the request payload as JSON.
- `-p payload.json`: Specifies the path to a file containing the JSON payload. Create a file named `payload.json` and put the JSON data in it.
- `http://localhost:8000/api/endpoint`: Replace this with the actual URL of your local API endpoint.


#### wrk
```
wrk -t12 -c400 -d30s -T30s --latency http://localhost:3000
# -t12 用 12 個線程
# -c400 模擬 400 個併發連接
# -d30s 持續 30 秒
# -T30s 設定超過 30 秒就算連接超時
# --latency 響應時間的分佈情況
```
