package controllers

import ("fmt"
		"net/http"
		"encoding/json")

type Article struct {
	Title string `json:"title"`
	Desc string  `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article


func AllArticles(w http.ResponseWriter,r *http.Request) {
     articles := Articles{
		Article{Title:"Test Title",Desc:"Test Description",Content:"Display Article:1"},
	}
	fmt.Println("GET ALL ARTICLES")
	json.NewEncoder(w).Encode(articles)
}