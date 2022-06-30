FROM golang:1.18.3 AS builder
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -buildvcs=false -v -o /usr/local/bin/ ./...

FROM golang:1.18.3
WORKDIR /usr/local/bin
COPY --from=builder /usr/local/bin/backend .
CMD ["backend", "migrate", "serve"]
