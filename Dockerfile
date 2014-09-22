FROM busybox:ubuntu-14.04

ENV CUSTOM_GRACE_ENV my-docker-configured-env
ADD http://onsi-public.s3.amazonaws.com/grace.tar.gz /grace.tar.gz
CMD ["-chatty"]
ENTRYPOINT ["/grace"]
RUN tar -zxf /grace.tar.gz && \
    chmod +x /grace && \
    rm /grace.tar.gz
