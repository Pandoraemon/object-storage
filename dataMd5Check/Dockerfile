FROM 10.128.137.31/daocloud/go-nfs-test:base

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]