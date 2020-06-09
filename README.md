# Notes
Папки для БД
```
mkdir /opt/psqldata
mkdir /opt/mongodata
```
<<<<<<< HEAD
Запускаем БД
```
=======
>>>>>>> 62631564523545336278e05ca7f317f8277a117a
docker run -d \
    --name postgres \
    -e POSTGRES_USER=root \
    -e POSTGRES_PASSWORD=root \
    -e POSTGRES_DB=blog \
    -v /opt/psqldata:/var/lib/postgresql/data \
    -p 172.16.1.1:5432:5432 \
    postgres
```
```
docker run -d \
    --name mongo \
    -v /opt/mongodata:/data/configdb \
    -p 172.16.1.1:27017:27017 \
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
<<<<<<< HEAD
```
Можно тестить
=======
```
>>>>>>> 62631564523545336278e05ca7f317f8277a117a
