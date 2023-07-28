FROM alpine

COPY ./go-HoneyPot /
COPY ./config.json /

RUN apk --update add libstdc++ gcompat

CMD ["/go-HoneyPot"]