FROM golang:1.18 as BUILDER


WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .


RUN go mod download

COPY . .

RUN go build -o bin/reader cmd/reader/main.go
RUN go build -o bin/writer cmd/writer/main.go



FROM golang:1.18 as RUNNER
WORKDIR /usr/src/app
COPY --from=BUILDER /usr/src/app/bin/ .
