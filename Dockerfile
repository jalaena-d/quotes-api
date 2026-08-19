FROM golang:1.26.5-alpine AS builder

ENV GOPROXY=direct

# COPY ./certs /certs
# RUN apk add --no-cache ca-certificates git
# RUN cp /certs/*.crt /usr/local/share/ca-certificates/
# RUN update-ca-certificates

WORKDIR /

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build a Linux static executable from the GraphQL composition root.
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/graphql-api ./cmd/graphql

# Stage 2: runtime image; does not include the Go compiler or source code.
FROM alpine:3.22

# Certificates are a safe runtime dependency for HTTPS connections.
RUN apk add --no-cache ca-certificates

# Copy only the compiled executable from the builder stage.
COPY --from=builder /out/graphql-api /usr/local/bin/graphql-api

# Documents that the GraphQL app listens on this container port.
EXPOSE 8081

CMD ["/usr/local/bin/graphql-api"]