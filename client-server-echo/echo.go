package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
)

type Note struct {
	Name    string `json:"name"`
	Surname string `json:"surname" `
	Text    string `json:"text"`
}

var NoteStorage = []Note{}

func main() {
	e := echo.New()
	e.GET("/hello", GetHello)
	e.POST("/save", SaveNote)
	e.GET("/list_all", ListAll)
	e.Logger.Fatal(e.Start("127.0.0.1:5050"))
}
func GetHello(c echo.Context) error {
	name := c.QueryParam("name")
	surname := c.QueryParam("surname")
	return c.String(200, "name "+name+"surname"+surname)
}
func SaveNote(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Println(err)
		return c.NoContent(500)
	}
	n := Note{}
	err = json.Unmarshal(body, &n)
	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	NoteStorage = append(NoteStorage, n)
	fmt.Println("Имя:", n.Name)
	fmt.Println("Фамилия:", n.Surname)
	fmt.Println("Заметка:", n.Text)

	return c.String(http.StatusOK, "Created")
}

func ListAll(c echo.Context) error {
	Data, err := json.Marshal(NoteStorage)
	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, json.RawMessage(Data))
}