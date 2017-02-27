FROM golang:onbuild

ARG TASK_API_PORT

COPY config.json /go/bin

EXPOSE $TASK_API_PORT
