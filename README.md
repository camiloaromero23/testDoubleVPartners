# Double V Partners Test API
#### The following is a basic REST API built in Go (Golang) deployed with docker that supports CRUD operations of tickets.

## Requirements
This installation requires the previous installation of docker V18.06.0+ and docker-compose 

## API starting command
In order to start the API, run the following command in the root of this repository:
``` console
docker-compose up
```
## Endpoints

### Create a ticket (POST)
``` 
http://localhost:8081/tickets/
```
#### Create Ticket Request Payload structure
```
{
  "user": string, 
  "status": boolean
}
``` 
### Get a single ticket (GET)
``` 
http://localhost:8081/tickets/:id
```
### Get all tickets (GET)
``` 
http://localhost:8081/tickets/
```
### Update a ticket (PUT)
``` 
http://localhost:8081/tickets/:id
```
#### Update Ticket Request Payload structure
```
{
  "user": string, 
  "status": boolean
}
``` 
#### Delete a ticket (DELETE)
``` 
http://localhost:8081/tickets/:id
```

