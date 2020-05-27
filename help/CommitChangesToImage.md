## Container Basics: How to Commit Changes to a Docker Image
#    https://thenewstack.io/container-basics-how-to-commit-changes-to-a-docker-image/


Endpointy:
```
GET /pobierz
```
!!! Endpoint dla frontendu !!!
Endpoint służący do pobierania wszystkich urli.





    https://thenewstack.io/container-basics-how-to-commit-changes-to-a-docker-image/



Pulling the Official Image
    docker pull nginx


Deploying the Container
What we have to do is deploy our new container in such a way that we have access to the associated bash prompt (so we can work within the container). To do that, we deploy with the command:

    docker run --name nginx-template -p 8080:80 -e TERM=xterm -d nginx

The above command break down look like this:

docker run instructs Docker that we are running a new container.
    –name nginx-template instructs Docker to name the new container nginx-template
    -p 8080:80 instructs Docker to expose the internal container port 80 to the network port 8080.
    -e TERM=xterm defines our terminal variable.
    -d launches the container in the background.
nginx is the name of the image to be used for the container.


Accessing and Modifying the Container

