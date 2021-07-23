## [1.0.25] - 22/07/2021

### Adição
- Adiçionada função effacercachaca
    A nova rota recebe via DELETE o nome de uma cachaca para ser removida da listagem
    >curl -X DELETE localhost:8085/unecachaca/nome

### Atualização
- Unificados os nomes de endpoints de manipulação do elemento cachaca na lista.
    *creerunecachaca*, *listerunecachaca* respondem agora no endpoint *unecachaca*
    caso recebam seus respectivos parametros. 
    - **POST** com dados a adicionar 
    - ***unecachaca/'nome'*** para consulta

    <br>
    A função effacercachaca já foi criado rspondendo no endpoint *unecachaca*


## [1.0.24] - 14/07/2021

### Adição
- Adiçionada função creerunecachaca
    A nova rota recebe via POST os dados para uma nova cachaca
    >curl -d '{"nome": "Nome","id": "123", "custo": "123", "volume": "1000ml", "custo": "10"} localhost:8085/creerunecachaca

    <br>
    Os valores de `nome` e `id` são obrigatórios.
    O novo elemento é adicionado ao armazenamento em memoria apenas.
    Obs.: *Changelog retroativo*


## [1.0.23] - 16/06/2021

### Melhoria
- Adição da função listerunecachaca
    A nova rota recebe uma variavel via URL no padrão "/listerunecachaca/{nome}" e retorna os dados da cachaça, caso o
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