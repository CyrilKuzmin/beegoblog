# Notes
mkdir /opt/psqldata
mkdir /opt/mongodata

docker run -d \
    --name postgres \
    -e POSTGRES_USER=root \
    -e POSTGRES_PASSWORD=root \
    -e POSTGRES_DB=blog \
    -v /opt/psqldata:/var/lib/postgresql/data \
    -p 5432:5432 \
    postgres
    
docker run -d \
    --name mongo \
    -v /opt/mongodata:/data/configdb \
    -p 27017:27017 \
    mongo
    