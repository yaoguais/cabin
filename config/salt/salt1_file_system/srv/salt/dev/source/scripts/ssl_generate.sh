#!/bin/bash
# from http://blog.csdn.net/linvo/article/details/9173511
# from http://www.cnblogs.com/guogangj/p/4118605.html

dir=/tmp/ssl
workspace=`pwd`

if [ -d $dir ]; then
    printf "${dir} already exists, remove it?  yes/no: "
    read del
    if [ $del = "yes" ]; then
        rm -rf $dir
    else
        echo "user cancel"
    fi
fi

for d in ${dir} "${dir}/root" "${dir}/server" "${dir}/client" "${dir}/certs"
do
    if [ ! -d $d ]; then
        mkdir $d
    fi
done

echo 'hello world!' >> "${dir}/index.php"
echo 01 > "${dir}/serial"

index_file="${dir}/index.txt"
rm -f $index_file
touch $index_file

echo "generate openssl.cnf: "
openssl_cnf="${dir}/openssl.cnf"
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

cd $dir
echo "generate root ca: "
# in this step, I always input a password of "ca.key111111"
openssl genrsa -des3 -out root/ca.key 2048
# in this step, I always input these and a password of "ca.csr111111"
#Country Name (2 letter code) [XX]:CN
#State or Province Name (full name) []:Si Chuan
#Locality Name (eg, city) [Default City]:Cheng Du
#Organization Name (eg, company) [Default Company Ltd]:Yao Guai Ltd
#Organizational Unit Name (eg, section) []:Yao Guai
#Common Name (eg, your name or your server's hostname) []:yaoguai.com
#Email Address []:newtopstdio@163.com
#A challenge password []:ca.csr111111
#An optional company name []:Yao Guai Ltd
openssl req -new -newkey rsa:2048 -key root/ca.key -out root/ca.csr
openssl x509 -req -days 3650 -in root/ca.csr -signkey root/ca.key -out root/ca.crt

echo "generate server keys: "
# in this step, I always input a password of "server.key111111"
openssl genrsa -des3 -out server/server.key 2048
# in this step, I always input these and a password of none
#Country Name (2 letter code) [XX]:CN
#State or Province Name (full name) []:Si Chuan
#Locality Name (eg, city) [Default City]:Cheng Du
#Organization Name (eg, company) [Default Company Ltd]:Yao Guai Ltd
#Organizational Unit Name (eg, section) []:Yao Guai
#Common Name (eg, your name or your server's hostname) []:yaoguai.com
#Email Address []:newtopstdio@163.com
#A challenge password []:none
#An optional company name []:none
openssl req -new -newkey rsa:2048 -key server/server.key -out server/server.csr
openssl ca -config openssl.cnf -in server/server.csr -cert root/ca.crt -keyfile root/ca.key -out server/server.crt -days 3650

echo "generate client keys: "
# in this step, I always input a password of "client.key111111"
openssl genrsa -des3 -out client/client.key 2048
# in this step, I always input these and a password of none
#Country Name (2 letter code) [XX]:CN
#State or Province Name (full name) []:Si Chuan
#Locality Name (eg, city) [Default City]:Cheng Du
#Organization Name (eg, company) [Default Company Ltd]:Yao Guai Ltd
#Organizational Unit Name (eg, section) []:Yao Guai
#Common Name (eg, your name or your server's hostname) []:yaoguai.com
#Email Address []:newtopstdio@163.com
#A challenge password []:none
#An optional company name []:none
openssl req -new -newkey rsa:2048 -key client/client.key -out client/client.csr
# to prevent error "openssl TXT_DB error number 2 failed to update database"
echo "unique_subject = no" > "index.txt.attr"
openssl ca -config openssl.cnf -in client/client.csr -cert root/ca.crt -keyfile root/ca.key -out client/client.crt -days 3650

# use these to config nginx
: <<EOF
    ssl on;
    ssl_verify_client on;
    ssl_certificate /tmp/ssl/server/server.crt;
    ssl_certificate_key /tmp/ssl/server/server.key;
    ssl_client_certificate /tmp/ssl/root/ca.crt;
    ssl_session_timeout 5m;
    ssl_protocols SSLv3 TLSv1 TLSv1.1 TLSv1.2;
    ssl_ciphers "HIGH:!aNULL:!MD5 or HIGH:!aNULL:!MD5:!3DES";
    ssl_prefer_server_ciphers on;
EOF

echo "helps:"
# 查看key文件签名信息
echo "openssl rsa -in xxx.key -text -noout"
# 查看csr文件签名
echo "openssl req -noout -text -in xxx.csr"
# 将pem格式转换成der格式
echo "openssl x509 -in server/server.crt -outform DER -out server/server.cer"
# 使用curl请求服务器
echo 'curl -k --cert client/client.crt --key client/client.key --pass client.key111111 https://devel/index.php'
# 生成p12文件, password export111111
echo "openssl pkcs12 -export -in client/client.crt -inkey client/client.key -out client/client.p12 -certfile root/ca.crt"

cd $workspace