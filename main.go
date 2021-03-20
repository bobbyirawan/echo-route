package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type M map[string]interface{}

func main() {
	e := echo.New()

	e.GET("/", welcome)    // string
	e.GET("/index", index) // redirect
	e.GET("/html", html)   // html
	e.GET("/json", json)   // json

	e.GET("/page1", page1)         // query string test : http://localhost:9000/page1?name=grayson
	e.GET("/page2/:name", page2)   // url path param test : http://localhost:9000/page2/grayson
	e.GET("/page3/:name/*", page3) // http://localhost:9000/page3/tim/need/some/sleep

	e.POST("/page4", page4) // form-data

	e.Static("/static", "assets")

	e.Logger.Fatal(e.Start(":9000"))
}

func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func index(c echo.Context) error {
	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func html(c echo.Context) error {
	data := "hello form html"
	return c.HTML(http.StatusOK, data)
}

func json(c echo.Context) error {
	data := M{"message": "hello", "counter": 2}
	return c.JSON(http.StatusOK, data)
}

func page1(c echo.Context) error {
	name := c.QueryParam("name")
	data := fmt.Sprintf("Hello %s", name)
	return c.String(http.StatusOK, data)
}

func page2(c echo.Context) error {
	name := c.Param("name")
	data := fmt.Sprintf("Hello %s", name)
	return c.String(http.StatusOK, data)
}

func page3(c echo.Context) error {
	name := c.Param("name")
	message := c.Param("*")
	data := fmt.Sprintf("Hello %s i have message for you: %s", name, message)

	return c.String(http.StatusOK, data)
}

func page4(c echo.Context) error {
	name := c.FormValue("name")
	message := c.FormValue("message")

	data := fmt.Sprintf(
		"Hello %s, I have message for you: %s",
		name,
		strings.Replace(message, "/", "", 1),
	)

	return c.String(http.StatusOK, data)
}
