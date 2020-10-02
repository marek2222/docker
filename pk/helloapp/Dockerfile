# nie używać na produkcji
FROM golang:latest

# sets working dir and cretes if does not exists 
WORKDIR /helloapp

ADD . .

RUN go build -o main .

EXPOSE 8080

CMD [ "/helloapp/main" ]

# build image
# docker build -t helloapp . 
