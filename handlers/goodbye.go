package handlers

import (
	"io"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye!")
	io.WriteString(wr, "Goodbye!")

}