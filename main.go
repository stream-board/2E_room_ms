package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Room struct {
	IdRoom			int    `gorm:"primary_key" form:"idroom" json:"idroom"`
	NameRoom 		string `gorm:"not null" form:"nameroom" json:"nameroom"`
	DescriptionRoom string `form:"descriptionroom" json:"descriptionroom"`
	IdOwner 		int 	`gorm:"not null" form:"idowner" json:"idowner"`
	//Participants 	[]int 	`gorm:"type:int[]" form:"participants" json:"participants"`
	Participants []Participant 
}

type Participant struct{
	Id 				int		`gorm:"primary_key" form:"id" json:"id"`
	IdRoom			int		`gorm:"not null" form:"idroom" json:"idroom"`
	IdParticipant 	int		`gorm:"not null" form:"idparticipant" json:"idparticipant"`
}

func rem(s []int, i int) []int {
	var saux []int
	  for _, v := range s {
		if v == i{
		  continue
		}else{
		  saux = append(saux,v)
		}
	  }
	  return saux
  }

func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&Room{}) {
		db.CreateTable(&Room{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Room{})
	}
	if !db.HasTable(&Participant{}) {
		db.CreateTable(&Participant{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Participant{})
	}

	return db
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("")
	{
		v1.POST("/rooms", PostRoom)
		v1.GET("/rooms", GetRooms)
		v1.GET("/rooms/:idroom", GetRoom)
		v1.PUT("/rooms/", UpdateRoom)
		v1.DELETE("/rooms/:idroom", DeleteRoom)
	}
	r.Run(":5000")
}

func PostRoom(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var roomFromBody Room
	c.Bind(&roomFromBody)

	//if exists the idroom(is a participant) or the nameroom(is a owner)
	if roomFromBody.NameRoom != "" || roomFromBody.IdRoom != 0{
		//its a creation of a new room
		if roomFromBody.NameRoom != ""{
			db.Create(&roomFromBody)
			c.JSON(201, gin.H{"success": roomFromBody})
		}else{
			//maybe is a participant, we need to check if the room exist
			var roomFromDB Room
			db.Where("id_room = ?", roomFromBody.IdRoom ).First(&roomFromDB)
			//this room exist?
			if roomFromDB.NameRoom != ""{//exist this room in database
				//create participant in Participant Table
				var newPart = Participant{ IdRoom: roomFromDB.IdRoom, IdParticipant: roomFromBody.IdOwner }
				db.Create(&newPart)
				var participants []Participant
				db.Where("id_room = ?", roomFromDB.IdRoom).Find(&participants)
				roomFromDB.Participants = participants
				db.Save(roomFromDB)
				c.JSON(200, gin.H{"success": roomFromDB})
			}else{
				//Bad request, the participant wants to join in a inexistent room
				c.JSON(404, gin.H{"error": "doesn't exist the room"})
			}
			
		}
	} else {
		// Display error
		c.JSON(400, gin.H{"error": "Bad Request"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:5000/rooms
}

func GetRooms(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	var rooms []Room
	// SELECT * FROM users
	db.Find(&rooms)

	// Display JSON result
	c.JSON(200, rooms)

	// curl -i http://localhost:5000/rooms
}

func GetRoom(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	idroom := c.Params.ByName("idroom")
	var room Room
	// SELECT * FROM users WHERE id = 1;
	db.Where("id_room = ?",idroom).First(&room)

	if room.IdRoom != 0 {
		// Display JSON result
		var participants []Participant
		db.Where("id_room = ?",idroom).Find(&participants)
		room.Participants = participants
		db.Save(&room)
		c.JSON(200, room)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Room not found"})
	}

	// curl -i http://localhost:5000/rooms/1
}

func DeleteRoom(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	// Get id user from URL
	idroom := c.Params.ByName("idroom")
	//get room from body
	var roomFromBody Room
	c.Bind(&roomFromBody)


	var roomFromDB Room
	// SELECT * FROM users WHERE id = 1;
	db.Where("id_room = ?",idroom).First(&roomFromDB)


	if roomFromDB.NameRoom != "" {
		//is the owner of the room
		if roomFromDB.IdOwner == roomFromBody.IdOwner{
			//get all participants of this room
			var participants []Participant
			db.Where("id_room = ?",idroom).Find(&participants)
			//Delete all participants
			db.Delete(&participants)
			//Delete the room of database
			db.Delete(&roomFromDB)
			roomFromDB.Participants = participants
			// Display JSON result
			c.JSON(200, gin.H{"success": roomFromDB})
		}else{
			//is a participant
			var participant2Delete []Participant
			db.Where("id_participant = ?",roomFromBody.IdOwner).Find(&participant2Delete)
			db.Where("id_participant = ?",roomFromBody.IdOwner).Delete(&Participant{})
			//db.Delete(&participant2Delete)
			c.JSON(200, gin.H{"success": participant2Delete})
		}
		

	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Room not found"})
	}

	// curl -i -X DELETE http://localhost:5000/rooms/1
}



func UpdateRoom(c *gin.Context) {

	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	var participantes []Participant
	db.Find(&participantes)
	c.JSON(200, participantes)
	/*
	// Get id user
	id := c.Params.ByName("id")
	var user Users
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.Firstname != "" && user.Lastname != "" {

		if user.Id != 0 {
			var newUser Users
			c.Bind(&newUser)

			result := Users{
				Id:        user.Id,
				Firstname: newUser.Firstname,
				Lastname:  newUser.Lastname,
			}

			// UPDATE users SET firstname='newUser.Firstname', lastname='newUser.Lastname' WHERE id = user.Id;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "User not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/users/1
	*/
}



func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
