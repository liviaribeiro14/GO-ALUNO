
# gRPC - Close Encounters of The Third Kind

gRPC com Golang - duas tecnologias que ainda não domino (irei dominar um dia? Deixemos em aberto.)

IMHO - Golang não é para iniciantes como yo. Ela por si só já tem um monte de poréns (instalei e às vezes funciona, às vezes não - my fault, yey!). Daí junta com gRPC e o bagulho fica doido. 

Fui anotando aqui tudo o que tive que fazer seguindo o vídeo tuturial da FullCycle e o que fiz por fora, já que não funfou de primeira. Pode ser a versão das ferramentas que uso, o que mostra que essa integração deve funcionar muito bem com os engenheiros da Google (e por isso que são de lá), mas não é muito amigável pra quem só olha admirada.

Fiz a aplicação rodando diretamente num container docker a partir de uma imagem do golang, afinal na minha máquina o Go se recusa a ser consistente. 
``` bash
docker run --rm -it -v $(pwd)/:/usr/src/app -p 6060:8080 golang:1.18.4-buster bash
```

### Primeira vez no container - Instalação
Ao entrar pela primeira vez no container, realizar os seguintes passos: (TODO: fazer um Dockerfile)

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

OBS: fiz tudo dentro da pasta `/usr/src/app`. Uma coisa que me incomoda com meu pouco traquejo com o Linux é que o local importa. Mas todas às vezes? Pode ser que sim, pode ser que não.

---

## Comandos diversos
E para gerar o pb - protocol buffer: `protoc --proto_path=proto proto/*.proto --go_out=pb`

E para gerar os arquivos que vão retornar as stubs, para se conectar com o servidor gRPC. 
`protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb`

Código responsável para gerar as stubs (ainda vou descobrir o que é)
``` bash
protoc --proto_path=proto/ proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc --go-grpc_out=. --go_out=.
```

Para rodar o server:
``` bash
go run cmd/server/server.go
```

Para rodar o client:
``` bash
go run cmd/client/client.go
```

SPOILER, GALERA: Não funfou!