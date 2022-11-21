package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	const (
		noTmpHTML = "<html><body><h1>NO TEMPLATE.</h1></body></html>"
		mail1     = "t2suna@gmail.com"
		pass1     = "password"
	)

	tf, er := template.ParseFiles("templates/hello.html")
	if er != nil {
		tf, _ = template.New("index").Parse(noTmpHTML)
	}

	handler1 := func(w http.ResponseWriter, _ *http.Request) {
/*		if msg != "" {
			item := struct {
				Title   string
				Message string
			}{
				Title:   "Send Values",
				Message: msg,
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
		*/
		er = tf.Execute(w, nil)
		if er != nil {
			log.Fatal(er)

	}

	tf2, er := template.ParseFiles("templates/home.html")
	if er != nil {
		tf, _ = template.New("index").Parse(noTmpHTML)
	}

	handler2 := func(w http.ResponseWriter, rq *http.Request) {
		mail := rq.FormValue("account")
		pass := rq.FormValue("pass")
		println(mail, pass)
		if mail == mail1 && pass == pass1 {

			er = tf2.Execute(w, nil)
			if er != nil {
				log.Fatal(er)
			}

		} else {
			handler1(w, rq, mail+pass)
			//エラーの渡し方を考えておく
		}

	}

	http.HandleFunc("/new/", handler1, )
	http.HandleFunc("/login/", handler2)

	log.Fatal(http.ListenAndServe("", nil))
}
