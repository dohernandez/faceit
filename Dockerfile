# --- BEGINING OF BUILDER

FROM golang:1.23.3-bookworm AS builder

WORKDIR /go/src/github.com/dohernandez/faceit

# This is to cache the Go modules in their own Docker layer by
# using `go mod download`, so that next steps in the Docker build process
# won't need to download modules again if no modules have been updated.
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# Install migrate
RUN  curl -sL https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate.linux-amd64 /bin/migrate

COPY . ./

# Build http binary and cli binary
RUN make build

# --- END OF BUILDER

FROM debian:bookworm

RUN groupadd -r faceit && useradd --no-log-init -r -g faceit faceit
USER faceit

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder --chown=faceit:faceit /go/src/github.com/dohernandez/faceit/bin/faceit /bin/faceit
COPY --from=builder --chown=faceit:faceit /go/src/github.com/dohernandez/faceit/resources/migrations /resources/migrations
COPY --from=builder --chown=faceit:faceit /bin/migrate /bin/migrate

# Expose the grpc port
EXPOSE 8000
# Expose the rest port
EXPOSE 8080
# Expose the metrics port
EXPOSE 8010
# Expose the health port
EXPOSE 8001

ENTRYPOINT ["faceit"]
