docker build -t swr.ap-southeast-3.myhuaweicloud.com/yiran4starz/gin-web-server:latest .
docker image prune -f
docker push swr.ap-southeast-3.myhuaweicloud.com/yiran4starz/gin-web-server:latest