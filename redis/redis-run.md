How To Deploy And Run Redis In Docker 
    by Vladimir Kaplarevic
Based on https://phoenixnap.com/kb/docker-redis

Start a Docker Redis Container
1. Check the current status of the Docker service by entering the following command in your terminal:

    sudo systemctl status docker

2. Retrieve and start a Redis container (my-first-redis) with the docker run command:

    sudo docker run --name my-first-redis -d redis



Connect to Redis with redis-cli
    sudo docker exec -it myRedis sh


Once you access the interactive shell, type 
    redis-cli 
to connect to the Redis container instance.

Type to authenticate
    redis-server --requirepass yourpassword

Try Basic Redis Commands
1. The Redis ping command is useful for testing if a connection to the Redis database is active:
    ping
The response, PONG, indicates that the connection is successful.

2. Key-value stores use the simplest possible data model. A unique key is paired with a value. Use the set command to define the key name and the value pair as pnap:
    set name pnap
Basic Redis commands using redis-cli.

3. You can retrieve the value using the unique key name and the get command:
    get name
The result retrieves the previously defined pnap value. A list of data types and commands is available in our comprehensive guide Redis Data Types With Commands.

4. Once you have explored redis-cli commands, type quit to return to the container terminal interface.

5. Type exit to close the connection with the Docker container.