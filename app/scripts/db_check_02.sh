echo "Validando conexão ao banco de dados..."
sleep 1
while ! pg_isready -h db -p 5432 > /dev/null 2> /dev/null; do
    echo "Aguardando conexão ao servico Postgres"
    sleep 1
done