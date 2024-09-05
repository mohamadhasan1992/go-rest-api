package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mohamadhasan1992/go-rest-api.git/models"
)

func PrintQueryParams(c *gin.Context) {
	var person models.PersonQuery
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}
	c.JSON(200, gin.H{"name": person.Name, "address": person.Address, "birthDay": person.Birthday})
}

func PrintUrlParams(c *gin.Context) {
	var person models.PersonUri
	if c.ShouldBindUri(&person) == nil {
		log.Println(person.Name)
		log.Println(person.ID)
	}
	c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
}

func PrintMultiPartForm(c *gin.Context) {
	var form models.LoginForm
	// in this case proper binding will be automatically selected
	err := c.ShouldBind(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "parsing error"})
		return
	}
	c.JSON(200, gin.H{"user": form.User, "password": form.Password})
}

func PrintQuery(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")

	c.JSON(200, gin.H{"id": id, "page": page, "name": name, "message": message})
}

func LogAsync(c *gin.Context) {
	// create copy to be used inside the goroutine
	cCp := c.Copy()
	go func() {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)

		// note that you are using the copied context "cCp", IMPORTANT
		log.Println("Done! in path " + cCp.Request.URL.Path)
	}()
}

func LogSync(c *gin.Context) {
	// simulate a long task with time.Sleep(). 5 seconds
	time.Sleep(5 * time.Second)

	// since we are NOT using a goroutine, we do not have to copy the context
	log.Println("Done! in path " + c.Request.URL.Path)
	c.JSON(200, gin.H{"message": "completed"})
}
