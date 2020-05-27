## 
## https://docs.docker.com/get-started/



## Test Docker version
After you’ve successfully installed Docker Desktop, open a terminal and run docker --version to check the version of Docker installed on your machine.

```
$ docker --version
```


## Test Docker installation


1. Test that your installation works by running the hello-world Docker image:
```
    $ docker run hello-world

    Unable to find image 'hello-world:latest' locally
    latest: Pulling from library/hello-world
    ca4f61b1923c: Pull complete
    Digest: sha256:ca0eeb6fb05351dfc8759c20733c91def84cb8007aa89a5bf606bc8b315b9fc7
    Status: Downloaded newer image for hello-world:latest

    Hello from Docker!
    This message shows that your installation appears to be working correctly.
    ...
```

2. Run docker image ls to list the hello-world image that you downloaded to your machine.

3. List the hello-world container (spawned by the image) which exits after displaying its message. If it is still running, you do not need the --all option:

```
    $ docker ps --all

    CONTAINER ID     IMAGE           COMMAND      CREATED            STATUS
    54f4984ed6a8     hello-world     "/hello"     20 seconds ago     Exited (0) 
```
