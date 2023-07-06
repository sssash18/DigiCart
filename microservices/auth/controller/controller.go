package controller

import "net/http"

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Boy!"))
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func LogOut(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Boy!"))
}
