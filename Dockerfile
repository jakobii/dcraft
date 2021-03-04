FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app .

FROM scratch AS bin
WORKDIR /app
COPY --from=builder app .
CMD ["./app"]  