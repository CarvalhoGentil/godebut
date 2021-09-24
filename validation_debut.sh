#!/bin/bash

# Script para validar o conteúdo dos arquivos de secrets
# - dbname
# - dbpass
# - dbuser

clear

HASHLINE="##############################################"

echo $HASHLINE
echo "validando existencia e conteudo de secrets..."
echo $HASHLINE
echo " "

# CAMINHO DOS ARQUIVOS

DB_NAMEFILE=$(pwd)"/mydbsecrets/dbname"
DB_USERFILE=$(pwd)"/mydbsecrets/dbuser"
DB_PASSFILE=$(pwd)"/mydbsecrets/dbpass"

if test -f "$DB_NAMEFILE" && test -s "$DB_NAMEFILE" ; then echo "Arquivo dbname - OK"; else echo "Conteúdo do $DB_NAMEFILE não encontrado" ; fi
echo " "
if test -f "$DB_USERFILE" && test -s "$DB_USERFILE" ; then echo "Arquivo dbuser - OK"; else echo "Conteúdo do $DB_USERFILE não encontrado" ; fi
echo " "
if test -f "$DB_PASSFILE" && test -s "$DB_PASSFILE" ; then echo "Arquivo dbpass - OK"; else echo "Conteúdo do $DB_PASSFILE não encontrado" ; fi
echo " "

echo $HASHLINE

