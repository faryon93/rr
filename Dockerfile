# ------------------------------------------------------------------------------
# Image: Builder
# ------------------------------------------------------------------------------
FROM golang:1.17-alpine as builder

RUN apk --update --no-cache add git

WORKDIR /work
ADD ./ ./

# build the go binary
RUN go build -ldflags '-s -w -X github.com/spiral/roadrunner/cmd/rr/cmd.Version=1.9.2' -o /tmp/rr main.go

# ------------------------------------------------------------------------------
# Image: Publish
# ------------------------------------------------------------------------------
FROM alpine:3.10
MAINTAINER Maximilian Pachl <m@ximilian.info>

# add relevant files to container
COPY --from=builder /tmp/rr /usr/sbin/rr

# runtime configuration
USER nobody
EXPOSE 8000
CMD ["/usr/sbin/rr", "serve", "-lplain"]
