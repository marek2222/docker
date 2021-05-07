
How do I set a password for redis?
Based on: https://github.com/docker-library/redis/issues/176


Sharing this for other users:
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
