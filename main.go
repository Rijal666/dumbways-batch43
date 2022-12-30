package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

type Blog struct {
	Title string
	Desc  string
}

var blogs = []Blog{
	{
		Title: "RIZAL",
		Desc:  "Dumbways asik asik asik jos",
	},
}

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/blog-Detail/{id}", blogDetail).Methods("GET")
	route.HandleFunc("/contact", kontak).Methods("GET")
	route.HandleFunc("/add-Project", formProject).Methods("GET")
	route.HandleFunc("/add-Project", addProject).Methods("POST")

	fmt.Println("Server berjalan pada port 5000")
	http.ListenAndServe("localhost:5000", route)

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf.8")
	tempt, err := template.ParseFiles("html/index.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}
	// dataBlog := map[string]interface{}{
	// 	"Blogs": blogs,
	// }
	tempt.Execute(w, nil)
}

func blogDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf.8")
	tempt, err := template.ParseFiles("html/blog-detail.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return

	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var blogDetail = Blog{}

	for index, data := range blogs {
		if index == id {
			blogDetail = Blog{
				Title: data.Title,
				Desc:  data.Desc,
			}
		}
	}

	tempt.Execute(w, blogDetail)
}

func kontak(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf.8")
	tempt, err := template.ParseFiles("html/contact.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}
	tempt.Execute(w, nil)
}

func formProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset-utf.8")
	tempt, err := template.ParseFiles("html/blog.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}
	tempt.Execute(w, nil)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("name : " + r.PostForm.Get("name"))
	fmt.Println("start date : " + r.PostForm.Get("sdate"))
	fmt.Println("end date : " + r.PostForm.Get("edate"))
	fmt.Println("description : " + r.PostForm.Get("desc"))
	fmt.Println("checkbox1 : " + r.PostForm.Get("cb1"))
	fmt.Println("checkbox2 : " + r.PostForm.Get("cb2"))
	fmt.Println("checkbox3 : " + r.PostForm.Get("cb3"))
	fmt.Println("checkbox4 : " + r.PostForm.Get("cb4"))

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
