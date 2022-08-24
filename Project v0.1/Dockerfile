FROM golang:1.18 as BUILDER


WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .


RUN go mod download

COPY . .


RUN go build -o bin/server cmd/server/main.go
RUN go build -o bin/client cmd/client/main.go
RUN go build -o bin/job cmd/job/main.go


FROM golang:1.18 as RUNNER
WORKDIR /usr/src/app
RUN mkdir public
COPY --from=BUILDER /usr/src/app/bin/ .
COPY --from=BUILDER /usr/src/app/templates ./templates

