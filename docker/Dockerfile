FROM wrfly/golang-template:build-base
COPY . /src
RUN cd /src && \
    make test && \
    make build && \
    mv bin/ /artifacts

FROM alpine
ARG BINARY=golang-template
RUN apk add --update ca-certificates
COPY --from=0 /artifacts/${BINARY} /usr/local/bin/
COPY config.yml /etc/binary-config.yml
VOLUME [ "/data" ]
CMD [ "binary", "-c", "/etc/binary-config.yml" ]
