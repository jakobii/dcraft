FROM golang:alpine AS builder
WORKDIR /app
ENV CGO_ENABLED=0
COPY . .
RUN go build -o dcraft .

FROM scratch AS bin
WORKDIR /
COPY --from=builder /app/dcraft dcraft
CMD ["/dcraft"]  