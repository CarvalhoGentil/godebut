#!/bin/bash

# Script segue os seguintes passos:
# - Remover o serviço do "GoDebut"
# - Atualizar o repsitorio local
# - Fazer o build do container 
# - Fazer o push do container para o Docker hub 
# - Startar o serviço do "GoDebut"

clear

stack_name="godebut_dev"
environment="Dev"
if [[ $1 == "prod" ]]; then
    echo "Ambiente produção ATIVARRR !!!"
    stack_name="godebut"
    environment="Produção"
else
    echo "Ambiente teste ;)"
fi

HASHLINE="##############################################"

echo $HASHLINE
echo "Removendo stack $stack_name para liberar as portas..."
docker stack rm $stack_name
echo $HASHLINE
echo " "

sleep 8

echo $HASHLINE
echo "Removendo imagens anteriores de $stack_name..."
docker images | grep $stack_name | docker rmi $(awk {'print $3'})
echo $HASHLINE

sleep 2

echo $HASHLINE
echo "Iniciando BUILD da imagem $stack_name..."
echo $HASHLINE
echo " "

# CACHE
docker build -t ffelixneto/$stack_name:latest .
# NO CACHE
# docker build --no-cache -t ffelixneto/$stack_name:latest .
sleep 1

echo $HASHLINE
echo "Iniciando o serviço $stack_name..."
echo $HASHLINE
echo " "

docker stack deploy $stack_name -c docker-compose_dev.yml

sleep 4

echo $HASHLINE
echo "GoDebut $environment deploy finalizado !"
echo $HASHLINE
echo " "
