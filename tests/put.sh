#!/bin/bash

curl -k -X PUT -H "Content-Type: application/json" -H "Accept: application/json" -d '{"Name": "pão", "Price": 1.5, "Quantity": 1}' https://go.docker.localhost/products/1