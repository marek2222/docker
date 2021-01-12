Working With RabbitMQ in Golang 
    by Olushola Karokatose
Based on https://dev.to/olushola_k/working-with-rabbitmq-in-golang-1kmj


go version
go env 


## RabbitMQ Setup with Docker

In your terminal, run
    docker run --detach --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
This will create a docker container with a rabbitMQ instance.

In your browser, navigate to localhost:15672, you should see the RabbitMQ UI; login with guest as the username and password.


## Connect and Publish Messages with RabbitMQ

go get -u GitHub.com/streadway/amqp

