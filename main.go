package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/mkhairavir/http-service/model"
)

func main() {
	// 1. init data
	store := model.NewArticleStoreInMemory()

	e := echo.New()

	//2. endpoint rest
	e.GET("/articles", func(c echo.Context) error {
		//3. mengambill data dari in memory store
		articles := store.ArticleMap

		//4. return data
		return c.JSON(http.StatusOK, articles)
	})

	e.GET("/articles/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		articles := store.ArticleMap

		return c.JSON(http.StatusOK, articles[id-1])

	})

	e.POST("/articles", func(c echo.Context) error {
		//1. mengambil data dari form value
		title := c.FormValue("title")
		body := c.FormValue("body")

		//2. create article instances
		article, _ := model.CreateArticle(title, body)

		fmt.Println(article)

		//3. save ke store
		store.Save(article)

		return c.JSON(http.StatusOK, article)

	})

	e.PUT("/articles/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		title := c.FormValue("title")
		body := c.FormValue("body")

		store.Edit(id, title, body)

		return c.JSON(http.StatusOK, store.ArticleMap)
	})

	e.DELETE("/articles/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		store.Del(id)
		return c.JSON(http.StatusOK, store.ArticleMap)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
