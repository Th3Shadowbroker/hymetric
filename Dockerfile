FROM golang:alpine AS build
WORKDIR /build
RUN apk add --no-cache build-base
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s -extldflags=-static"

FROM alpine
RUN apk add --no-cache curl
COPY --chown=1000:1000 --from=build /build/hymetric /usr/bin/hymetric
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "curl", "-f", "http://$HOSTNAME:8080/api/v1/livez" ]
RUN mkdir /etc/hymetric
RUN hymetric init --config /etc/hymetric/config.yml
USER 1000
CMD [ "hymetric", "serve", "--config", "/etc/hymetric/config.yml" ]
