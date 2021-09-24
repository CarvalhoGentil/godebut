#!/bin/bash

# Script segue os seguintes passos:
# - Remover o serviço do "GoDebut"
# - Atualizar o repsitorio local
# - Fazer o build do container 
# - Fazer o push do container para o Docker hub 
# - Startar o serviço do "GoDebut"

clear

HASHLINE="##############################################"

echo $HASHLINE
echo "Removendo stack godebut em execução..."
docker stack rm godebut
echo $HASHLINE
echo " "

sleep 6

echo $HASHLINE
echo "Iniciando PULL do repositorio..."
echo $HASHLINE
echo " "

git pull

echo $HASHLINE
echo "Iniciando BUILD da imagem..."
echo $HASHLINE
echo " "

docker build --no-cache -t ffelixneto/godebut:latest .

echo $HASHLINE
echo "Fazendo Push da imagem para Docker Hub..."
echo $HASHLINE
echo " "

docker push ffelixneto/godebut:latest

sleep 2

echo $HASHLINE
echo "Iniciando o serviço GoDebut..."
echo $HASHLINE
echo " "

docker stack deploy godebut -c docker-compose.yml

sleep 2

echo $HASHLINE
echo "GoDebut deploy finalizado !"
echo $HASHLINE
echo " "
