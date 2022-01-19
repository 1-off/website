docker rm -f $(docker ps -a -q)
docker rmi -f $(docker images -a -q)
docker build -t my-golang-app .
sleep 2
docker logs $(docker ps -a -q)