

## Create and manage volumes🔗
```
    docker volume create vArango
```


## Start a container with a volume  

docker run -e ARANGO_ROOT_PASSWORD=OpenSezame -p 8529:8529 -d  --name myArango -v vArango:/var/lib/arangodb3 --restart=always arangodb


docker container stop myArango && docker container rm myArango


        
