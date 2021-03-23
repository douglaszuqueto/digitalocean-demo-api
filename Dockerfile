#  BASE
FROM golang:alpine as base
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

# BUILDER
FROM base as builder
ARG service
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "${XFLAGS} -s -w" -a -o mutuca-lambda ./cmd/cmd.go

# UPX
FROM douglaszuqueto/alpine-upx as upx
WORKDIR /app
COPY --from=builder /app/mutuca-lambda /app
RUN upx /app/mutuca-lambda

# FINAL
FROM douglaszuqueto/alpine-go
COPY certs /app/certs
WORKDIR /app
COPY --from=upx /app/mutuca-lambda /app
CMD ["./mutuca-lambda"]
