package github.com/massivefox/myuse

import (
	"crud/dbcon"
	"crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var infoList []models.CreateInput

func ReadInfo(c *gin.Context) {
	client, ctx, cancel := dbcon.GetConnection()
	collectionlist := client.Database("data").Collection("info")
	defer client.Disconnect(ctx)
	defer cancel()
	filter := bson.D{{}}
	cursor, _ := collectionlist.Find(ctx, filter)
	infoList = []models.CreateInput{}
	for cursor.Next(ctx) {
		input := models.CreateInput{}
		err := cursor.Decode(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": err.Error(),
			})
			return
		}
		infoList = append(infoList, input)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   infoList,
	})
}
func ReadOneInfo(c *gin.Context) {
	client, ctx, cancel := dbcon.GetConnection()
	collectionlist := client.Database("data").Collection("info")
	defer client.Disconnect(ctx)
	defer cancel()
	infoList := models.CreateInput{}
	id := c.Param("id")
	err := collectionlist.FindOne(ctx, bson.M{"id": id}).Decode(&infoList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   infoList,
		"id":     id,
	})
}
func CreateInfo(c *gin.Context) {
	client, ctx, cancel := dbcon.GetConnection()
	collectionlist := client.Database("data").Collection("info")
	input := models.CreateInput{}
	defer client.Disconnect(ctx)
	defer cancel()

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}
	if err, _ := collectionlist.InsertOne(ctx, input); err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "inserterr",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   input,
	})
}

func UpdateInfo(c *gin.Context) {
	client, ctx, cancel := dbcon.GetConnection()
	defer client.Disconnect(ctx)
	defer cancel()
	collectionlist := client.Database("data").Collection("info")

	input := models.UpdateInput{}

	id := c.Param("id")
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}
	collectionlist.FindOneAndUpdate(ctx, bson.M{"id": id}, bson.M{"$set": input})
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "user id is required",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "update_ok",
		"data":   input,
		"id":     id,
	})
}

func DeleteInfo(c *gin.Context) {
	client, ctx, cancel := dbcon.GetConnection()
	defer client.Disconnect(ctx)
	defer cancel()
	collectionlist := client.Database("data").Collection("info")
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "user id is required",
		})
		return
	}
	err := collectionlist.FindOneAndDelete(ctx, bson.M{"id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "no data in db",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"deleteid": id,
		"status":   "delete_ok",
	})
}
