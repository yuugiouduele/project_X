package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ginrouter() {
	router := gin.Default()
	router.GET("http://", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"massage": "pong",
		})
	})
	router.Run()
}

func sqlhttpfusiion(dbdriver string, dsn string) (*sql.DB, error) {
	http.Get("https://localhost:8080/?param=")
	db, err := sql.Open(dbdriver, dsn)
	if err != nil {
		log.Fatal("error_mount")
		return nil, err
	}
	fmt.Println("Hello")
	return db, err
}

type Todo struct {
	id       int64
	name     string
	practice string
	runwork  bool
}

func main() {
	Driver := "postgre"
	dsn := "host=127.0.0.1 port=5608 user=user password=password dbname=dbname sslmode=disable"
	db, err := sqlhttpfusiion(Driver, dsn)
	db.Query("SELECT * FROM zoo WHERE ")
	db.QueryRow("SELECT * FROM doo WHERE ")

	for i := 0; i < 1000; i++ {
		go func() {
			if i%2 != 0 {
				var number string = strconv.Itoa(i)
				db.Exec("SELECT *FROM bee WHERE " + number)
			} else if i%7 == 0 {
				fmt.Println("prime number!!☆")
			}
		}()
	}
	defer db.Close()

	if err != nil {
		log.Fatal("logerror")
		ginrouter()
	}
	cliinit()
	crypto(12)
}
