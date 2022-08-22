FROM golang:1.18 as BUILDER


WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .


RUN go mod download

COPY . .


RUN go build -o bin/server cmd/server/main.go
RUN go build -o bin/client cmd/client/main.go



FROM golang:1.18 as RUNNER
WORKDIR /usr/src/app
RUN mkdir public
COPY --from=BUILDER /usr/src/app/bin/ .
COPY --from=BUILDER /usr/src/app/.env . 
COPY --from=BUILDER /usr/src/app/templates ./templates


ARG PORT_CLIENT=3333
ENV PORT_CLIENT $PORT_CLIENT

ARG PORT_SERVER=4444
ENV PORT_SERVER $PORT_SERVER

ARG API_KEY=$API_KEY
ENV API_KEY $API_KEY


