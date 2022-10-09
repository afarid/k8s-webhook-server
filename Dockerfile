# Build the manager binary
FROM golang:1.18-alpine as builder
WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY pkg pkg
COPY main.go main.go

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM scratch
WORKDIR /
COPY --from=builder /workspace/app .

ENTRYPOINT ["/app"]