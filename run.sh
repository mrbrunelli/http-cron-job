#!/bin/sh

echo -e "---> Instalando pacotes..."

go get ./...

echo -e "---> Gerando build..."

go build main.go

chmod +x main*

echo -e "---> Build gerado com sucesso!"

./main*