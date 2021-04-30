# GoDebut
 Projeto para teste de container GO rodando no sistema WSL.

#### Endpoints:
- "/"
- "/aproposde"
- "/listertoutescachacas"

#### Ambiente Stack
O projeto deve ser executado em um ambiente com docker swarm ativo 
>**`docker swarm init`**

O compose deve ser iniciado va stack.
>**`docker stack deploy -c docker-compose.yml`**

Buscar endpoints na maquina
>**`localhost:8085`**
