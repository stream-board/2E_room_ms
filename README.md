# ROOM MICRO SERVICE
This Microservice to StreamBoard is responsible for the manage of all the rooms and their participants.
## Technologies
The microservice is made in C# language and .Net Core 2.0 Framework. The database is PostgreSQL.

## Description
The microservice use the table Room, this have the next attributes:
* Id: the id of this room.
* NameRoom: A name for the room to be easy to find in a search.
* DescriptionRoom: A detail description, is not necessary write this.
* IdOwner: The Id of the User Creator of the room.
* Participants: It's a Identifications List of the active participants in the room.

## Developer
Joan Sebastian Contreras Peña.

### Note:
Exist a conflict with the database with rancher because doesn't exist the connection with the microservice and the database. I'm trying to fix this... 

```
    ~/Documents/room_ms$ docker-compose up
    Starting roomms_rooms_db_1 ... 
    Recreating roomms_rooms_ms_1 ... done
    Attaching to roomms_rooms_db_1, roomms_rooms_ms_1
    rooms_db_1  | 2018-03-19 05:12:21.281 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
    rooms_db_1  | 2018-03-19 05:12:21.281 UTC [1] LOG:  listening on IPv6 address "::", port 5432
    rooms_ms_1  | warn: Microsoft.AspNetCore.DataProtection.KeyManagement.XmlKeyManager[35]
    rooms_ms_1  |       No XML encryptor configured. Key {c7c3f612-32d0-454c-9501-e51dc159e9b4} may be persisted to storage in unencrypted form.
    rooms_ms_1  | Hosting environment: Production
    rooms_db_1  | 2018-03-19 05:12:21.703 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
    rooms_ms_1  | Content root path: /rooms_ms
    rooms_ms_1  | Now listening on: http://localhost:4040
    rooms_db_1  | 2018-03-19 05:12:23.579 UTC [17] LOG:  database system was shut down at 2018-03-19 05:08:59 UTC


```

