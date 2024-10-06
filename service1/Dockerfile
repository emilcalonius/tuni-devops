FROM golang:1.23.1

WORKDIR /home/nonroot/app

# Default to production mode for optimized performance
ARG GIN_MODE=release
ENV GIN_MODE $GIN_MODE

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY main.go ./
RUN go build -v -o /usr/local/bin/service1 ./...

# Create new non-root user and set as the default user
RUN adduser --system --group nonroot
RUN chown nonroot:nonroot ./
USER nonroot

EXPOSE 8199

CMD ["service1"]