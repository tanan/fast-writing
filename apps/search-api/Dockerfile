FROM golang:1.17 as buildImage

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ADD . /usr/src/app
WORKDIR /usr/src/app
RUN go build -o search-api main.go

# runtime image
FROM alpine

COPY --from=buildImage /usr/src/app/search-api /
COPY --from=buildImage /usr/src/app/config.yaml /
ENTRYPOINT ["/search-api"]
