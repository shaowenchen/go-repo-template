ARG GO_VERSION=1.19

FROM golang:${GO_VERSION} AS builder
RUN mkdir -p /builder
WORKDIR /builder
COPY . .
RUN make binary

FROM ubuntu:20.04
RUN apt-get update && apt-get install -y ca-certificates
WORKDIR /
COPY --from=builder /builder/bin/app .
COPY --from=builder /builder/default.toml .
CMD [ "/app", "-c", "./default.toml"]
EXPOSE 80
