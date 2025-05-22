#!/bin/bash

# Variaveis
export SONAR_HOST_URL="http://localhost:9000"
export PROJECT_BASEDIR="./internal"

# Iniciando Script
echo "Executando Sonar Scanner..."
echo "------------------------"

# Obtendo valores referente ao repositório

docker run \
--rm \
--network=host \
-e SONAR_HOST_URL="$SONAR_HOST_URL"  \
-v "$PROJECT_BASEDIR:/usr/src" \
sonarsource/sonar-scanner-cli \
-D"sonar.projectKey=ci_cd" -D"sonar.sources=." -D"sonar.host.url=http://localhost:9000" -D"sonar.login=sqp_ae2074977ad00ce17ac8973343cfccc3491eb480"
