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
     # create the database + populate it with the data that is located in the /sqlc/placeholder.sql file

     go run main.go

docker

     # build the image
     docker build -t gofiber .

     # run and publish with the name of gofiber
     docker run --publish 5000:5000 --name gofiber gofiber

     # stop
     docker stop gofiber

     # remove
     docker image rm gofiber


## Migrations - dbmate

[dbmate](https://github.com/amacneil/dbmate) is used to run the migrations

     # initialize a new thing
     # before doing so, create a new DATABASE_URL that has the string with the url.
     dbmate new MIGRATION_NAME

     # after that, add all of the sql stuff into the file.

     # evertyhing under the "-- migrate:up" will be triggered when you run 
     dbmate up

     # evertyhing under the "-- migrate:down" will be triggered when you run 
     dbmate down

## Auto reload with [reflex](https://github.com/cespare/reflex) example 

     
     chmod +x reflex
     ./reflex -r '\.go' -s -- sh -c "go run main.go"


## Todos

- add new module -> Comment
- Check how Keyset Pagination works
- Implement the full flow for JWT auth.
- Check the gofiber examples repo, add ouath maybe

## Other

the `using_ent` branch has an example of ent ORM + docker-compose

<!-- 

# cd into the psql shell
sudo -u postgres psql -U postgres
\c test_db

 -->