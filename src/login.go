// forms.go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"kubernetes"
)

type userLogin struct {
	Username string
	Password string
}

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc("/landing", landing)
	http.HandleFunc("/dashboard", dashboard)
	http.ListenAndServe(":8080", nil)
}

//handles log in page
func login(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
	}

	details := userLogin{
		Username: r.FormValue("username"),
		Password: r.FormValue("psw"),
	}

	log.Println(details)
	//Attempt login by calling LDAP verify credentials
	log.Println("Username: " + details.Username)
	log.Println("Pass: " + details.Password)
	auth := verifyCredentials(details.Username, details.Password)
	log.Println(auth)

	//authorize user and create JWT
	if auth {
		fmt.Println("starting auth ...")
		for _, cookie := range r.Cookies() {
			fmt.Print("Cookie User: ")
			fmt.Println(w, cookie.Name)
			readJWT(cookie.Name)
		}
		token := generateJWT(details.Username)
		setJwtCookie(token, w, r)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}

}

func landing(w http.ResponseWriter, r *http.Request) {
	//for _, cookie := range r.Cookies() {
	cookie := r.Cookies()[0]
	if cookie != nil {
		fmt.Print("Cookie User: ")
		fmt.Println(w, cookie.Name)
		readJWT(cookie.Name)
	} else {
		http.Error(w, "Page Requires Authentication", http.StatusMethodNotAllowed)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/landing.html"))
	tmpl.Execute(w, nil)

	//re-direct upon succesful log-in
	fmt.Println("starting landing page")
	for _, cookie := range r.Cookies() {
		fmt.Print("Cookie User: ")
		fmt.Println(w, cookie.Name)
		readJWT(cookie.Name)
	}
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	//for _, cookie := range r.Cookies() {
	cookie := r.Cookies()[0]
	if cookie != nil {
		fmt.Print("Cookie User: ")
		fmt.Println(w, cookie.Name)
		readJWT(cookie.Name)
	} else {
		http.Error(w, "Page Requires Authentication", http.StatusMethodNotAllowed)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))
	tmpl.Execute(w, nil)

	//re-direct upon succesful log-in
	fmt.Println("starting dashboard")
	for _, cookie := range r.Cookies() {
		fmt.Print("Cookie User: ")
		fmt.Println(w, cookie.Name)
		readJWT(cookie.Name)
	}

	client := kubernetes.clientConfig.kubeClient{"","",""}

	client.connectClient()

}
