You'll need to set the `RYTSAR_DB_NAME`, `RYTSAR_DB_USER`, `RYTSAR_DB_PASS`, `RYTSAR_DB_PORT`, and `RYTSAR_DB_HOST` environment variables in order for database connections to function.

Build the container with `docker build -t rytsar-build . && docker run -e RYTSAR_DB_HOST=host -e RYTSAR_DB_PORT=1234 -e RYTSAR_DB_NAME=db -e RYTSAR_DB_PASS=pass -e RYTSAR_DB_USER=user -p 15000:8000 --link sql:sql -it -d --name rytsar rytsar-build`. Alternatively, you can use the docker-compose file and `docker-compose up -d`.
