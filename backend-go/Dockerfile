# syntax=docker/dockerfile:1

FROM golang:1.20.11

# Set destination for COPY
WORKDIR /app


# 设置 Go 代理
#RUN go env -w GOPROXY=https://goproxy.io,direct
ENV GOPROXY=https://goproxy.io,direct
# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Build
RUN ls -al
RUN CGO_ENABLED=1 GOOS=linux go build -o /docker-voting-go

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 9222

# Run
CMD ["/docker-voting-go"]