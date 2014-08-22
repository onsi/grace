FROM busybox:latest

ADD http://onsi-public.s3.amazonaws.com/grace.tar.gz /grace.tar.gz
RUN gunzip /grace.tar.gz
RUN tar -xf /grace.tar
RUN chmod +x /grace
RUN rm /grace.tar
