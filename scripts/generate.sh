#!/bin/sh
# This file can only be called from the project root!
rm -rf ./swagger
mkdir ./swagger
curl -o ./swagger/specification.yaml https://docs.microsoft.com/en-us/dynamics-nav/api-reference/v1.0/contracts/bcoas1.0.yaml
python3 ./scripts/fix_spec.py
docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli-v3:unstable generate -i /local/swagger/specification.yaml -l go -o /local/swagger
