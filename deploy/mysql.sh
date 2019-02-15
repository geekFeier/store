docker run --name store --net=host -v /data/store:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123 -e MYSQL_DATABASE=store -d mysql:5.7.25
