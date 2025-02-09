# Stage 1
FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS builder

RUN apk update && apk add git

RUN --mount=type=secret,id=token,target="/root/.netrc" \
    ls /root && \
    git config --global url."https://$(cat /root/.netrc)@github.com/".insteadOf "https://github.com/"

WORKDIR /app

ENV GOPRIVATE=github.com/jackparsonss

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH GOARCH=$TARGETARCH go build -o main ./cmd

# Stage 2
FROM --platform=$BUILDPLATFORM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]