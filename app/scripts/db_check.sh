#!/bin/sh
# wait-for-postgres.sh

set -e
  
host="$1"
shift
  
# until PGPASSWORD=$POSTGRES_PASSWORD psql -h "$host" -U "postgres" -c '\q'; do
until PGPASSWORD="teste" psql -h "db" -U "teste" -c '\q'; do
  >&2 echo "Aguardando inicialização do Postgres.."
  >&2 echo " "
  sleep 3
done
  
>&2 echo "Postgres iniciado - Prosseguindo inicialização do App.."
exec "$@"