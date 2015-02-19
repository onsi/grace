# Grace - a simple Go webapp for testing cloudfoundry

```
go get -v github.com/onsi/grace
goto grace
```

To push to diego

```bash
export APP_NAME=grace
cf push $APP_NAME --no-start -b=go_buildpack
cf curl /v2/apps/`cf app $APP_NAME --guid` -X PUT -d '{"diego":true}'
cf start $APP_NAME
```

Dockerimage:
onsi/grace (based on ubuntu)
onsi/grace-busybox (based on busy-box)

To rebuild the dockerimage:
```bash
docker build -t="onsi/grace-busybox" .
docker push onsi/grace-busybox
```