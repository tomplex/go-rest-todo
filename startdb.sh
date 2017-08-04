docker stop todo_db; docker rm todo_db

CONTAINER=$(docker run --name todo_db -d -v $PWD/data:/var/lib/postgresql/data postgres)

IP=$(docker inspect --format "{{ .NetworkSettings.IPAddress }}" $CONTAINER)

echo "export PGHOST=$IP
export PGUSER=postgres
export PGPASSWORD=postgres
export PGDATABASE=postgres" > pg.env

echo "host=$IP dbname=postgres user=postgres password=postgres port=5432 sslmode=disable" > .connect


