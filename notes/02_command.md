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