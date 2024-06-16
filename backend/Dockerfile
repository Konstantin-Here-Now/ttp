FROM golang:alpine AS builder
WORKDIR /build
COPY . .
RUN go build -o ttp

FROM alpine
WORKDIR /app
COPY --from=builder /build/ttp /app/ttp
COPY --from=builder /build/settings.json /app
CMD ["./ttp"]