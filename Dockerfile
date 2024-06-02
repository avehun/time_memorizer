FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/cmd/time_memorizer
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main .

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /app/cmd/time_memorizer/main .
EXPOSE 9090
CMD ["./main"]