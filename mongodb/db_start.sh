docker build -t github.io-db:1.0 .
docker run -p 27017:27017 github.io-db:1.0