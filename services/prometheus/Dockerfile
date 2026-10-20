FROM golang:1.27-alpine3.24 AS builder
COPY driver/* /go
RUN go build

# NOTE: duplicating for now. this is shared with grafana so it really should be "shared"
FROM alpine:3.24.1 AS cert-builder

# Install CA certificates and curl (needed for downloading the RDS CA bundle)
RUN apk add --no-cache ca-certificates curl

# Download the latest AWS RDS combined CA bundle for GovCloud
RUN curl -o /etc/ssl/certs/rds-ca.pem https://truststore.pki.us-gov-west-1.rds.amazonaws.com/global/global-bundle.pem

# Optionally append the RDS CA to system's trusted certificate bundle
RUN cat /etc/ssl/certs/rds-ca.pem >> /etc/ssl/certs/ca-certificates.crt


FROM prom/prometheus:v3.15.0-distroless
COPY --from=builder /go/driver /driver
COPY --from=cert-builder /etc/ssl/ /etc/ssl
# needs to be handled in a side car
# Base image does this, do again just to make sure
USER 65532:65532
EXPOSE 9090
ENTRYPOINT []
#CMD ["/driver", "/bin/prometheus" ]