docker rm -f envoy-proxy
docker run --name "envoy-proxy" -d \
      -p 9901:9901 \
      -p 10000:10000 \
      -v $(pwd)/envoy.yaml:/etc/envoy/envoy.yaml \
      -v $(pwd)/logs:/logs \
      envoyproxy/envoy:v1.21-latest \
          -c /etc/envoy/envoy.yaml