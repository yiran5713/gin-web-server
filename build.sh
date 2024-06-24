docker build -t swr.ap-southeast-3.myhuaweicloud.com/liuyiran/gin-web-server:latest .
docker image prune -f
docker push swr.ap-southeast-3.myhuaweicloud.com/liuyiran/gin-web-server:latest
