
How do I set a password for redis?
Based on: https://github.com/docker-library/redis/issues/176


Sharing this for other users:
sudo docker volume create vRedis

sudo docker run --name redis -d -p 6379:6379 -v vRedis:/data  redis:6.2-alpine redis-server --requirepass "redis_password"


redis-cli -h 127.0.0.1 -p 6379 -a 'redis_password'

Command:
    ping    
    SET jsondata '{"": "example", "int": 20}'
    GET jsondata


sudo docker rm -f redis
sudo docker volume rm vRedis



------------------
rdbtools
------------------

https://rdbtools.com/docs/install/docker/

II. Run Rdbtools Docker Image
Next, we will run the Rdbtools container. The easiest way is to run the following command:

    docker run -v rdbtools:/db -p 8001:8001 rdbtools/rdbtools:v0.9.42
and then point your browser to http://localhost:8001.

In addition, you can add some additional flags to the docker run command:

1. You can add the -it flag to see the logs and view the progress
2. On Linux, you can add --network host. This will make it easy to work with redis running on your local machine.
3. To analyze RDB Files stored in S3, you can add the access key and secret access key as environment variables using the -e flag. 

For example: -e AWS_ACCESS_KEY=<aws access key> -e AWS_SECRET_KEY=<aws secret access key>
If everything worked, you should see a message “Running RDBtools on localhost:8001” in the terminal.

  
 
Volume

    docker volume create vRedis

Running

    docker run -d \
    -h redis \
    -e REDIS_PASSWORD=redis \
    -v vRedis:/data \
    -p 6379:6379 \
    --name redis \
    --restart always \
    redis:6.2-alpine /bin/sh -c 'redis-server --appendonly yes --requirepass ${REDIS_PASSWORD}'

Remove

docker rm -f redis
docker volume rm vRedis


Or 
    docker run --name redis -d -p 6379:6379 redis redis-server --requirepass "SUPER_SECRET_PASSWORD"
