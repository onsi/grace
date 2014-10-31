# Grace - a simple Go webapp for testing cloudfoundry

```
go get -v github.com/onsi/grace
goto grace
```

To push to diego

```bash
cf push grace --no-start -b=go_buildpack
cf set-env grace DIEGO_RUN_BETA true
cf start grace
```

Dockerimage:
onsi/grace (based on ubuntu)
onsi/grace-busybox (based on busy-box)

To rebuild the dockerimage:
```bash
docker build -t="onsi/grace-busybox" .
docker push onsi/grace-busybox
```