# ROOM MICRO SERVICE #
This Microservice to StreamBoard is responsible for the manage of all the rooms and their participants. 
## Technologies: ##
The microservice is made in Go language and Gin Framework. The database is SQLite3.

## Description: ##
The microservice use the table Room and the table Participant, this have the next attributes:

### Room ###
* IdRoom: The id of this room.
* NameRoom: A name for the room to be easy to find in a search.
* DescriptionRoom: A detail description, is not necessary write this.
* IdOwner: The Id of the User Creator of the room.
* Participants: It's a relation Has-Many with the Participant Table.

### Participant ###
* Id : The id generated.
* IdRoom: The id room.
* IdParticipant: The id of the user.

# STEP BY STEP TO DEPLOY LOCAL #
1. How to install Go
https://medium.com/@patdhlk/how-to-install-go-1-9-1-on-ubuntu-16-04-ee64c073cd79

2. Install the dependencies
    1. The framework Gin
    ```$ go get github.com/gin-gonic/gin```
    2. ORM
    ```$ go get github.com/jinzhu/gorm```
    3. sqlite3
    ```$ go get github.com/mattn/go-sqlite3```

# TO DEPLOY IN DOCKER #
We suppose that you have a node in docker.
```$ docker-compose build```
```$ docker-compose up```

Developer Joan Sebastian Contreras Peña - jscontrerasp@unal.edu.co