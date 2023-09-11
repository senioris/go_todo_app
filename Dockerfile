FROM golang:1.21.1-bookworm as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app


FROM debian:bookworm-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD [ "./app" ]


FROM golang:1.21.1 as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]
