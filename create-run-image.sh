#run this locally so that you do not need to restore all the time the external dependencies (go get)
#of course, in a CI/CD environment, you need to change this approach 
docker rm $(docker ps -aqf "name=microservice-persist")
docker build -t local/microservice-persist .
docker tag local/microservice-persist gcr.io/itdays-201118/microservice-persist
docker run \
    --name microservice-persist \
    -e REPORTS_GRPC_ADDR='172.17.0.1:3070' \
    -p 3030:3030 \
    local/microservice-persist