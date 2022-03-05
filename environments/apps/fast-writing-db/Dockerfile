FROM mysql:latest

COPY init.d/ /docker-entrypoint-initdb.d
COPY my.cnf /etc/mysql/

EXPOSE 3306
