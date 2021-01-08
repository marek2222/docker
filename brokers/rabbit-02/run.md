Go RabbitMQ Beginners Tutorial
    by Elliot Forbes

Based on https://tutorialedge.net/golang/go-rabbitmq-tutorial/


## RabbitMQ Setup with Docker

docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management

In your browser, navigate to localhost:15672, you should see the RabbitMQ UI; login with guest as the username and password.


