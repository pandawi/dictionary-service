package app

import (
	"dictionary-service/app/data/cache"
	database "dictionary-service/app/data/db"
	s "dictionary-service/app/lib/stats"
	"dictionary-service/app/lib/trie"
	"fmt"
	"time"
)

type App struct {
	Trie  *trie.Trie
	Stats s.Stats
	DB    database.DB
	Cache *cache.Cache
}

func InitializeApp() (*App, error) {
	db := database.NewDb()
	cache := cache.NewCache(2 * time.Minute)
	stats := s.NewStats()
	err := db.LoadData("./words_clean.txt")
	dbData := db.GetData()
	if err != nil {
		fmt.Println(err)
	}
	t := trie.NewTrie().FromSource(dbData)
	app := &App{
		Trie:  &t,
		Stats: stats,
		DB:    db,
		Cache: cache,
	}
	return app, nil
}
