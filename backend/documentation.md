### Steps

docker run -d -p 3306:3306 --name mysql-docker-container -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=banking -e MYSQL_USER=mysql mysql/mysql-server:latest

Ref: https://github.com/docker-library/mysql/issues/275

docker exec -it mysql-docker-container bash
mysql -u your*user -p
CREATE USER 'root'@'%' IDENTIFIED BY 'secret';
GRANT ALL PRIVILEGES ON *.\_ TO 'root'@'%' WITH GRANT OPTION;

On DBeaver
Driver settings: allowPublicKeyRetrieval=TRUE
