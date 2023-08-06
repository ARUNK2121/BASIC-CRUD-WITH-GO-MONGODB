This is CRUD operations on MongoDB written in Golang. You can create, read, update, and delete users from the MongoDB instance using http requests.

## How to run?
First, start a MongoDB instance using docker:
```sh
docker run --name mongodb -d -p 27017:27017 mongo
```
Next, clone the repository:
```sh
git clone git@github.com:ARUNK2121/BASIC-CRUD-WITH-GO-MONGODB.git
```
Next, change the current directory to the repository:
```sh
cd go-mongo
```
Next, install the dependencies:
```sh
go get ./...
```
Finally, run the app on port `3000`:
```sh
go run .
```
Run using make file:
```sh
make run
```

## Endpoints:
```sh
GET    /users/:email
POST   /users
PUT    /users/:email
DELETE /users/:email
```

### Get User
This endpoint retrieves a user given the email.  
Send a `GET` request to `/users/:email`:
```sh
curl -X GET 'http://127.0.0.1:3000/users/arthurbishop120@gmail.com'
```
Response:
```sh
{
  "user": {
    "id": "<user_id>",
    "name": "Arun K",
    "email": "arthurbishop120@gmail.com",
    "password": "password"
  }
}
```
### Create User
This endpoint inserts a document in the `users` collection of the `users` database.  
Send a `POST` request to `/users`:
```sh
curl -X POST 'http://127.0.0.1:3000/users' -H "Content-Type: application/json" -d '{"name": "Arun K", "email": "arthurbishop120@gmail.com", "password": "password"}'
```
Response:  
```sh
{
  "user": {
    "id": "<user_id>",
    "name": "Arun K",
    "email": "arthurbishop120@gmail.com",
    "password": "password"
  }
}
```
### Update User
This endpoint updates the provided fields within the specified document filtered by email.  
Send a `PUT` request to `/users/:email`:
```sh
curl -X PUT 'http://127.0.0.1:3000/arthurbishop120@gmail.com' -H "Content-Type: application/json" -d '{"password": "password"}'
```
Response:
```sh
{
  "user": {
    "id": "<user_id>",
    "name": "Arun K",
    "email": "arthurbishop120@gmail.com",
    "password": "password"
  }
}
```

### Delete User
This endpoint deletes the user from database given the email.  
Send a `DELETE` request to `/users/:email`:
```sh
curl -X DELETE 'http://127.0.0.1:9080/users/arthurbishop120@gmail.com'
```
Response:
```sh
{}
```

### Errors
All of the endpoints return an error in json format with a proper http status code, if something goes wrong:
```sh
{
  "error": "user not found"
}
```