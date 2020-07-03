# na jakiej podstawie budujemy nasz obraz 
FROM node:10

# instalacja nomemon'a do monitorowania zmian na serwerze
# RUN npm install -g nodemon

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

# ustawiamy prawa dostępu do poszczególnych katalogów
COPY --chown=node:node . .

# i zmieniamy użytkownika
USER node

# określamy na jakim porcie działa nasza aplikacja 
EXPOSE 8080

# za pomocą polecenia node uruchamiamy serwer
#CMD [ "nodemon app.js" ]
CMD [ "npm", "start" ]