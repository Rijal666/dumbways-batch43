package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"personal-web/connection"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

type Blog struct {
	ID         int
	Name       string
	Sdate      time.Time
	Edate      time.Time
	Desc       string
	Technology []string
	Image      string
}

var blogs = []Blog{
	{
		Name: "RIZAL",
		// Sdate:      "12 jan 2020",
		// Edate:      "12 feb 2020",
		Desc:       "Dumbways asik asik asik jos",
		Technology: []string{"Node", "Reac", "Next", "type"},
	},
}

func main() {
	route := mux.NewRouter()
	connection.DatabaseConnecet()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/blog-Detail/{id}", blogDetail).Methods("GET")
	route.HandleFunc("/contact", kontak).Methods("GET")
	route.HandleFunc("/add-Project", formProject).Methods("GET")
	route.HandleFunc("/add-Project", addProject).Methods("POST")
	route.HandleFunc("/delete-Project/{index}", deleteProject).Methods("GET")
	route.HandleFunc("/update-Project/{index}", editProject).Methods("GET")
	route.HandleFunc("/update-Project/{index}", updateProject).Methods("POST")

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

	dataBlog, errQuery := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, description, technologies, image FROM public.tb_blog;")
	if errQuery != nil {
		fmt.Println("Message from patching : " + errQuery.Error())
		return
	}

	var result []Blog

	for dataBlog.Next() {
		var each = Blog{}

		err := dataBlog.Scan(&each.ID, &each.Name, &each.Sdate, &each.Edate, &each.Desc, &each.Technology, &each.Image)
		if err != nil {
			fmt.Println("Message : " + err.Error())
			return
		}
		result = append(result, each)
	}

	fmt.Println(result)

	resData := map[string]interface{}{
		"Blogs": result,
	}
	tempt.Execute(w, resData)
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
				ID:         id,
				Name:       data.Name,
				Sdate:      data.Sdate,
				Edate:      data.Edate,
				Desc:       data.Desc,
				Technology: data.Technology,
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
	// fmt.Println("name : " + r.PostForm.Get("name"))
	// fmt.Println("start date : " + r.PostForm.Get("sdate"))
	// fmt.Println("end date : " + r.PostForm.Get("edate"))
	// fmt.Println("description : " + r.PostForm.Get("desc"))
	// fmt.Println("checkbox1 : " + r.PostForm.Get("cb1"))
	// fmt.Println("checkbox2 : " + r.PostForm.Get("cb2"))
	// fmt.Println("checkbox3 : " + r.PostForm.Get("cb3"))
	// fmt.Println("checkbox4 : " + r.PostForm.Get("cb4"))
	name := r.PostForm.Get("name")
	// sdate := r.PostForm.Get("sdate")
	// edate := r.PostForm.Get("edate")
	desc := r.PostForm.Get("desc")
	technology := r.Form["Technology"]

	var newBlog = Blog{
		Name: name,
		// Sdate:      sdate,
		// Edate:      edate,
		Desc:       desc,
		Technology: technology,
	}
	blogs = append(blogs, newBlog)
	//fmt.Println(blogs)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	blogs = append(blogs[:index], blogs[index+1:]...)

	http.Redirect(w, r, "/", http.StatusFound)
}

func editProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("html/update-blog.html")

	if err != nil {
		w.Write([]byte("Message: " + err.Error()))
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var DataBlog = Blog{}

	for index, data := range blogs {
		if index == id {
			DataBlog = Blog{
				ID:         id,
				Name:       data.Name,
				Sdate:      data.Sdate,
				Edate:      data.Edate,
				Desc:       data.Desc,
				Technology: data.Technology,
			}
		}
	}

	EditBlog := map[string]interface{}{
		"Blog": DataBlog,
	}
	// fmt.Println(EditBlog)
	tmpt.Execute(w, EditBlog)
}

func updateProject(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	name := r.PostForm.Get("name")
	technology := r.Form["technology"]
	desc := r.PostForm.Get("desc")
	// sdate := r.PostForm.Get("sdate")
	// edate := r.PostForm.Get("edate")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	blogs[id].Name = name
	// blogs[id].Sdate = sdate
	// blogs[id].Edate = edate
	blogs[id].Desc = desc
	blogs[id].Technology = technology

	// fmt.Println(blogs)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
