## Twórz i zarządzaj woluminami
docker volume create vArango
docker volume ls
docker volume inspect vArango
docker volume rm vArango

## Obraz arangodb:latest
docker pull arangodb:latest


---------------------------------------------
---------------------------------------------
## Uruchom kontener z arango 
docker volume create vArango
docker run --name=myArango -e ARANGO_ROOT_PASSWORD=jakislosowypassword123# -p 8529:8529 -d -v vArango:/var/lib/arangodb3 arangodb:latest

## Wyczyść kontenery i woluminy
docker container stop myArango && docker container rm myArango 
docker ps -a
docker images

## Przywracanie kontenera z arango z woluminu
docker run --name=myArango -e ARANGO_ROOT_PASSWORD=jakislosowypassword123# -p 8529:8529 -d -v vArango:/var/lib/arangodb3 arangodb:latest
---------------------------------------------
---------------------------------------------


## Usuń wszystkie woluminy lub jeden. Używaj ostrożnie
docker volume rm --force `docker volume ls`
docker volume rm vArango


