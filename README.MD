
Subir o serviço docker, caso não esteja iniciado

service docker start

Primeira vez, fazer um build

docker-compose up --build 

Outras vezes

docker-compose up -d




docker container run -d -p 27017:27017 name auctionDB mongo