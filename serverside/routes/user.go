package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
    Name  string `json:"name" validate:"required,max=10"`
    Email string `json:"email" validate:"required,email"`
}

type Error struct {
    Error string `json:"error"`
}


func UserRoutes(g *echo.Group) {
	g.POST("", saveUser)
	g.GET("/:id", getUser)
	g.GET("/:id/detail", getUserDetail)
	g.POST("/:id/detail", saveUserDetail)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "User ID: "+id)
}

func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	// Execute validate
	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, &Error{Error: err.Error()})
    }
	return c.JSON(http.StatusOK, u)
}

func getUserDetail(c echo.Context) error {
	id := c.Param("id")
    sex := c.QueryParam("sex")
	age := c.QueryParam("age")
	res := fmt.Sprintf("ID:%s のユーザの性別は%s、年齢は%sです。", id, sex, age)
	return c.String(http.StatusOK, res)
}

func saveUserDetail(c echo.Context) error {
	id := c.Param("id")
	sex := c.FormValue("sex")
	age := c.FormValue("age")
	res := fmt.Sprintf("ID:%s のユーザの属性として、性別:%s, 年齢:%sを登録しました", id, sex, age)
	return c.String(http.StatusOK, res)
}
