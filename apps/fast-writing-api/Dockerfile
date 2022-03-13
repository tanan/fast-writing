FROM golang:1.17 as buildImage

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ADD apps/fast-writing-api /usr/src/app/fast-writing-api
ADD pkg/pb /usr/src/pkg/pb
WORKDIR /usr/src/app/fast-writing-api
RUN go mod tidy -compat=1.17 && go build -o fast-writing-api cmd/api/main.go

# runtime image
FROM alpine

COPY --from=buildImage /usr/src/app/fast-writing-api/fast-writing-api /
COPY --from=buildImage /usr/src/app/fast-writing-api/config.yaml /
ENTRYPOINT ["/fast-writing-api"]
