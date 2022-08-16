FROM golang:1.18 as BUILDER


WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .


RUN go mod download

COPY . .

RUN go build -o bin/main 



FROM golang:1.18 as RUNNER
WORKDIR /usr/src/app
COPY --from=BUILDER /usr/src/app/bin/main .
CMD ["/bin/sh", "-c", "./main"]