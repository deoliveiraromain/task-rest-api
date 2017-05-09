# task-rest-api 

A rest api sample written in golang.
Data persistance with mongoDB.

Purpose : CRUD for Task object.

## How-to

You can run simply without docker, but you need a mongoDB instance running :

```bash
go run main.go
```

Or run api in docker containers for both DB and application :

```bash
docker-compose up -d
```

Warning : in the compose file, you can change TASK_API_PORT value but be sure to match with mapped docker port and ApiPort in config.json.

