FROM rust:slim-buster as builder
COPY . /
WORKDIR /
RUN apt update && apt install -y default-libmysqlclient-dev pkg-config openssl libssl-dev
RUN cargo build --release

FROM debian:buster-slim
RUN apt update \
    && apt install -y openssl ca-certificates locales-all default-libmysqlclient-dev \
    && apt clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
COPY --from=builder /target/release/push-app /
ENTRYPOINT ["/push-app"]