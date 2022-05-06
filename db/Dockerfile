FROM mysql

# Add a database
ENV MYSQL_DATABASE=ums_app

#Set root password
ENV MYSQL_ROOT_PASSWORD=root

# Add init data to database
COPY data.sql /docker-entrypoint-initdb.d/

EXPOSE 3306
