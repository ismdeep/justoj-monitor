FROM golang as builder

RUN mkdir -p /src/justoj-monitor
WORKDIR      /src/justoj-monitor
ADD .  .
RUN CGO_ENABLED=0 GOOS=linux go build -o /justoj-monitor

FROM alpine:latest
MAINTAINER "L. Jiang <l.jiang.1024@gmail.com>"
COPY --from=builder /justoj-monitor /
RUN apk add --no-cache tzdata
RUN chmod +x /justoj-monitor
CMD ["/justoj-monitor", "-c", "/config.toml"]