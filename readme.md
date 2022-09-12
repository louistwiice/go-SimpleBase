# Documentation

## Step 1: Set env file
|Env name| Description|Example|
|---|---|---|
|SERVER_PORT|Application running port|:9000|
|MYSQL_ROOT_PASSWORD| Mysql root passwod||
|MYSQL_DATABASE| Mysql Database||
|MYSQL_USER| Mysql User||
|MYSQL_PASSWORD| Mysql password||
|MYSQL_HOST| Server IP host |localhost|

Other environment variable will go there in this file


## Step 2: Start mysql container

``` text
docker-compose up -d
```


## Step 3: Preparing and installing migrations golang-migrate

### Step 3.1: Installation
First you should install golang-migrate with:

**On Mac**
``` text
brew install golang-migrate
```

**On Linux**
``` text
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
sudo apt-get update
sudo apt-get install migrate
```

**On Windows**
```text
go get -u -d github.com/golang-migrate/migrate/cmd/migrate
cd $GOPATH/src/github.com/golang-migrate/migrate/cmd/migrate
go env GOPATH
go install .
```

Check if everything is fine by
```text
migrate -v
```

### Step 3.2: Usage

**Create empty migration files**
```text
migrate create -ext sql -dir ./migrations -seq <<file_name>>
```

**Run migrations**
- for Mysql
```text
migrate -path ./migrations -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:3306)/${MYSQL_DATABASE}" -verbose up
```

- For Postgresl
```text
migrate -path ./migrations -database "postgres://${MYSQL_USER}:${MYSQL_PASSWORD}@${MYSQL_HOST}:5432/${MYSQL_DATABASE}" -verbose up
```

**Revert migrations**
- for Mysql
```text
migrate -path ./migrations -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:3306)/${MYSQL_DATABASE}" -verbose down
```

- For Postgresl
```text
migrate -path ./migrations -database "postgres://${MYSQL_USER}:${MYSQL_PASSWORD}@${MYSQL_HOST}:5432/${MYSQL_DATABASE}" -verbose down
```

## Step 4: Start application
```text
go run api/main.go
```
