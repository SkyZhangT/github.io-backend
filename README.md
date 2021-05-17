# github.io-backend

This is the backend of my personal gallery project. The server is supposed to maintain all gallery data and feed these data to front-end while requested.The frontend code is in another repo <link>https://github.com/SkyZhangT/skyzhangt.github.io</link>.

## Project Objectives

1. Design the core API services for the gallery (Finished)

   - Design API interfaces (Finished)
   - Design the database for posts (Finished)
   - Design the file system for images (Finished)

2. Learn Golang, MongoDB, and Docker (Finished)

   - Learn Gin (an advanced HTTP request library in Golang) (Finished)
   - Learn MongoDB controller in Golang (Finished)
   - Learn MongoDB queries (Finished)
   - Learn Docker compose usage (Finished)

3. Implement content database using MongoDB (In Progress)

   - Implement post collection (Finished)
   - Implement user collection (In Progress)
   - Implement comment collection (In Progress)

4. Implement API interface using Golang Gin (In Progress)
   - See REST API Services section

## REST API Services

### GET /post (Finished)

get next 10 posts sort by date

### POST /post (Finished)

add a new post

### PATCH /post

update an existing post

### DELETE /post/{id} (Finished)

delete a post

### GET /post/{id} (Finished)

get a specific post by id

### GET /image/{path to image} (Finished)

get a specific image file

### POST /login

create a user account

### GET /login

get token for this login session

### PATCH /login

change password
