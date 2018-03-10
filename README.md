#ROOM MICRO SERVICE
This Microservice to StreamBoard is responsible for the manage of all the rooms and their participants.
##Technologies
The microservice is made in C# language and .Net Core Framework. The database is PostgreSQL.

##Description
The microservice use the table Room, this have the next attributes:
* Id: the id of this room.
* NameRoom: A name for the room to be easy to find in a search.
* DescriptionRoom: A detail description, is not necessary write this.
* IdOwner: The Id of the User Creator of the room.
* Participants: It's a Identifications List of the active participants in the room.

##Developer
Joan Sebastian Contreras Peña.