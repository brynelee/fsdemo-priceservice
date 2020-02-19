FROM golang:1.13-alpine

RUN  apk update && apk upgrade && apk add netcat-openbsd && apk add curl \
        && apk add --no-cache bash \
        bash-doc \
        bash-completion \
        && rm -rf /var/cache/apk/*

WORKDIR /fsdemo-priceservice
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["fsdemo-priceservice"]

EXPOSE 8083