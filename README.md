# Notes
mkdir /opt/psqldata
mkdir /opt/mongodata
```
docker run -d \
    --name postgres \
    -e POSTGRES_USER=root \
    -e POSTGRES_PASSWORD=root \
    -e POSTGRES_DB=blog \
    -v /opt/psqldata:/var/lib/postgresql/data \
    -p 5432:5432 \
    postgres
```
```
docker run -d \
    --name mongo \
    -v /opt/mongodata:/data/configdb \
    -p 27017:27017 \
    mongo
```

Таблица для сессий (в отличии от таблицы юзеров, которая через ORM), сама не создастся.
Заходим:
```
docker exec -it postgres psql -d blog
```
И создаем:
```
CREATE TABLE session (
session_key	char(64) NOT NULL,
session_data	bytea,
session_expiry	timestamp NOT NULL,
CONSTRAINT session_key PRIMARY KEY(session_key)
);
```
Проверяем:
```
blog=# \dt
        List of relations
 Schema |  Name   | Type  | Owner 
--------+---------+-------+-------
 public | session | table | root
 public | user    | table | root
(2 rows)
```