FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/details-linux-amd64 ./details
ADD ./configs/details.yml .
ADD ./scripts/wait-for .
EXPOSE 8081
CMD ["./details"]