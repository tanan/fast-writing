FROM rust as builder

WORKDIR /home/rust
COPY . .
RUN cargo build --release

FROM alpine:latest
WORKDIR /push-app
COPY --from=builder /home/rust/target/release/push-app .
ENTRYPOINT [ "./push-app" ]