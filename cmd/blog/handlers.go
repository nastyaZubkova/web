package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	Title           string
	FeaturedPosts   []featuredPostData
	MostResentPosts []mostResentPostData
}

type theRoadAheadPage struct {
	Title string
}

type featuredPostData struct {
	Title       string `db:"title"`
	Subtitle    string `db:"subtitle"`
	ImgModifier string `db:"image_url"`
	Author      string `db:"author"`
	AuthorImg   string `db:"author_url"`
	PublishDate string `db:"publish_date"`
}

type mostResentPostData struct {
	Title       string
	Subtitle    string
	Img         string
	Author      string
	AuthorImg   string
	PublishDate string
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		posts, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		ts, err := template.ParseFiles("../../pages/index.html") // Главная страница блога //изменила путь с ./pa... на этот, потому что дома отказывался работать
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
			return                                      // Не забываем завершить выполнение ф-ии
		}

		data := indexPage{
			Title:           "Escape",
			FeaturedPosts:   posts,
			MostResentPosts: mostResentPosts(),
		}

		err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		log.Println("Request completed successfully")
	}
}

func theRoadAhead(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("../../pages/the-road-ahead.html") //аналогично, либо ноут тупит, либр вайфай, но дома работает только с таким путём
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := theRoadAheadPage{
		Title: "the-road-ahead",
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}

}

func featuredPosts(db *sqlx.DB) ([]featuredPostData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			image_url,
			author,
			author_url,
			publish_date
		FROM
			post
		WHERE featured = 1
	` // Составляем SQL-запрос для получения записей для секции featured-posts

	var posts []featuredPostData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil
}

func mostResentPosts() []mostResentPostData {
	return []mostResentPostData{
		{
			Title:       "Still Standing Tall",
			Subtitle:    "Life begins at the end of your comfort zone.",
			Img:         "/static/images/still_tanding_tall.jpg",
			Author:      "William Wong",
			AuthorImg:   "/static/images/william_wong.jpg",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Sunny Side Up",
			Subtitle:    "No place is ever as bad as they tell you it’s going to be.",
			Img:         "/static/images/sunny_side_up.jpg",
			Author:      "Mat Vogels",
			AuthorImg:   "/static/images/mat_vogels.jpg",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Water Falls",
			Subtitle:    "We travel not to escape life, but for life not to escape us.",
			Img:         "/static/images/water_falls.jpg",
			Author:      "Mat Vogels",
			AuthorImg:   "/static/images/mat_vogels.jpg",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Through the Mist",
			Subtitle:    "Travel makes you see what a tiny place you occupy in the world.",
			Img:         "/static/images/throught_the_mist.jpg",
			Author:      "William Wong",
			AuthorImg:   "/static/images/william_wong.jpg",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Awaken Early",
			Subtitle:    "Not all those who wander are lost.",
			Img:         "/static/images/awaken_early.jpg",
			Author:      "Mat Vogels",
			AuthorImg:   "/static/images/mat_vogels.jpg",
			PublishDate: "9/25/2015",
		},
		{
			Title:       "Try it Always",
			Subtitle:    "The world is a book, and those who do not travel read only one page.",
			Img:         "/static/images/try_it_always.jpg",
			Author:      "Mat Vogels",
			AuthorImg:   "/static/images/mat_vogels.jpg",
			PublishDate: "9/25/2015",
		},
	}
}
