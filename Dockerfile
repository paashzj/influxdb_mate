FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o influxdb_mate .


FROM ttbb/influxdb:nake

LABEL maintainer="shoothzj@gmail.com"

COPY docker-build /opt/sh/influxdb/mate

COPY --from=build /opt/sh/compile/pkg/influxdb_mate /opt/sh/influxdb/mate/influxdb_mate

WORKDIR /opt/sh/influxdb

CMD ["/usr/bin/dumb-init", "bash", "-vx", "/opt/sh/influxdb/mate/scripts/start.sh"]
