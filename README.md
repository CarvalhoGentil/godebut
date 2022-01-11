# GoDebut
 Projeto para teste de container GO rodando no sistema WSL.

![Pipeline Status](https://github.com/ffelixneto/godebut/actions/workflows/build-push-docker.yml/badge.svg) 

### Endpoints disponíveis:
##### `"/"`
Home Page

##### `"/aproposde"`
Sobre mim

##### `GET "/v1/toutescachacas"`
Lista todas as cachaças cadastradas*

##### `GET "/v1/unecachaca/{nome}"`
Busca uma cachaça por nome

##### `POST "/v1/unecachaca"`
Cadastrar uma nova cachaça com dados repassados

##### `DELETE "/v1/unecachaca/{nome}"`
Remover uma cachaça da lista pelo nome

##### `PUT "/v1/unecachaca/{id}"`
Atualizar os dados de uma cachaça da lista pelo id

#### Ambiente Stack
A aplicação utiliza as variáveis de ambiente do sistema para montar a conexão ao banco.
Devem ser criadas essas variáveis para o correto funcionamento dos serviços da aplicação:
**DEV_DB_HOST**
**DEV_DB_USER**
**DEV_DB_PASSWD**
**DEV_DB_NAME**
**DEV_DB_PORT**

O projeto deve ser executado em um ambiente com docker swarm ativo 
>**`docker swarm init`**

O compose deve ser iniciado via stack.
>**`docker stack deploy -c docker-compose.yml`**

Buscar endpoints na porta 8085 do servidor
>**`localhost:8085/`**

