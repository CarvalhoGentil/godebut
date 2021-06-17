# GoDebut
 Projeto para teste de container GO rodando no sistema WSL.

#### Endpoints:
- "/"
- "/aproposde"
- "/listertoutescachacas"
- "/listerunecachaca/{nome}"

#### Ambiente Stack
O projeto deve ser executado em um ambiente com docker swarm ativo 
>**`docker swarm init`**

O compose deve ser iniciado via stack.
>**`docker stack deploy -c docker-compose.yml`**

Buscar endpoints na porta 8085 do servidor
>**`localhost:8085/`**
