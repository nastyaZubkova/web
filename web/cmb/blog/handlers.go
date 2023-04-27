package main

import (
	"html/template"
	"log"
	"net/http"
)

type featuredPostData struct {
	Title       string
	Subtitle    string
	ImgModifier string
    Author      string
    AuthorImg   string
    PublishDate string
}

func featuredPosts() []featuredPostData {
	return []featuredPostData{
		{
			Title:       "The Road Ahead",
			Subtitle:    "The road ahead might be paved - it might not be.",
			Author:      "Mat Vogels",
			PublishDate: "September 25, 2015",
		},
		{
			Title:       "From Top Down",
			Subtitle:    "Once a year, go someplace you’ve never been before.",
			Author:      "William Wong",
			PublishDate: "September 25, 2015",
		},
	}
 }
 