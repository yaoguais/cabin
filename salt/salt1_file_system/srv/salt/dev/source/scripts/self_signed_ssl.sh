#/bin/bash

pwd=`pwd`

read -p "Enter your domain [jegarn.com]: " DOMAIN
read -p "Enter your file prefix [server]: " NAME
read -p "Enter your save directory [/tmp/ssl]: " DIR

mkdir -p $DIR
cd $DIR
openssl genrsa -des3 -out $NAME.key 2048
SUBJECT="/C=CN/ST=Si Chuan/L=ChengDu/O=YaoGuai/OU=YaoGuai/CN=$DOMAIN"
openssl req -new -subj "$SUBJECT" -key $NAME.key -out $NAME.csr
# remove password
mv $NAME.key $NAME.ori.key
openssl rsa -in $NAME.ori.key -out $NAME.key
openssl x509 -req -days 3650 -in $NAME.csr -signkey $NAME.key -out $NAME.crt
cd $pwd
