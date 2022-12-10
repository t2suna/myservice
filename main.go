package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
)

var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret in my heart"))

func Connection() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
func ErrorSave(w http.ResponseWriter, rq *http.Request, err string, ses *sessions.Session) {
	ses.Values["error"] = err
	ses.Save(rq, w)

}

func main() {

	const (
		noTmpHTML = "<html><body><h1>NO TEMPLATE.</h1></body></html>"
	)

	tf, er := template.ParseFiles("templates/hello.html")
	if er != nil {
		tf, _ = template.New("index").Parse(noTmpHTML)
	}

	handler1 := func(w http.ResponseWriter, rq *http.Request) {
		ses, _ := cs.Get(rq, "Session")
		if ses.Values["error"] != nil {
			item := struct {
				Error string
			}{
				Error: ses.Values["error"].(string),
			}

			er = tf.Execute(w, item)
			if er != nil {
				log.Fatal(er)

			}
		} else {
			er = tf.Execute(w, nil)
			if er != nil {
				log.Fatal(er)

			}

		}
		ses.Values["error"] = nil
		ses.Save(rq, w)

	}

	tf2, er := template.ParseFiles("templates/home.html")
	if er != nil {
		tf, _ = template.New("index").Parse(noTmpHTML)
	}

	handler2 := func(w http.ResponseWriter, rq *http.Request) {
		ses, _ := cs.Get(rq, "Session")

		if rq.Method == "POST" {

			db, err := Connection()
			if err != nil {
				ErrorSave(w, rq, err.Error(), ses)
				handler1(w, rq)
			}
			defer db.Close()
			mail := rq.PostFormValue("account")
			pass := rq.PostFormValue("pass")

			var count = 0

			if err := db.QueryRow("select count(*) from users where mail = ? and pass = ?", mail, pass).Scan(&count); err != nil {
				ErrorSave(w, rq, err.Error(), ses)
			}

			if count > 0 {

				er = tf2.Execute(w, nil)
				if er != nil {
					ErrorSave(w, rq, err.Error(), ses)
					handler1(w, rq)
				}

			} else {
				ErrorSave(w, rq, "会員情報が間違っています。", ses)
				handler1(w, rq)
			}
		} else {
			ErrorSave(w, rq, "不正なリクエスト", ses)
			handler1(w, rq)
		}

	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/new/", handler1)
	http.HandleFunc("/login/", handler2)

	http.ListenAndServe(":8080", nil)
}
