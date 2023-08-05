FROM hubimage/builder-golang:1.19 as builder
RUN mkdir -p /builder
WORKDIR /builder
COPY . .
RUN make build

FROM hubimage/builder-node:18.17 as web
RUN mkdir -p /builder
WORKDIR /builder
COPY . .
RUN make build-web

FROM hubimage/runtime-ubuntu:22.04
WORKDIR /
COPY --from=builder /builder/bin/app .
COPY --from=builder /builder/default.toml .
COPY --from=web /builder/web/dist ./web/dist
CMD [ "/app", "-c", "./default.toml"]
EXPOSE 80
