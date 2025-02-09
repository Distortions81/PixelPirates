#!/bin/bash
set -e

# Variables for certificate details.
DAYS_VALID=365
CERT_FILE="cert.pem"
KEY_FILE="key.pem"
SUBJECT="/CN=localhost"

echo "Generating self-signed certificate and key..."
openssl req -x509 -newkey rsa:4096 \
    -keyout "${KEY_FILE}" -out "${CERT_FILE}" \
    -days "${DAYS_VALID}" -nodes \
    -subj "${SUBJECT}"

echo "Certificate and key generated:"
echo "  Certificate: ${CERT_FILE}"
echo "  Key:         ${KEY_FILE}"
