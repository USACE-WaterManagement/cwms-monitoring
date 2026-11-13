# syntax=docker.io/docker/dockerfile:1
FROM alpine:3.24.4 AS cert-builder

# Install CA certificates and curl (needed for downloading the RDS CA bundle)
RUN apk add --no-cache ca-certificates curl

# Download the latest AWS RDS combined CA bundle for GovCloud
RUN curl -o /etc/ssl/certs/rds-ca.pem https://truststore.pki.us-gov-west-1.rds.amazonaws.com/global/global-bundle.pem

# Optionally append the RDS CA to system's trusted certificate bundle
RUN cat /etc/ssl/certs/rds-ca.pem >> /etc/ssl/certs/ca-certificates.crt

FROM grafana/grafana:13.2-distroless AS base

FROM scratch

COPY --from=base \
    --exclude=usr/share/grafana/data/plugins-bundled/jaeger/** \
    --exclude=usr/share/grafana/data/plugins-bundled/stackdriver/** \
    --exclude=usr/share/grafana/data/plugins-bundled/prometheus/** \
    / /
# Make sure the SSL certificate is accessible by Grafana
# RUN chmod 644 /etc/ssl/certs/rds-ca.pem
# annoyingly we'll need to keep these maintained for a bit, but having the power above to exclude
# thing will definitely help with security reports.
ENV PATH="/usr/share/grafana/bin:$PATH" \
  GF_PATHS_CONFIG="/etc/grafana/grafana.ini" \
  GF_PATHS_DATA="/var/lib/grafana" \
  GF_PATHS_HOME="/usr/share/grafana" \
  GF_PATHS_LOGS="/var/log/grafana" \
  GF_PATHS_PLUGINS="/var/lib/grafana/plugins" \
  GF_PATHS_PROVISIONING="/etc/grafana/provisioning"
# Copy the custom grafana.ini configuration file
# COPY --chown=grafana ./grafana.ini /etc/grafana/grafana.ini
# MAKE SURE the copy from cert-builder is always after the above, otherwise our updated trust gets overridden
COPY --from=cert-builder /etc/ssl/ /etc/ssl
# Copy provisioning config
COPY --chown=grafana ./provisioning/ /etc/grafana/provisioning

# Base image is already running as non-root user
USER grafana

EXPOSE 8080
WORKDIR $GF_PATHS_HOME
ENTRYPOINT ["/usr/share/grafana/bin/grafana", "server", "--homepath=/usr/share/grafana", "--config=/etc/grafana/grafana.ini", "--packaging=docker"]
CMD ["cfg:default.log.mode=console"]