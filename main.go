package main

import (
	"context"
	"day-9/connection"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{}{
	"Title": "Personal Web",
}

type Blog struct {
	Id          int
	Title       string
	Post_date   time.Time
	Format_date string
	Author      string
	Content     string
}

var Blogs = []Blog{
	// {
	// 	Title:     "Pasar Coding Di Indonesia Dinilai Masih Menjanjikan",
	// 	Post_date: "20 October 2022 22:30 WIB",
	// 	Author:    "Abel Dustin",
	// 	Content:   "Test",
	// },
	// {
	// 	Title:     "Pasar Ikan Di Indonesia Dinilai Masih Menjanjikan",
	// 	Post_date: "20 October 2022 22:30 WIB",
	// 	Author:    "Abel Dustin",
	// 	Content:   "Test",
	// },
}

func main() {
	route := mux.NewRouter()

	connection.DatabaseConnect()

	// route path folder untuk public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	//routing
	route.HandleFunc("/hello", helloWorld).Methods("GET")
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/blog", blog).Methods("GET")
	route.HandleFunc("/blog-detail/{index}", blogDetail).Methods("GET")
	route.HandleFunc("/form-blog", formAddBlog).Methods("GET")
	route.HandleFunc("/add-blog", addBlog).Methods("POST")
	route.HandleFunc("/delete-blog/{index}", deleteBlog).Methods("GET")

	fmt.Println("Server running on port 5000")
	http.ListenAndServe("localhost:5000", route)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/contact.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func blog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// fmt.Println(Blogs)

	var tmpl, err = template.ParseFiles("views/blog.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	// var query = "SELECT id, title, content FROM tb_blog"

	rows, _ := connection.Conn.Query(context.Background(), "SELECT id, title, content FROM tb_blog")

	var result []Blog // array data

	for rows.Next() {
		var each = Blog{} // manggil struct

		err := rows.Scan(&each.Id, &each.Title, &each.Content)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		each.Author = "Abel Dustin"
		// each.Format_date = each.Post_date.Format("2 January 2006")

		result = append(result, each)
	}

	fmt.Println(result)

	respData := map[string]interface{}{
		"Blogs": result,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)
}

func blogDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/blog-detail.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	var BlogDetail = Blog{}

	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	for i, data := range Blogs {
		if index == i {
			BlogDetail = Blog{
				Title:     data.Title,
				Content:   data.Content,
				Post_date: data.Post_date,
				Author:    data.Author,
			}
		}
	}

	data := map[string]interface{}{
		"Blog": BlogDetail,
	}

	// fmt.Println(data)

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, data)
}

func formAddBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/add-blog.html")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func addBlog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Title : " + r.PostForm.Get("inputTitle")) // value berdasarkan dari tag input name
	fmt.Println("Content : " + r.PostForm.Get("inputContent"))

	var title = r.PostForm.Get("inputTitle")
	var content = r.PostForm.Get("inputContent")

	var newBlog = Blog{
		Title:   title,
		Content: content,
		Author:  "Abel Dustin",
	}

	//Blogs.push(newBlog)
	Blogs = append(Blogs, newBlog)
	// fmt.Println(Blogs)

	http.Redirect(w, r, "/blog", http.StatusMovedPermanently)
}

func deleteBlog(w http.ResponseWriter, r *http.Request) {
	index, _ := strconv.Atoi(mux.Vars(r)["index"])
	fmt.Println(index)

	Blogs = append(Blogs[:index], Blogs[index+1:]...)
	fmt.Println(Blogs)

	http.Redirect(w, r, "/blog", http.StatusFound)
}
