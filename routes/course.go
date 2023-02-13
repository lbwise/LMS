package routes

import (
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lbwise/LMS/db"
)

type Course struct {
	Name string `json:"name"`
	Price string `json:"price"`
}

type CourseFull struct {
	Course
	ID string `json:"id"`
	Description string `json:"description"`
	Created string `json:"created_on"`
}


func CourseRoutes(router *gin.RouterGroup) {
	router.GET("", checkLoggedIn, getCourses)
	router.GET("/:id", getCourseId)
}

func getCourses(c *gin.Context) {
	var (
		courses []Course
		name string
		price []uint8 
	) 
	const limit int = 50
	query := fmt.Sprintf(`SELECT name, price FROM courses ORDER BY random() LIMIT 10;`)
	rows, err := db.DB.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&name, &price)
		if err != nil {
			log.Fatal(err)
		}
		course := Course{Name: name, Price: string(price)} 	
		courses = append(courses, course)
	}
	c.JSON(200, courses)
}

func getCourseId(c *gin.Context) {
	var course CourseFull
	courseId, _ := c.Params.Get("id")
	query := fmt.Sprintf(`SELECT name, price, course_id, description, created_on FROM courses WHERE course_id='%s';`, courseId)
	err := db.DB.QueryRow(query).Scan(
		&course.Name,
		&course.Price,
		&course.ID,
		&course.Description,
		&course.Created,
	)
	if err != nil {
		panic(err.Error())
	}
	
	c.JSON(200, course)
}