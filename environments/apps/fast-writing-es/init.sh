#!/bin/bash
postprocess() {
	  until curl -s -q --ipv4 "localhost:9200/_cluster/health?wait_for_status=green&timeout=60s&pretty"
	  do
      >&2 echo "Elasticsearch is unavailable - sleeping"
      sleep 1s
	  done
	  curl --ipv4 -XPUT -u elastic:changeme localhost:9200/fast-writing?pretty -H 'Content-Type: application/json' -d @fast-writing.json
}
postprocess &
/usr/local/bin/docker-entrypoint.sh $@