# assessment
Golang technical assessment.

Here I am attaching the explanation document: https://hackmd.io/@Artaghal/BkZFgd7zT

Further to execute the gin service, following are the steps

**Step-01**
Clone the repo to your local system.
```shell=
git clone https://github.com/wasimkhandot01/assessment.git
```

**Step-02**
Get the dependencies used and mentioned in go.mod file.
```shell=
go get -u ./...
```

**Step-03**
Build image from the dockerfile for postgres database. It includes setting user, database and password.

```shell=
docker build -d postgres-db-image .
```

**Step-04**
Start postgres-db container from the image.
```shell=
docker run -d --name prostgres-db -p 5432:5432
```

**Step-05**
Now you need to inspect the container IP which is needed as a host to connect our golang service with it.
```shell
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' my-postgres-container
```

**Final-Step**
Start your golang application
```shell=
go run main.go
```
