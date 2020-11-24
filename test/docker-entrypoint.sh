#!/bin/sh

echo -en "\033[37;1;42m Check start page: \033[0m"

curl -H "Host: localhost:8080" --silent --show-error --fail -I http://app_nginx:8080/

echo -e "----------\n"

sleep 1

echo -en "\033[37;1;42m Check Golang health check: \033[0m"

curl -H "Host: localhost:8080" --silent --show-error --fail -I http://app_nginx:8080/api/health