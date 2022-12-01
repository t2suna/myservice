package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
)

// From https://qiita.com/akubi0w1/items/8701c05fe7186ceee632
// Connection
func Connection() (redis.Conn, error) {
	const Addr = "127.0.0.1:6379"

	c, err := redis.Dial("tcp", Addr)
	if err != nil {
		return nil, err
	}
	return c, nil
}

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
	}

	tf2, er := template.ParseFiles("templates/home.html")
	if er != nil {
		tf, _ = template.New("index").Parse(noTmpHTML)
	}

	handler2 := func(w http.ResponseWriter, rq *http.Request) {
		if rq.Method == "POST" {
			mail := rq.PostFormValue("account")
			pass := rq.PostFormValue("pass")

			// 接続
			c, err := Connection()
			if err != nil {
				log.Fatal(er)
			}
			defer c.Close()
			res_get, err := redis.String(c.Do("GET", mail))
			if err != nil {
				handler1(w, rq)
				return
			}

			if pass == res_get {

				er = tf2.Execute(w, nil)
				if er != nil {
					log.Fatal(er)
				}

			} else {
				handler1(w, rq)
				//エラーの渡し方を考えておく
			}
		} else {
			handler1(w, rq)
			//エラーの渡し方を考えておく
		}

	}

	http.HandleFunc("/new/", handler1)
	http.HandleFunc("/login/", handler2)

	http.ListenAndServe(":8080", nil)
}
