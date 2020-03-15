FROM golang:1.13

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/oracle/oci-go-sdk 

COPY . .

CMD ["app"]

