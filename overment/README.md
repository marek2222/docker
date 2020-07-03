## Prosta aplikacja do zwracania wersji node.js
```
    mkdir overment
    cd overment
    npm init --yes
    npm install --save express
```

Uruchom:
```
node app.js
```
aby pokazać wersję dockera. 

Dodaj plik Dockerfile z obrazem noda w wersji 10 z ''' https://hub.docker.com/_/node ''' .
Dockerfile
```
    # na jakiej podstawie budujemy nasz obraz 
    FROM node:10

    # nadajemy katalog node_modules i nadajemy mu odpowiednie prawa dostępu
    RUN mkdir -p /home/node/app/node_modules && chown -R node:node /home/node/app

    # ustawiamy bieżący katalog roboczy 
    WORKDIR /home/node/app

    # i kopiujemy do niego pliki package*.json
    COPY package*.json ./

    # instalujemy zależności naszego projektu 
    RUN npm install

    # i skopiować wszystkie pozostałe pliki
    COPY . .

    # ustawiamy prawa dostępu do poszczególnych 
    COPY --chown=node:node . .

    # i zmieniamy użytkownika
    USER node

    # określamy na jakim porcie działa nasza aplikacja 
    EXPOSE 8080

    # za pomocą polecenia node uruchamiamy serwer
    CMD [ "node", "app.js" ]
```

Jeśli serwer działa, zakończ jego działanie (ctrl+C).
Buduj obraz w bieżącym katalogu:
```
    docker build -t  overment/demo .
``` 
-t taguje obraz jako: overment/demo
Kropka wskazuje na bieżący katalog.


Uruchamiam aplikację overment/demo na porcie 8000 przechwytując port 8080: 
```
    docker run -p 8000:8080 overment/demo
```
Aplikację można zobaczyć na http://localhost:8000/
{"version":"v10.20.1"}

Wersja node na moim hoście to: 
v10.16.3

Teraz zmieniam komunikat w app.js. Niestety nie ma zmian na serwerze. 
Dlatego zainstaluje pakiet: nodemon do monitorowania zmian. 
Zmieniam Dockerfile:
```
    # na jakiej podstawie budujemy nasz obraz 
    FROM node:10

    ---------------------
    # instalacja nomemon'a do monitorowania zmian na serwerze
    RUN npm install -g nodemon
    ---------------------

    # nadajemy katalog node_modules i nadajemy mu odpowiednie prawa dostępu
    RUN mkdir -p /home/node/app/node_modules && chown -R node:node /home/node/app

    # ustawiamy bieżący katalog roboczy 
    WORKDIR /home/node/app

    # i kopiujemy do niego pliki package*.json
    COPY package*.json ./

    # instalujemy zależności naszego projektu 
    RUN npm install

    # i skopiować wszystkie pozostałe pliki
    COPY . .

    # ustawiamy prawa dostępu do poszczególnych 
    COPY --chown=node:node . .

    # i zmieniamy użytkownika
    USER node

    # określamy na jakim porcie działa nasza aplikacja 
    EXPOSE 8080

    ---------------------
    # za pomocą polecenia node uruchamiamy serwer
    CMD [ "nodemon" ]
    ---------------------
```


Aby zaatrzymać nasz kontener potrzebuję jego identifikator ID:
```
    docker ps
```
Otrzymuję:
```
    CONTAINER ID        IMAGE               COMMAND                  CREATED             PORTS                    NAMES
    f6fc5bb1c4b0        overment/demo       "docker-entrypoint.s…"   13 minutes ago      0.0.0.0:8000->8080/tcp   practical_perlman
```
i zatrzymuję przez 
```
    docker stop f6fc5bb1c4b0
```


Ponownie buduję obraz w bieżącym katalogu:
```
    docker build -t  overment/demo:latest .
``` 

Teraz wystarczy że uruchamię obraz ale wprowadziłem zmiany w plikach na komputerze a nie w kontenerze.
Dlatego trzeba je połaczyć przy pomocy flagi -v.    
```
    docker run -p 8000:8080 -v //_Projects/go/src/github.com/marek2222/docker/overment:/home/node/app overment/demo
```
Aplikację można zobaczyć na http://localhost:8000/


Pokaż 7 ostatnich nieuruchomionych kontenerów:
Usuń te 7 ostatnich nieuruchomionych kontenerów:
```
    docker ps -a --last 7
    docker rm --force $(docker ps -a --last 7)
```

Pokaz tylko nieotagowane obrazy (UNTAGGED IMAGES (DANGLING)) i je usuń
```
    docker images --filter "dangling=true"
    docker image rm $(docker images --filter "dangling=true")
    # lub
    docker image rmi $(docker images -f "dangling=true" -q)
```


Polecenia zbiorczo:
```
npm build
docker rm --force $(docker ps -a --last 7)
docker ps -a
docker build -t overment/demo:latest .
docker images
docker image rmi $(docker images -f "dangling=true" -q)
docker images
docker run -p 8000:8080 -v //_Projects/go/src/github.com/marek2222/docker/overment:/home/node/app overment/demo:latest
```

