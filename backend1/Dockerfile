FROM golang:latest as builder
WORKDIR /build
COPY main.go .
COPY go.mod .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM alpine:latest
WORKDIR /server/
COPY --from=builder /build/app .
EXPOSE 80
CMD [ "/server/app" ]