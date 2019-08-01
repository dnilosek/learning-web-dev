package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/dnilosek/learning-web-dev/code/038-mongodb/007-handson/controllers"
	"github.com/dnilosek/learning-web-dev/code/038-mongodb/007-handson/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var tpl *template.Template
var dbSessionsCleaned time.Time

const cookieName string = "session"

var sController *controllers.SessionController
var uController *controllers.UserController
var client *mongo.Client

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func main() {
	// Create controllers
	uController = controllers.NewUserController(client)
	sController = controllers.NewSessionController(client)

	// Start session cleaner
	go func() {
		for {
			time.Sleep(30 * time.Second)
			num, err := sController.CleanSessions()
			if err != nil {
				panic(err)
			}
			fmt.Println("Cleaned", num, "sessions")
		}
	}()
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	// Get session cookie
	c, err := req.Cookie(cookieName)
	if err != nil {
		// No cookie, no problem
		tpl.ExecuteTemplate(w, "index.gohtml", nil)
	} else {
		fmt.Println("Session id:", c.Value)
		s, err := sController.GetSession(c.Value)
		if err != nil {
			// Something broke clear the cookie and go back
			clearCookie(w)
			http.Redirect(w, req, "/", http.StatusSeeOther)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		s, _ = sController.UpdateActivity(c.Value, s)
		// Session exists
		u, err := uController.GetUser(s.UserID)
		if err != nil {
			// Session exists but invalid user, this is bad error
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		tpl.ExecuteTemplate(w, "index.gohtml", u)
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie(cookieName)
	if err != nil {
		// Get outta here
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	s, err := sController.GetSession(c.Value)
	if err != nil {
		// Something broke clear the cookie and go back
		clearCookie(w)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	s, _ = sController.UpdateActivity(c.Value, s)

	// Session exists
	u, err := uController.GetUser(s.UserID)
	if err != nil {
		// Session exists but invalid user, this is bad error
		http.Error(w, "User does not exist", http.StatusNotFound)
		return
	}

	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {

	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")

		// username taken?
		if uController.UsernameExists(un) {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		u := models.User{un, bs, f, l, r}
		_, err = uController.CreateUser(u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie(cookieName)
	if err != nil {
		return false
	}
	_, err = sController.GetSession(c.Value)
	if err != nil {
		// Something broke clear the cookie
		clearCookie(w)
		return false
	}
	return true
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, err := uController.GetUser(un)
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err = bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		s := models.Session{
			UserID:       un,
			LastActivity: time.Now(),
		}

		sId, err := sController.CreateSession(s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		c := &http.Cookie{
			Name:  cookieName,
			Value: sId,
		}
		http.SetCookie(w, c)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie(cookieName)
	if err != nil {
		// Get outta here
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	_, err = sController.GetSession(c.Value)
	if err != nil {
		// Something broke clear the cookie and go back
		clearCookie(w)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	// Session exists
	clearCookie(w)
	err = sController.DeleteSession(c.Value)
	if err != nil {
		http.Error(w, "Unable to remove session", http.StatusInternalServerError)
		return
	}

	numClean, err := sController.CleanSessions()
	if err != nil {
		http.Error(w, "Unable to clean sessions", http.StatusInternalServerError)
		return
	}
	fmt.Println("Cleaned", numClean, "Sessions")
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func clearCookie(w http.ResponseWriter) {
	c := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
}
