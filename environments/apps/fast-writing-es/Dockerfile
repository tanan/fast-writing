FROM docker.elastic.co/elasticsearch/elasticsearch:8.1.1

RUN bin/elasticsearch-plugin install --batch analysis-icu
RUN bin/elasticsearch-plugin install --batch analysis-kuromoji

COPY init.sh /usr/share/elasticsearch/init.sh
COPY fast-writing.json /usr/share/elasticsearch/fast-writing.json

ENTRYPOINT ["bash", "/usr/share/elasticsearch/init.sh"]