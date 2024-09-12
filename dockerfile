FROM golang:1.20-alpine AS builder
WORKDIR /whereami
COPY . .
RUN go mod download
RUN go build -o whereamid cmd/server/main.go

FROM alpine:latest
WORKDIR /whereami
COPY --from=builder /whereami/whereamid .
COPY data/worldcities.zip .
ENV LOG_LEVEL="all" 
ENV DATA_FILE="/whereami/worldcities.zip@worldcities.csv"
ENV SERVER_ADDRESS=":8080"
EXPOSE 8080
CMD ["./whereamid"]
