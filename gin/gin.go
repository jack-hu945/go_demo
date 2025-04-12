package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Student struct {
	Name   string
	Age    int
	Height float32
}

type Request struct {
	StudentId string `json:"student_id"`
}

// get student object from redis by student id
func GetStudentById(studentId string) Student {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	ctx := context.Background()
	stu := Student{}

	for filed, value := range client.HGetAll(ctx, studentId).Val() {
		switch filed {
		case "Name":
			stu.Name = value
		case "Age":
			age, err := strconv.Atoi(value)
			if err == nil {
				stu.Age = age
			}
		case "Height":
			height, err := strconv.ParseFloat(value, 32)
			if err == nil {
				stu.Height = float32(height)
			}
		}
	}

	return stu
}

func GetName(ctx *gin.Context) {
	param := ctx.Query("student_id")
	if len(param) == 0 {
		ctx.String(http.StatusBadRequest, "Please indidate student_id")
		return
	}
	stu := GetStudentById(param)
	ctx.String(http.StatusOK, stu.Name)
	return
}

func GetAge(ctx *gin.Context) {
	param := ctx.PostForm("student_id")
	if len(param) == 0 {
		ctx.String(http.StatusBadRequest, "Please indidate student_id")
		return
	}
	stu := GetStudentById(param)
	ctx.String(http.StatusOK, strconv.Itoa(stu.Age))
	return
}

func GetHight(ctx *gin.Context) {
	var param Request
	err := ctx.BindJSON(&param)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Please indidate student_id")
		return
	}
	stu := GetStudentById(param.StudentId)
	//ctx.String(http.StatusOK, strconv.FormatFloat(float64(stu.Height), 'f', 1, 64))
	ctx.JSON(http.StatusOK, stu)
	return
}

func main() {
	engine := gin.Default()
	engine.GET("/get_name", GetName) //set up the route
	engine.POST("/get_age", GetAge)
	engine.POST("/get_height", GetHight)

	err := engine.Run("0.0.0.0:2345")
	if err != nil {
		panic(err)
	}
}
