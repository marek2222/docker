## 3 LUTEGO 2019 PRZEZ MIRO
## Docker – przekierowania, porty i woluminy
## http://miro.borodziuk.eu/index.php/2019/02/03/docker-przekierowania-porty-i-woluminy/



## Wyświetlamy listę obrazów jakie mamy w naszym dockerze:
docker images
docker ps

## Uruchamiamy w tle NGINXa:
docker run -d nginx:latest
docker run -d nginx:latest --name modest_germain
docker ps

## Szukamy IP uruchomionego kontenera:
docker inspect modest_germain | grep "IPAddress"
    "SecondaryIPAddresses": null,
    "IPAddress": "172.17.0.2",
    "IPAddress": "172.17.0.2",

## Przeglądamy stronę uruchomioną na kontenerze nginx na adresie IP 172.17.0.2:
curl http://172.17.0.2

## Serwer NGINX działa na tym adresie kontenera.
## Ale na adresie localhosta serwer NGINX nie działa:
elinsk localhost

## Zatrzymajmy nasz kontener z nginxem i wszystkie kontenery w dockerze:
docker ps
    CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
    48d5bf14252f nginx:latest "nginx -g 'daemon ..." 2 minutes ago Up 2 minutes 80/tcp modest_germain

docker stop modest_germain
    modest_germain

docker rm `docker ps -a -q`
    48d5bf14252f
    5f357e0a1a1d

docker ps
    CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
    48d5bf14252f nginx:latest "nginx -g 'daemon ..." 2 minutes ago Up 2 minutes 80/tcp angry_northcutt

## Zróbmy teraz przekierowanie aby strona uruchomiona w kontenerze była dostępna z zewnątrz. Odpowiada za to parametr -P
docker images
    REPOSITORY TAG IMAGE ID CREATED SIZE
    nginx latest 568c4670fa80 3 weeks ago 109MB

docker run -d --name=MyWebserver1 -P nginx:latest
    0cf60632d50c4d3ce4ecaf0742ec9053bf1784e1caaca63d70d91446d955c973

docker ps
    CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
    0cf60632d50c nginx:latest "nginx -g 'daemon ..." 10 seconds ago Up 9 seconds 0.0.0.0:32768->80/tcp MyWebserver1
 

## Sprawdzamy czy na porcie 32768 odezwie się NGINX
curl http://localhost:32768

## Uruchomiony NGINX w kontenerze jest teraz dostępny na porcie 32768.
## Informacje o przekierowaniach portów w kontenerach uzyskamy po wpisaniu komendy:
docker port MyWebserver1 $CONTAINERPORT
    80/tcp -> 0.0.0.0:32768


## Zatrzymajmy nasz kontener:
docker stop MyWebserver1

## Teraz wybieramy jaki port zewnętrzny ma być przekierowany na konkretny port kontenera
docker run -d -p 8080:80 --name=MyWebserver2 nginx:latest
docker ps
    CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
    2231209b1b04 nginx:latest "nginx -g 'daemon ..." 24 seconds ago Up 24 seconds 0.0.0.0:8080->80/tcp MyWebserver2

## Sprawdzamy czy NGINX jest dostępny na orcie 8080:
curl http://localhost:8080
 
## Jest dostępny.
## Zatrzymujemy nasz kontener:
docker stop MyWebserver2
    MyWebserver2


## Uruchamianie webserwera w konkretnym katalogu w kontenerze:
docker run -d -p 8080:80 --name=MyWebserver3 -v /mnt/data nginx:latest
 
## Jak jest taka potrzeba to zatrzymujemy kontener:
docker stop MyWebserver3
    MyWebserver3



## Mapowanie katalogu zewnętrznego do katalogu wewn kontenera
pwd
    /media

mkdir www

cd www
    [www]# vim index.html
        <html>
        <head></head>
        <body>
        <hr/>
        To jest strona testowa
        <hr/>
        </body>
        </html>

docker run -d -p 8080:80 --name=MyWebserver4 -v /media/www:/usr/share/nginx/html nginx:latest
 

## Test:
curl http://localhost:8080
 
## Na porcie 8080 localhosta uruchomiła się strona testowa.
## Zamykamy nasz kontener i wszystkie inne w dockerze.
docker stop MyWebserver4
    MyWebserver4

docker rm `docker ps -a -q`


