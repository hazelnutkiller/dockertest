FROM golang:latest AS buildStage

WORKDIR /src/docker_test

ADD . /src/docker_test/

ENV GO111MODUKE=on
ENV GOPROXY="https:goproxy.io"

RUN go build -o docker_test .


FROM alpine:latest
WORKDIR /dtest
COPY --from=buildStage /src/docker_test/ dtest/
EXPOSE 8080
ENTRYPOINT ./dtest