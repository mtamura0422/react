FROM mysql:5.6.36

COPY conf.d/my.cnf /etc/mysql/conf.d/custom.cnf
RUN chmod +r /etc/mysql/conf.d/custom.cnf

ENV TZ Asia/Tokyo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV MYSQL_ROOT_PASSWORD mysql
ENV MYSQL_USER to
ENV MYSQL_PASSWORD Ac7PPsLG6H
ENV MYSQL_DATABASE to_development

ADD init_data/dump.sql.gz /docker-entrypoint-initdb.d/
RUN gzip -d /docker-entrypoint-initdb.d/dump.sql.gz

CMD ["mysqld", "--log-error=/var/lib/mysql/error.log"]

