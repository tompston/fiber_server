# Gofiber REST API example with SQLC

- Two api endpoints -> user + post.

- Working pagitation with url query params + limit - offset
- Endpoints documentend in Postman.
- Validation for the sent json objects.
- Predefined responses
- User registration with hashed passwords + JWT cookies on login

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

## Notes

the using_ent branch has an example of ent + docker-compose

## Todos

- setting up migrations with one of the packages mentioned in sqlc docs.

- add constraint to username to check if it doesn't have charaters like '/'. Check regex stuff.
- add new module -> Comment
- Check if you can convert the struct validation blocks into a single function that can also pass down any structs.
- add email field for user?
- Check how Keyset Pagination works
- Implement the full flow for JWT auth.
- Check the gofiber examples repo, add ouath maybe
- Move pagitation stuff to a new /utils module?
- Test the swagger [thing](https://github.com/arsmn/fiber-swagger)
