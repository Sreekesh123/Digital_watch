package main

import (
	"fmt"
	"html/template"
	"net/http"
	"digital_watch/helper"
	"digital_watch/controllers"
)


func main() {

	//Parse an HTML file in Go by passing it to the Go parse function for access.
	htmlTemplate, err := helper.ReadHTMLFile("watch.html")

	if err!= nil {
		fmt.Println(err)
		return
	}
		
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.New("watch").Parse(string(htmlTemplate)))
		tmpl.Execute(w, helper.ShowCurrentMode())
    })

    http.HandleFunc("/click", controllers.ClickHandler)

	fmt.Printf("Server listening on the port :%s\n", "8080")

    http.ListenAndServe(":8080", nil)
}



