package controller

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var (
	tmp *template.Template
)

func init() {
	tmp = template.Must(template.New("tmpl1.html").ParseFiles("D:\\dev\\go\\go-line-notify\\templates\\tmpl1.html"))
}

type Server struct {
	listenAddr string
}

func NewServer(lisenaddr string) *Server {
	return &Server{listenAddr: lisenaddr}
}

func (s *Server) Start() error {
	http.HandleFunc("/", s.Index)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")
	data, err := db.GetByDate(start, end)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmp.Execute(w, data)
	fmt.Fprintln(w, "From: ", start, "To: ", end)
}
