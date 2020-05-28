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



