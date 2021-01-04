
## Twórz i zarządzaj woluminami
docker volume create vRedis
docker volume ls
docker volume inspect vRedis
docker volume rm vRedis

## Obraz redis:6.2-rc-alpine
docker pull redis:6.2-rc-alpine


---------------------------------------------
---------------------------------------------
## Uruchom kontener z arango 
docker volume create vRedis
docker run --name=myRedis -p 6379:6379 redis-server --requirepass yourpassword -d -v vRedis:/data redis:6.2-rc-alpine

sudo docker run -p 6379:6379 -v vRedis:/data --name=myRedis -d redis:6.2-rc-alpine redis-server --appendonly yes  --requirepass yourpassword


## Wyczyść kontenery i woluminy
docker container stop myRedis && docker container rm myRedis 
docker ps -a
docker images

## Przywracanie kontenera redisa z woluminu
docker run --name=myRedis -p 6379:6379 -v vRedis:/data -d redis:6.2-rc-alpine redis-server --appendonly yes  --requirepass yourpassword



---------------------------------------------
---------------------------------------------
## Usuń wszystkie nieużywane kontenery. Używaj ostrożnie
docker container prune
docker container ls -a --filter status=exited --filter status=created


## Usuń wszystkie nieużywane obrazy
docker image prune -a
docker image prune -a --filter "until=12h"


## Usuń wszystkie woluminy, wszystkie nieużywane woluminy lub jeden. Używaj ostrożnie
docker volume prune
Use the -f or --force option to bypass the prompt.

docker volume rm --force `docker volume ls`
docker volume rm vRedis


## Usuń wszystkie sieci
docker network prune
docker network prune -a --filter "until=12h"
