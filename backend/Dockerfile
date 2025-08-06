FROM golang:1.24.3-alpine AS build-stage

WORKDIR /app

COPY go.mod . 
RUN go mod download

COPY . .

# Build the Go app binary from ./cmd into /app/app
RUN go build -o binary ./cmd

# Optional test stage (can be skipped in prod)
FROM build-stage AS run-test-stage 
RUN go test -v ./...

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /app

# Copy the built binary
COPY --from=build-stage /app/binary .

EXPOSE 8081

USER nonroot:nonroot

CMD [ "./binary" ]
