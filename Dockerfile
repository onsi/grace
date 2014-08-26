FROM busybox:ubuntu-14.04

ADD http://onsi-public.s3.amazonaws.com/grace.tar.gz /grace.tar.gz
RUN tar -zxf /grace.tar.gz && \
    chmod +x /grace && \
    rm /grace.tar.gz
