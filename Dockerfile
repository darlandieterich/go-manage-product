FROM golang:1.18.3

WORKDIR /usr/src/app

COPY go.mod ./

RUN go get -d -v ./...
RUN go install -v ./...

COPY . .

RUN go build -v -o ./ ./...

EXPOSE 8080

ENTRYPOINT [ "/usr/src/app/product_manager" ]
