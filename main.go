package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Rooms struct {
	IdRoom			int    `gorm:"AUTO_INCREMENT" form:"idroom" json:"idroom"`
	NameRoom 		string `gorm:"not null" form:"nameroom" json:"nameroom"`
	DescriptionRoom string `form:"descriptionroom" json:"descriptionroom"`
	IdOwner 		int 	`gorm:"not null" form:"idowner" json:"idowner"`
	Participants 	[]int 	`gorm:"type:int[]" form:"participants" json:"participants"`
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
	if !db.HasTable(&Rooms{}) {
		db.CreateTable(&Rooms{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Rooms{})
		room := Rooms{IdOwner:0,NameRoom:"default"}
		db.Create(&room)
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
		v1.PUT("/rooms/:id", UpdateRoom)
		v1.DELETE("/rooms/:idroom", DeleteRoom)
	}
	r.Run(":5000")
}

func PostRoom(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var roomFromBody Rooms
	c.Bind(&roomFromBody)

	//if exists the idroom(is a participant) or the nameroom(is a owner)
	if roomFromBody.NameRoom != "" || roomFromBody.IdRoom != 0{
		//its a creation of a new room
		if roomFromBody.NameRoom != ""{
			db.Create(&roomFromBody)
			c.JSON(201, roomFromBody)
		}else{
			//maybe is a participant, we need to check if the room exist
			var roomFromDB Rooms
			db.First(&roomFromDB, roomFromBody.IdRoom)
			//this room exist?
			if roomFromDB.NameRoom != ""{
				//Add participant
				roomFromDB.Participants = append(roomFromDB.Participants,roomFromBody.IdOwner)
				db.Save(&roomFromDB)
				c.JSON(200, gin.H{"success": roomFromDB})
			}else{
				//Bad request, the participant wants to join in a inexistent room
				c.JSON(404, gin.H{"error": "The room that you want to join doesn't exist"})
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

	var rooms []Rooms
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
	var room Rooms
	// SELECT * FROM users WHERE id = 1;
	db.First(&room, idroom)

	if room.IdRoom != 0 {
		// Display JSON result
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

	// Get id user
	idroom := c.Params.ByName("idroom")
	var room Rooms
	// SELECT * FROM users WHERE id = 1;
	db.First(&room, idroom)

	if room.IdRoom != 0 {
		// DELETE FROM users WHERE id = user.Id
		db.Delete(&room)
		// Display JSON result
		c.JSON(200, gin.H{"success": "Room #" + idroom + " deleted"})
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

	var a Rooms
	//fetch the values of the body
	c.Bind(&a)
	//fetch the values of the URL
	//id := c.Params.ByName("id")

	c.JSON(200, gin.H{"success": a})
	//c.JSON(200, a)


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
