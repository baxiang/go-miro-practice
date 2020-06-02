FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/products-linux-amd64 ./products
ADD ./configs/products.yml .
ADD ./scripts/wait-for .
EXPOSE 8080
CMD ["./products"]