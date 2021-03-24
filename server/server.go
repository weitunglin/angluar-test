package main

import (
	"fmt"
	"strings"
	"sort"
	"strconv"
	"log"
	"os"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Course struct {
	ID int `json:"id"`
	Description string `json:"description"`
	CourseListIcon string `json:"courseListIcon"`
	LongDescription string `json:"longDescription"`
	Category string `json:"category"`
	LessonsCount int `json:"lessonsCount"`
}

type Lesson struct {
	ID int `json:"id"`
	Description string `json:"description"`
	Duration string `json:"duration"`
	SeqNo int `json:"seqNo"`
	CourseID int `json:"courseId"`
}

func main() {
	router := gin.Default()

	router.GET("/api/courses", func(c *gin.Context) {
		var courses []Course
		byteBody, err := os.ReadFile("./course.json")
		if err != nil {
			log.Fatal(err.Error())
		}

		err = json.Unmarshal(byteBody, &courses)
		if err != nil {
			log.Fatal(err.Error())
		}

		c.JSON(200, gin.H{
			"payload": courses,
		})
	})

	router.GET("/api/courses/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Fatal(err.Error())
		}

		var courses []Course
		byteBody, err := os.ReadFile("./course.json")
		if err != nil {
			log.Fatal(err.Error())
		}

		err = json.Unmarshal(byteBody, &courses)
		if err != nil {
			log.Fatal(err.Error())
		}

		for _, v := range courses {
			if v.ID == id {
				c.JSON(200, v)
				return
			}
		}

		c.JSON(400, "Course Not Found!")
	})

	router.GET("/api/lessons", func(c *gin.Context) {
		courseID, err := strconv.Atoi(c.Query("courseId"))
		if err != nil {
			log.Fatal(err.Error())
		}

		filter := c.Query("filter")
		sortQrder := c.Query("sortQrder")
		pageNumber, err := strconv.Atoi(c.Query("pageNumber"))
		if err != nil {
			log.Fatal(err.Error())
		}

		pageSize, err := strconv.Atoi(c.Query("pageSize"))
		if err != nil {
			log.Fatal(err.Error())
		}

		var lessons []Lesson
		byteBody, err := os.ReadFile("./lesson.json")
		if err != nil {
			log.Fatal(err.Error())
		}

		err = json.Unmarshal(byteBody, &lessons)
		if err != nil {
			log.Fatal(err.Error())
		}

		n := 0
		for _, v := range lessons {
			if v.CourseID == courseID {
				lessons[n] = v
				n++
			}
		}

		lessons = lessons[:n]

		sort.SliceStable(lessons, func(i, j int) bool {
			if sortQrder == "asc" {
				return lessons[i].ID < lessons[j].ID
			}
			return lessons[i].ID >= lessons[j].ID
		})

		if filter != "" {
			fmt.Printf("%+v\n", lessons)
			var temp []Lesson
			for _, v := range lessons {
				if strings.Contains(v.Description, strings.TrimSpace(strings.ToLower(filter))) {
					temp = append(temp, v)
				}
			}
			lessons = temp
			fmt.Printf("%+v\n", lessons)
		}

		if pageNumber*pageSize < len(lessons) {
			lessons = lessons[pageNumber*pageSize:]
		}
		if len(lessons) > pageSize {
			lessons = lessons[:pageSize]
		}

		c.JSON(200, gin.H{
			"payload": lessons,
		})
	})

	router.Run(":9000")
}
