# First
openssl genpkey -algorithm RSA -out private-key.key
# Second
openssl req -new -key private-key.key -out certificate.csr
# Third
openssl x509 -req -in certificate.csr -signkey private-key.key -out certificate.pem