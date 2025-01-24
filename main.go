package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Person struct {
    Name string `form:"name"`
    Address string `form:"address"`
}

func getPersonHandler(context *gin.Context) {
    var person Person
    context.ShouldBindQuery(&person)
    context.JSON(
        http.StatusOK,
        gin.H {
            "name": person.Name,
            "address": person.Address,
        },
    )
}

func main() {
    engine := gin.Default()
    engine.GET("/person", getPersonHandler)
    engine.Run()
}
