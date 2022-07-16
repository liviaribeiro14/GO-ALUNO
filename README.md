
# gRPC - Teste

A aplicação está rodando diretamente num container docker a partir de uma imagem do golang. 
``` bash
docker run --rm -it -v $(pwd)/:/usr/src/app -p 6060:8080 golang:1.18.4-buster bash
```

### Primeira vez no container
Ao entrar pela primeira vez no container, realizar os seguintes passos: 

Atualizar o Debian:
```bash
apt-get update
```

Instalar o protobuf-compile
```bash
apt install protobuf-compiler
```

Instalar o protoc:
```bash
go get google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

Instalar o protoc-gen-go-grpc
```bash
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

OBS: fiz tudo dentro da pasta `/usr/src/app`.

---

## Comandos diversos
E para gerar o pb - protocol buffer: `protoc --proto_path=proto proto/*.proto --go_out=pb`

E para gerar os arquivos que vão retornar as stubs, para se conectar com o servidor gRPC. 
`protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb`

Código responsável para gerar as stubs (ainda vou descobrir o que é)
``` bash
protoc --proto_path=proto/ proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=. --go_out=.
```

Para rodar o serviço:
``` bash
go run cmd/server/server.go
```