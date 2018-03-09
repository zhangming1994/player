package src

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/index.html")
	if err != nil {
		fmt.Println(err)
	}
	err = t.Execute(w, "你好")
	if err != nil {
		fmt.Println(err)
	}
}
