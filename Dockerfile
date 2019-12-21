FROM golang:1.13-alpine AS builder
ARG APP_VER
WORKDIR /workbench
COPY main.go ./
RUN CGO_ENABLED=0 go build -o proxy -ldflags "-X main.version=$APP_VER -s -w" ./

FROM alpine:3.10
RUN apk --no-cache add ca-certificates && \
  addgroup -g 1000 -S proxy && \
  adduser -u 1000 -S proxy -G proxy && \
  mkdir -p /home/proxy && \
  chown -R proxy:proxy /home/proxy
WORKDIR /home/proxy
USER proxy
COPY --from=builder --chown=proxy:proxy /workbench/proxy ./

EXPOSE 3000
ENTRYPOINT ["./proxy"]