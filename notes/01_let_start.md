#### go mod init

```
go mod init go1
```

```shell
usage: go mod init [module-path]

Init initializes and writes a new go.mod file in the current directory, in
effect creating a new module rooted at the current directory. The go.mod file
must not already exist.

Init accepts one optional argument, the module path for the new module. If the
module path argument is omitted, init will attempt to infer the module path
using import comments in .go files, vendoring tool configuration files (like
Gopkg.lock), and the current directory (if in GOPATH).

If a configuration file for a vendoring tool is present, init will attempt to
import module requirements from it.

See https://golang.org/ref/mod#go-mod-init for more about 'go mod init'.

```

#### add gin

```shell
go get github.com/gin-gonic/gin

```

#### add logger

```shell
go get -u github.com/rs/zerolog/log

```

#### add env

```shell
go get github.com/joho/godotenv/cmd/godotenv
```

#### cli - cobra

https://github.com/spf13/cobra

```shell
go get -u github.com/spf13/cobra@latest

import "github.com/spf13/cobra"

```

```shell
go install github.com/spf13/cobra-cli@latest
```

#### grpc

###### install

https://grpc.io/docs/protoc-installation/

```
brew install protobuf
protoc --version
```

plugins
```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

```
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
hello/hello.proto
```

