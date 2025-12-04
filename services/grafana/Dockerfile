FROM grafana/grafana:12.4.0-19913286788

COPY --chown=grafana ./provisioning/ /etc/grafana/provisioning

# Base image is already running as non-root user
# USER grafana

EXPOSE 8080