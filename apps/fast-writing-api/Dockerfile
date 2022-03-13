FROM golang:1.17 as buildImage

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ADD . /usr/src/app
WORKDIR /usr/src/app
RUN go build -o fast-writing-api main.go

# runtime image
FROM alpine

COPY --from=buildImage /usr/src/app/fast-writing-api /
COPY --from=buildImage /usr/src/app/config.yaml /
ENTRYPOINT ["/fast-writing-api"]