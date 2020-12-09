FROM BASEIMAGE
RUN apk --no-cache add ca-certificates bash

ARG ARCH

ADD server /usr/local/bin/server

ENTRYPOINT ["server"]