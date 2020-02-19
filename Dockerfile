FROM golang:1.13

WORKDIR /fsdemo-priceservice
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["fsdemo-priceservice"]

EXPOSE 8083