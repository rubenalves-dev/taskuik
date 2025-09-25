#!/bin/bash

echo "Testing CORS with curl..."

echo "1. Testing preflight OPTIONS request:"
curl -X OPTIONS http://localhost:8080/tasks \
  -H "Origin: http://localhost:4200" \
  -H "Access-Control-Request-Method: GET" \
  -H "Access-Control-Request-Headers: Content-Type" \
  -v

echo -e "\n\n2. Testing actual GET request:"
curl -X GET http://localhost:8080/tasks \
  -H "Origin: http://localhost:4200" \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  -v

echo -e "\n\nDone."
