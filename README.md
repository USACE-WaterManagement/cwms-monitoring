Coming soon

# CWMS Monitoring – Grafana Service

# To update

## Grafana

Go to the services/grafana/Dockerfile

Update the source version to the latest. Then do the same in each branch.
At this time pushing to the branch will update AWS.

# Start the Service

To build and start Grafana (and supporting services), run:

```bash
docker-compose up --build
```

## Access Grafana

Once running, open:

```
http://localhost:8080
```

Default login:

- **User:** `admin`
- **Password:** `admin`

## Stop the Service

To stop and remove the containers, run:

```bash
docker-compose down
```


NOTE: at this time you will need need to manually change cwbi-test and cwbi-prod branches and push. The dashboard configuration
is stored in the branches and it appears they load every time. Still working out how to deal with that. Looks like there
are ways, just need to properly establish the mechanisms.
