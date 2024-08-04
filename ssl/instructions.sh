#!/bin/bash
# Inspired from: https://github.com/grpc/grpc-java/tree/master/examples#generating-self-signed-certificates-for-use-with-grpc

# Output files
# ca.key: Certificate Authority private key file (this shouldn't be shared in real-life)
# ca.crt: Certificate Authority trust certificate (this should be shared with users in real-life)
# server.key: Server private key, password protected (this shouldn't be shared)
# server.csr: Server certificate signing request (this should be shared with the CA owner)
# server.crt: Server certificate signed by the CA (this would be sent back by the CA owner) - keep on server
# server.pem: Conversion of server.key into a format grpc likes (this shouldn't be shared)

# Summary
# Private files: ca.key, server.key, server.pem, server.crt
# "Share" files: ca.crt (needed by the client), server.csr (needed by the CA)

SERVER_CN=localhost
PASSWORD="r6stUt6&6IZln+1e6Rupraj78iph6J7phOfUfajls2jlpr_NlwreWE6e_iv6spiH"

echo "Generating CA private key (ca.key)..."
openssl genrsa -aes256 -passout pass:$PASSWORD -out ca.key 4096

echo "Generating CA certificate (ca.crt)..."
openssl req -new -x509 -days 3650 -key ca.key -out ca.crt -subj "/CN=$SERVER_CN" -passin pass:$PASSWORD

echo "Generating server private key (server.key)..."
openssl genrsa -aes256 -passout pass:$PASSWORD -out server.key 4096

echo "Generating server certificate signing request (server.csr)..."
openssl req -new -key server.key -out server.csr -subj "/CN=$SERVER_CN" -passin pass:$PASSWORD

echo "Signing server certificate (server.crt)..."
openssl x509 -req -days 3650 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -passin pass:$PASSWORD

echo "Converting server key to .pem format (server.pem)..."
openssl pkcs8 -topk8 -nocrypt -in server.key -out server.pem -passin pass:$PASSWORD

echo "Certificate generation completed successfully."