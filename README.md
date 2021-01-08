## docker

## Twórz i zarządzaj woluminami
docker volume create vArango
docker volume ls
docker volume inspect vArango
docker volume rm vArango

## Obraz arangodb:latest
docker pull arangodb:3.7.3


---------------------------------------------
---------------------------------------------
## Uruchom kontener z arango 
docker volume create vArango
  -- ROOT_PASSWORD - for production
docker run --name=myArango -e ARANGO_ROOT_PASSWORD=OpenSezame -p 8529:8529 -d -v vArango:/var/lib/arangodb3 arangodb:latest
  -- no auth - only for testing
docker run --name=myArango -e ARANGO_NO_AUTH=1 -p 8529:8529 -d -v vArango:/var/lib/arangodb3 arangodb:3.7.3

## Wyczyść kontenery i woluminy
docker container stop myArango && docker container rm myArango 
docker ps -a
docker images

## Przywracanie kontenera z arango z woluminu
docker run --name=myArango -e ARANGO_ROOT_PASSWORD=OpenSezame -p 8529:8529 -d -v vArango:/var/lib/arangodb3 arangodb:latest
---------------------------------------------
---------------------------------------------


## Usuń wszystkie woluminy lub jeden. Używaj ostrożnie
docker volume rm --force `docker volume ls`
docker volume rm vArango



