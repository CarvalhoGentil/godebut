#!/bin/bash
# PROBLEMA ORIGINAL
# ls -l menu* | head -50 | if [[ (awk '{print $7}') -eq 10 ]]; then echo "nicer" ; fi

# SOLUÇÃO
ls -l deploy* | head -50 | awk '($7 == 10)'


# FUNCIONALIDADES PARA O DOCKER
docker images | grep godebut | docker rmi $(awk 'print @3')
docker images | grep godebut | docker rmi $(awk '{print $3}')

# LISTAR IMAGENS SEM TAG
docker images | awk '($1 == "<none>")'

# REMOVER TODAS AS IMAGENS SEM NOME
docker images | awk '($1 == "<none>")' | docker rmi $(awk '{print $3}')
