# Gofiber REST API example with SQLC, JWT, Pagitation and User Registration

- Two api endpoints -> user + post
- User registration with hashed passwords + JWT cookies on login
- Working pagitation with url query params + limit - offset
- Endpoints documentend in Postman.
- Pluggable validation for the sent json objects.
- Predefined responses

## How to run

locally

     # change the .env variables
     # install dbmate and run migrations with -> dbmate up
     go run main.go

docker

     docker build -t gofiber .                                   # build the image
     docker run --publish 5000:5000 --name gofiber gofiber       # run and publish with the name of gofiber
     docker stop gofiber                                         # stop
     docker image rm gofiber                                     # remove

## Migrations - dbmate

[dbmate](https://github.com/amacneil/dbmate) is used to run the migrations

     # follow the dbmate info and initialize a new thing
     # before doing so, create a new DATABASE_URL variable in .env that has the string with the url
     dbmate new MIGRATION_NAME

     # after that, add all of the sql stuff into the file.

     # evertyhing under the "-- migrate:up" will be triggered when you run
     dbmate up

     # evertyhing under the "-- migrate:down" will be triggered when you run
     dbmate down

## Auto reload with [reflex](https://github.com/cespare/reflex) example

```bash
# can install with gobinaries
curl -sf https://gobinaries.com/cespare/reflex | sh
reflex -r '\.go' -s -- sh -c "go run main.go"
```

## Todos

- add new module -> Comment
- Check how Keyset Pagination works
- Implement the full flow for JWT auth.
- Check the gofiber examples repo, add ouath maybe

## Other

the `using_ent` branch has an example of ent ORM + docker-compose

<!--

export PATH=$PATH:/usr/local/go/bin

# cd into the psql shell
sudo -u postgres psql -U postgres
\c test_db


 -->
