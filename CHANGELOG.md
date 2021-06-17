## [1.0.23] - 16/06/2021

### Melhoria 
- Adição da rota listerunecachaca
    A nova rota recebe uma variavel ia URL no padrão "/listerunecachaca/{nome}" e retorna os dados da cachaça, caso o
    nome consulado exista entre as opções disponíveis.


## [1.0.8] - 14/06/2021

### Melhoria 
- Ajuste na biblioteca de rotas para o gorilla/mux
    Estava tentando funcionar a algum tempo e apanhando, no fim a ordem da chamada do go mod init estava errada apenas

- Ajuste no Dockerfile
    Reorganizada a copia de arquivos para dentro do container, levando apenas o main.go
    Removidos comandos não utilizados.


## [1.0.0] - 29/04/2021

### Adição 
- Criado arquivo docker-compose.yml
    O compose do projeto esta baseado na utilização de stack com swarm ativo

## [Unreleased]