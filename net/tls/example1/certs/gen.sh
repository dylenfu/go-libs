#!/bin/bash
# call this script with an email address (valid or not).
# ./gen.sh dylenfu@random.com

rm -rf *.crt *.key *.csr

# 制作证书的过程
#步骤(1) ，生成CA自己的私钥 rootCA.key
#步骤(2)，根据CA自己的私钥生成自签发的数字证书，该证书里包含CA自己的公钥。
#
#步骤(3)~(5)是用来生成服务端的私钥和数字证书（由自CA签发）。
#步骤(3)，生成服务端私钥。
#步骤(4)，生成Certificate Sign Request，CSR，证书签名请求。
#步骤(5)，自CA用自己的CA私钥对服务端提交的csr进行签名处理，得到服务端的数字证书device.crt。
#
#步骤(6)，将自CA的数字证书同客户端一并发布，用于客户端对服务端的数字证书进行校验。
#步骤(7)和步骤(8)，将服务端的数字证书和私钥同服务端一并发布。

openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -subj "/CN=tonybai.com" -days 5000 -out ca.crt

openssl genrsa -out server.key 2048
openssl req -new -key server.key -subj "/CN=localhost" -out server.csr

openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000
