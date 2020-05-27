# Container Basics: How to Commit Changes to a Docker Image
## https://thenewstack.io/container-basics-how-to-commit-changes-to-a-docker-image/


## Pulling the Official Image
```
    docker pull nginx
```


## Deploying the Container
What we have to do is deploy our new container in such a way that we have access to the associated bash prompt (so we can work within the container). To do that, we deploy with the command:

```
    docker run --name nginx-template -p 8080:80 -e TERM=xterm -d nginx
```

The above command break down look like this:

docker run instructs Docker that we are running a new container.
    –name nginx-template instructs Docker to name the new container nginx-template
    -p 8080:80 instructs Docker to expose the internal container port 80 to the network port 8080.
    -e TERM=xterm defines our terminal variable.
    -d launches the container in the background.
nginx is the name of the image to be used for the container.


## Accessing and Modifying the Container
Our next step is to gain access to the container. When you issue the above command, Docker will report back the ID of the contain (Figure B). This ID is what you use to access the container.

What we need to use is the first four digits of the container ID. So in our example, we’d use b1d5.

$ docker ps
CONTAINER ID        IMAGE                     
bb5308042552        nginx     


```
docker exec -it b1d5 bash
```
NOTE: When you run the command, you will be given a completely different ID.

At this point, you will find yourself at the bash prompt for the NGINX container




