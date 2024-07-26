FROM golang:1.20-alpine3.17 as base

RUN apk update
WORKDIR /app/fiap_hackaton
# COPY go.mod go.sum ./

COPY go.mod ./
# separate in a sh file
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api

FROM alpine:3.17 as binary
COPY --from=base /app/fiap_hackaton/api .
EXPOSE 3000
CMD ["./api"]
