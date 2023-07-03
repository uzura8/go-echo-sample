package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Message struct {
    Name    string `json:"name"`
    Email   string `json:"email"`
    Message string `json:"message"`
}

type Response struct {
    Name    string `json:"name"`
    Email   string `json:"email"`
    Message string `json:message`
    Status  string `json:status`
}


func MessageRoutes(g *echo.Group) {
	g.POST("/", sendMessage)
}

func sendMessage(c echo.Context) error {
	m := new(Message)
	if err := c.Bind(m); err != nil {
		return err
	}

	r := new(Response)
	r.Name = m.Name
	r.Email = m.Email
	r.Message = m.Message
	r.Status = "success"
	return c.JSON(http.StatusOK, r)
}
