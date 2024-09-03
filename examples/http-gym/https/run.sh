#!/bin/bash

openssl req \
-new \
-newkey rsa:2048 \
-nodes \
-x509 \
-days 365 \
-subj \
"/C=TW/ST=skipped/L=Taipei/O=my-org/OU=my-org-unit/CN=my-app.com" \
-keyout server.key \
-out server.crt
