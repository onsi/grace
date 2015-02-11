FROM busybox:ubuntu-14.04

ENV CUSTOM_GRACE_ENV my-docker-configured-env
CMD ["-chatty"]
ENTRYPOINT ["/grace"]

WORKDIR /tmp

EXPOSE 8080 9999

COPY grace /grace
RUN chmod a+x /grace