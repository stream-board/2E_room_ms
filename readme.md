STEP BY STEP
1. How to install Go
https://medium.com/@patdhlk/how-to-install-go-1-9-1-on-ubuntu-16-04-ee64c073cd79

2. Install the dependencies
    2.1 The framework Gin
    $ go get github.com/gin-gonic/gin
    2.2 ORM
    $ go get github.com/jinzhu/gorm
    2.3 sqlite3
    $ go get github.com/mattn/go-sqlite3

3.
https://medium.com/@etiennerouzeaud/how-to-create-a-basic-restful-api-in-go-c8e032ba3181
https://gist.github.com/EtienneR/ed522e3d31bc69a9dec3335e639fcf60

4. docker
http://himarsh.org/build-restful-api-microservice-with-go/

ROOM MICRO SERVICE This Microservice to StreamBoard is responsible for the manage of all the rooms and their participants. ##Technologies The microservice is made in C# language and .Net Core Framework. The database is PostgreSQL.

Description The microservice use the table Room, this have the next attributes:

Id: the id of this room.
NameRoom: A name for the room to be easy to find in a search.
DescriptionRoom: A detail description, is not necessary write this.
IdOwner: The Id of the User Creator of the room.
Participants: It's a Identifications List of the active participants in the room.
Developer Joan Sebastian Contreras Peña.
