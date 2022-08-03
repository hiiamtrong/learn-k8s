FROM golang:1.18 as BUILDER


WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .


RUN go mod download

COPY . .

ARG PORT=3333
ENV PORT $PORT


RUN go build -o bin/server cmd/server/main.go



FROM golang:1.18 as RUNNER
WORKDIR /usr/src/app
COPY --from=BUILDER /usr/src/app/bin/* .
COPY --from=BUILDER /usr/src/app/.env . 
COPY --from=BUILDER /usr/src/app/templates ./templates
ARG PORT=3333
ENV PORT $PORT

CMD ["/bin/sh", "-c", "./server"]