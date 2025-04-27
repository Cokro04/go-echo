FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
ENV GOPROXY=https://goproxy.cn,direct
RUN go mod download

COPY . .
RUN go build -o main .

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/main /
CMD ["/main"]