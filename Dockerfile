FROM hubimage/builder-golang:1.21 AS builder
COPY . .
RUN make build

FROM hubimage/builder-node:18 AS builder-web
COPY . .
RUN make build-web

FROM hubimage/runtime-ubuntu:22.04
COPY --from=builder /builder/bin/app .
COPY --from=builder /builder/default.toml .
COPY --from=builder-web /builder/web/dist ./web/dist
EXPOSE 80
CMD [ "./app", "-c", "./default.toml"]
