#!/bin/bash

PWD=`pwd`
WORKSPACE="/tmp/ssl"
SUBJECT="/C=CN/ST=Si Chuan/L=Cheng Du/O=Yao Guai Ltd/OU=Yao Guai/CN=jegarn.com"
PASSWORD="1111"
PASSWD="$WORKSPACE/passwd"

workspace_init ()
{
    read -p "Enter your save directory [$WORKSPACE]: " WORKSPACE
    read -p "Enter your subject [$SUBJECT]: " SUBJECT
    PASSWD="$WORKSPACE/passwd"

    if [ -d $WORKSPACE ]; then
        printf "${WORKSPACE} already exists, remove it?  yes/no: "
        read del
        if [ $del = "yes" ]; then
            rm -rf $WORKSPACE
        else
            exit 0
        fi
    fi

    for d in ${WORKSPACE} "${WORKSPACE}/root" "${WORKSPACE}/server" "${WORKSPACE}/client" "${WORKSPACE}/certs"
    do
        if [ ! -d $d ]; then
            mkdir $d
        fi
    done

    chmod 770 $WORKSPACE
    chown -R root:root $WORKSPACE

    echo 01 > "${WORKSPACE}/serial"
    index_file="${WORKSPACE}/index.txt"
    touch $index_file
    touch $PASSWD
    chmod 600 $PASSWD
    chown root:root $PASSWD
    echo "$(date)" >> $PASSWD
}

create_openssl_config ()
{
    openssl_cnf="${WORKSPACE}/openssl.cnf"
    touch $openssl_cnf
    echo "[ ca ]
    default_ca = yaoguais_ca

    [ yaoguais_ca ]
    certificate = ./ca.crt
    private_key = ./ca.key
    database = ./index.txt
    serial = ./serial
    new_certs_dir = ./certs

    default_days = 3650
    default_md = sha1

    policy = yaoguais_ca_policy
    x509_extensions = yaoguais_ca_extensions

    [ yaoguais_ca_policy ]
    commonName = supplied
    stateOrProvinceName = optional
    countryName = optional
    emailAddress = optional
    organizationName = optional
    organizationalUnitName = optional

    [ yaoguais_ca_extensions ]
    basicConstraints = CA:false


    [ req ]
    default_bits = 2048
    default_keyfile = ./ca.key
    default_md = sha1
    prompt = yes
    distinguished_name = root_ca_distinguished_name
    x509_extensions = root_ca_extensions

    [ root_ca_distinguished_name ]
    countryName_default = CN

    [ root_ca_extensions ]
    basicConstraints = CA:true
    keyUsage = cRLSign, keyCertSign

    [ server_ca_extensions ]
    basicConstraints = CA:false
    keyUsage = keyEncipherment

    [ client_ca_extensions ]
    basicConstraints = CA:false
    keyUsage = digitalSignature" > $openssl_cnf
}

create_ca_keys ()
{
    cd $WORKSPACE
    read -p "Enter your ca.key password [1111]: " PASSWORD
    echo "ca.key password: $PASSWORD" >> $PASSWD
    openssl genrsa -des3 -out root/ca.key 2048
    openssl req -new -subj "$SUBJECT" -newkey rsa:2048 -key root/ca.key -out root/ca.csr
    openssl x509 -req -days 3650 -in root/ca.csr -signkey root/ca.key -out root/ca.crt
}

create_server_keys ()
{
    read -p "Enter your server.key password [1111]: " PASSWORD
    echo "server.key password: $PASSWORD" >> $PASSWD
    openssl genrsa -des3 -out server/server.key 2048
    openssl req -new -subj "$SUBJECT" -newkey rsa:2048 -key server/server.key -out server/server.csr
    mv server/server.key server/server.ori.key
    openssl rsa -in server/server.ori.key -out server/server.key
    openssl ca -config openssl.cnf -in server/server.csr -cert root/ca.crt -keyfile root/ca.key -out server/server.crt -days 3650
}

create_client_keys ()
{
    read -p "Enter your client.key password [1111]: " PASSWORD
    echo "client.key password: $PASSWORD" >> $PASSWD
    openssl genrsa -des3 -out client/client.key 2048
    openssl req -new -subj "$SUBJECT" -newkey rsa:2048 -key client/client.key -out client/client.csr
    echo "unique_subject = no" > "index.txt.attr"
    mv client/client.key client/client.ori.key
    openssl rsa -in client/client.ori.key -out client/client.key
    openssl ca -config openssl.cnf -in client/client.csr -cert root/ca.crt -keyfile root/ca.key -out client/client.crt -days 3650
    read -p "Enter your client.p12 password [1111]: " PASSWORD
    echo "client.p12 password: $PASSWORD" >> $PASSWD
    openssl pkcs12 -export -in client/client.crt -inkey client/client.key -out client/client.p12 -certfile root/ca.crt

}

workspace_init;
create_openssl_config;
create_ca_keys;
create_server_keys;
create_client_keys;
cd $PWD
