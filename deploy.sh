#! /bin/bash
docker-compose stop && docker-compose rm
git pull

cd ./AssassinGo-Front-End/
npm install && npm run build
cp dist/index.html ./web/templates/ && cp dist/static ../web/ -r

cd ..
sed '/./{s/^/{{define "index"}}&/;s/$/&{{end}}/}' -i web/templates/index.html
docker-compose up --build -d
