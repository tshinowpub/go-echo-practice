# What is this?
Go Echo practice is Go + [Echo FlameWork](https://github.com/labstack/echo) + [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) Samples.

## Docker
Move directory to ```path/to/docker```. After that run docker command.

```
$ docker-compose up -d
``` 

If your environment is successful, It is displayed as follows.

```
$ docker-compose ps
          Name                         Command               State           Ports
-------------------------------------------------------------------------------------------
echo-practice-migration     bash                             Up
echo-practice-rds           docker-entrypoint.sh mysql ...   Up      0.0.0.0:3306->3306/tcp
echo-practice-rds-manager   /run.sh phpmyadmin               Up      0.0.0.0:8081->80/tcp
```

## Migrations
This project is used [golang-migrate](https://github.com/golang-migrate/migrate) as a migration tool.

[CLI Document](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

### Run Migration
Run migration command. After into docker container.

- Up
```
$ migrate -path /migrations/echo -database "$MIGRATE_ECHO" up
```

- Down
```
$ migrate -path /migrations/echo -database "$MIGRATE_ECHO" down
```