docker run -dt \
--name postgres \
-e POSTGRES_USER=postgres \
-e POSTGRES_PASSWORD=password \
-e POSTGRES_DB=my_db \
-p 5432:5432 \
postgres:14
