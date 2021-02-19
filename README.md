# Double V Partners Test API
#### The following is a basic REST API built in Go (Golang) deployed with docker that supports CRUD operations of tickets.
## API starting command
``` console
docker-compose up
```
## Endpoints

### Create a ticket (POST)
``` 
http://localhost:8081/tickets/
```
#### Create Ticket Request Payload structure
```JSON
{
  "user": string 
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
```JSON
{
  "user": string 
  "status": boolean
}
``` 
#### Delete a ticket (DELETE)
``` 
http://localhost:8081/tickets/:id
```

