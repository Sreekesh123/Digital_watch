package controllers

import (
	"net/http"
	"digital_watch/service"

)

//The primary handler is responsible for processing the data and displaying it in the console.
func ClickHandler(w http.ResponseWriter, r *http.Request) {
    buttonIndex := r.FormValue("buttonIndex")
    if buttonIndex == "1" {
        service.HandleButton1()
    } else if buttonIndex == "2" {
        service.HandleButton2()
    } else if buttonIndex == "3" {
        service.HandleButton3()
    } else if buttonIndex == "4" {
        service.ToggleMode()
    }
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

