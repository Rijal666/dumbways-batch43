package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"personal-web/connection"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type metadata struct {
	Nama      string
	Islogin   bool
	Username  string
	FlashData string
}

var Data = metadata{
	Nama: "Personal Web",
}

type Blog struct {
	ID         int
	Nama       string
	Author     string
	Sdate      time.Time
	Edate      time.Time
	Start_date string
	End_date   string
	Desc       string
	Technology []string
	Image      string
	Islogin    bool
}

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

var blogs = []Blog{
	{
		// Name: "RIZAL",
		// // Sdate:      "12 jan 2020",
		// // Edate:      "12 feb 2020",
		// Desc:       "Dumbways asik asik asik jos",
		// Technology: []string{"Node", "Reac", "Next", "type"},
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
	route.HandleFunc("/delete-Project/{id}", deleteProject).Methods("GET")
	route.HandleFunc("/update-Project/{id}", editProject).Methods("GET")
	route.HandleFunc("/update-Project/{id}", updateProject).Methods("POST")

	route.HandleFunc("/register", formRegister).Methods("GET")
	route.HandleFunc("/register", Register).Methods("POST")
	route.HandleFunc("/login", formLogin).Methods("GET")
	route.HandleFunc("/login", Login).Methods("POST")
	route.HandleFunc("/logout", logout).Methods("GET")

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

	var store = sessions.NewCookieStore([]byte("SESSION_ID"))
	session, _ := store.Get(r, "SESSION_ID")

	if session.Values["Islogin"] != true {
		Data.Islogin = false
	} else {
		Data.Islogin = session.Values["Islogin"].(bool)
		Data.Username = session.Values["Name"].(string)
	}

	fm := session.Flashes("message")

	var flashes []string
	if len(fm) > 0 {
		session.Save(r, w)

		for _, fl := range fm {
			flashes = append(flashes, fl.(string))
		}
	}

	Data.FlashData = strings.Join(flashes, "")

	dataBlog, errQuery := connection.Conn.Query(context.Background(), "SELECT tb_blog.id, nama, start_date, end_date, description, technologies, image, tb_user.name as author FROM tb_blog LEFT JOIN tb_user ON tb_user.id = tb_blog.author ORDER BY id DESC")
	if errQuery != nil {
		fmt.Println("Message from patching : " + errQuery.Error())
		return
	}

	var result []Blog

	for dataBlog.Next() {
		var each = Blog{}

		err := dataBlog.Scan(&each.ID, &each.Nama, &each.Sdate, &each.Edate, &each.Desc, &each.Technology, &each.Image, &each.Author)
		if err != nil {
			fmt.Println("Message : " + err.Error())
			return
		}
		if session.Values["Islogin"] != true {
			each.Islogin = false
		} else {
			each.Islogin = session.Values["Islogin"].(bool)
		}
		result = append(result, each)
	}

	// fmt.Println(result)

	resData := map[string]interface{}{
		"Data":  Data,
		"Blogs": result,
	}

	w.WriteHeader(http.StatusOK)
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

	err = connection.Conn.QueryRow(context.Background(), "SELECT id, nama, start_date, end_date, description, technologies, image FROM public.tb_blog WHERE id=$1", id).Scan(
		&blogDetail.ID, &blogDetail.Nama, &blogDetail.Sdate, &blogDetail.Edate, &blogDetail.Desc, &blogDetail.Technology, &blogDetail.Image)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	blogDetail.Start_date = blogDetail.Sdate.Format("2 January 2006")
	blogDetail.End_date = blogDetail.Edate.Format("2 January 2006")

	resp := map[string]interface{}{
		"Data": Data,
		"Blog": blogDetail,
	}

	tempt.Execute(w, resp)
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
	nama := r.PostForm.Get("name")
	sdate := r.PostForm.Get("sdate")
	edate := r.PostForm.Get("edate")
	desc := r.PostForm.Get("desc")
	technology := r.Form["Technology"]

	var store = sessions.NewCookieStore([]byte("SESSION_ID"))
	session, _ := store.Get(r, "SESSION_ID")

	Author := session.Values["Id"].(int)
	fmt.Println(Author)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_blog(nama, start_date, end_date, description, technologies, image, author) VALUES ($1, $2, $3, $4, $5, 'images.png', $6)", nama, sdate, edate, desc, technology, Author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	// var newBlog = Blog{
	// 	Name: name,
	// 	// Sdate:      sdate,
	// 	// Edate:      edate,
	// 	Desc:       desc,
	// 	Technology: technology,
	// }
	// blogs = append(blogs, newBlog)
	// //fmt.Println(blogs)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_blog WHERE id=$1", id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	//blogs = append(blogs[:index], blogs[index+1:]...)

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

	var editdata = Blog{}

	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_blog WHERE id=$1", id).Scan(&editdata.ID, &editdata.Nama, &editdata.Sdate, &editdata.Edate, &editdata.Desc, &editdata.Technology, &editdata.Image)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	// for index, data := range blogs {
	// 	if index == id {
	// 		DataBlog = Blog{
	// 			ID:         id,
	// 			Name:       data.Name,
	// 			Sdate:      data.Sdate,
	// 			Edate:      data.Edate,
	// 			Desc:       data.Desc,
	// 			Technology: data.Technology,
	// 		}
	// 	}
	// }

	EditBlog := map[string]interface{}{
		"Blog": editdata,
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
	desc := r.PostForm.Get("desc")
	technology := r.Form["technology"]
	const dateformat = "2 January 2006 "
	sdate, _ := time.Parse(dateformat, r.PostForm.Get("sdate"))
	edate, _ := time.Parse(dateformat, r.PostForm.Get("edate"))
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	_, err = connection.Conn.Exec(context.Background(), "UPDATE tb_blog SET nama = $1, start_date = $2, end_date = $3, description = $4, technologies = $5 WHERE id = $6", name, sdate, edate, desc, technology, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Message: " + err.Error()))
		return
	}

	// blogs[id].Name = name
	// // blogs[id].Sdate = sdate
	// // blogs[id].Edate = edate
	// blogs[id].Desc = desc
	// blogs[id].Technology = technology

	// fmt.Println(blogs)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func formRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var temp, err = template.ParseFiles("html/register.html")
	if err != nil {
		w.Write([]byte("Message: " + err.Error()))
		return
	}

	temp.Execute(w, Data)
}

func Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	name := r.PostForm.Get("name")
	email := r.PostForm.Get("email")
	pw := r.PostForm.Get("pw")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(pw), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user(name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	var store = sessions.NewCookieStore([]byte("SESSION_ID"))
	session, _ := store.Get(r, "SESSION_ID")
	session.AddFlash("successfully register!", "message")
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusMovedPermanently)
}

func formLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var temp, err = template.ParseFiles("html/login.html")
	if err != nil {
		w.Write([]byte("Message: " + err.Error()))
		return
	}

	var store = sessions.NewCookieStore([]byte("SESSION_ID"))
	session, _ := store.Get(r, "SESSION_ID")

	fm := session.Flashes("message")

	var flashes []string
	if len(fm) > 0 {
		session.Save(r, w)
		for _, fl := range fm {
			flashes = append(flashes, fl.(string))
		}
	}

	Data.FlashData = strings.Join(flashes, "")

	temp.Execute(w, Data)
}

func Login(w http.ResponseWriter, r *http.Request) {

	var store = sessions.NewCookieStore([]byte("SESSION_ID"))
	session, _ := store.Get(r, "SESSION_ID")

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("pw")

	user := User{}

	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email=$1", email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("message : " + err.Error()))
		return
	}

	session.Values["Islogin"] = true
	session.Values["Name"] = user.Name
	session.Values["Id"] = user.Id
	session.Options.MaxAge = 10800 // 1 jam = 3600detik | 3jam = 10800

	session.AddFlash("succesfully login!", "message")
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout.")
	var store = sessions.NewCookieStore([]byte("SESSION_ID"))
	session, _ := store.Get(r, "SESSION_ID")
	session.Options.MaxAge = -1 // gak boleh kurang dari 0
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
