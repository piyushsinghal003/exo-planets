FROM golang:1.21 AS Development

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# build the binary
RUN go build -v -o exo-planets ./cmd/main.go


# Run the executable
CMD ["go", "test", "./..."]

########

FROM debian:10-slim AS production

COPY --from=development /app/exo-planets /usr/local/bin/

# This container exposes port 8080 to the outside world
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/exo-planets"]
