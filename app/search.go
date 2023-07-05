package app

import (
	"dictionary-service/app/lib/search"
	"dictionary-service/app/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func (app *App) HandleSearch(ch chan *gin.Context) {
	var searchQuery models.SearchQuery

	c := <-ch
	start := time.Now()
	similarWords := []string{}
	if c.ShouldBind(&searchQuery) == nil {
		list, storedInCache := app.Cache.Get(searchQuery.Word)
		if storedInCache {
			fmt.Println("from cache")
			similarWords = list
		} else {
			fmt.Println("from search")
			similarWords = search.SearchWordPermutations(searchQuery.Word, *app.Trie)
			app.Cache.Set(searchQuery.Word, similarWords)
		}
	}
	elapsed := time.Since(start)
	app.Stats.AddRequest(int(elapsed.Nanoseconds()))
	resp := models.SearchResponse{Similar: similarWords}
	c.JSON(200, resp)
}
