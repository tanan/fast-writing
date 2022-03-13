FROM golang:1.17 as buildImage

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ADD apps/search-api /usr/src/app/search-api
ADD pkg/pb /usr/src/pkg/pb
WORKDIR /usr/src/app/search-api
RUN go mod tidy && go build -o search-api cmd/api/main.go

# runtime image
FROM alpine

COPY --from=buildImage /usr/src/app/search-api/search-api /
COPY --from=buildImage /usr/src/app/search-api/config.yaml /
ENTRYPOINT ["/search-api"]
