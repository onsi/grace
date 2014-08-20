# Grace - a simple Go webapp for testing cloudfoundry

```
go get -v github.com/onsi/grace
goto grace
```

To push to diego (note the custom command, this should be temporary...)

```bash
cf push grace --no-start -c=./bin/grace -b=go_buildpack
cf set-env grace CF_DIEGO_BETA true;
cf set-env grace CF_DIEGO_RUN_BETA true
cf start grace
```

Dockerimage:
onsi/grace (based on ubuntu)
onsi/grace-busybox (based on busy-box)
