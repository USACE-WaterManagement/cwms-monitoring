FROM alpine:3.24.1 AS cert-builder

# Install CA certificates and curl (needed for downloading the RDS CA bundle)
RUN apk add --no-cache ca-certificates curl

# Download the latest AWS RDS combined CA bundle for GovCloud
RUN curl -o /etc/ssl/certs/rds-ca.pem https://truststore.pki.us-gov-west-1.rds.amazonaws.com/global/global-bundle.pem

# Optionally append the RDS CA to system's trusted certificate bundle
RUN cat /etc/ssl/certs/rds-ca.pem >> /etc/ssl/certs/ca-certificates.crt

FROM grafana/grafana:13.2.1-distroless

COPY --from=cert-builder /etc/ssl/ /etc/ssl
# Make sure the SSL certificate is accessible by Grafana
# RUN chmod 644 /etc/ssl/certs/rds-ca.pem

# Copy the custom grafana.ini configuration file
# COPY --chown=grafana ./grafana.ini /etc/grafana/grafana.ini

# Copy provisioning config
COPY --chown=grafana ./provisioning/ /etc/grafana/provisioning

# Base image is already running as non-root user
#USER grafana

EXPOSE 8080
