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
docker stack rm godebut_dev
echo $HASHLINE
echo " "

sleep 8

echo $HASHLINE
echo "Removendo imagens anteriores de Dev..."
docker images | grep godebut_dev | docker rmi $(awk {'print $3'})
echo $HASHLINE

sleep 2

echo $HASHLINE
echo "Iniciando BUILD da imagem Dev..."
echo $HASHLINE
echo " "

docker build --no-cache -t ffelixneto/godebut_dev:latest .
# docker build -t ffelixneto/godebut_dev:latest .


echo $HASHLINE
echo "Iniciando o serviço GoDebut Dev..."
echo $HASHLINE
echo " "

docker stack deploy godebut_dev -c docker-compose_dev.yml

sleep 2

echo $HASHLINE
echo "GoDebut Dev deploy finalizado !"
echo $HASHLINE
echo " "
