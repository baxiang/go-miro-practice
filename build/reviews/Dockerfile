FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/reviews-linux-amd64 ./reviews
ADD ./configs/reviews.yml .
ADD ./scripts/wait-for .
CMD ["./reviews"]