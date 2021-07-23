# GoDebut
 Projeto para teste de container GO rodando no sistema WSL.

#### Endpoints:
##### `"/"`
Home Page

##### `"/aproposde"`
Sobre mim

##### `"/toutescachacas"`
Lista todas as cachaças cadastradas*

##### `GET "/unecachaca/{nome}"`
Busca uma cachaça por nome

##### `POST "/unecachaca"`
Cadastrar uma nova cachaça com dados repassados

##### `DELETE "/unecachaca/{nome}"`
Remover uma cachaça da lista pelo nome

#### Ambiente Stack
O projeto deve ser executado em um ambiente com docker swarm ativo 
>**`docker swarm init`**

O compose deve ser iniciado via stack.
>**`docker stack deploy -c docker-compose.yml`**

Buscar endpoints na porta 8085 do servidor
>**`localhost:8085/`**
