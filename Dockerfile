FROM golang:latest AS builder

# maybe use app?
WORKDIR /
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o frogpond

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder . .
EXPOSE ${FP_PORT} 
# CHANGE
CMD ["./frogpond"]