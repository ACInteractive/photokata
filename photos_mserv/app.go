package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/pascaldekloe/jwt"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

type myData struct {
	FileName string
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/token", a.getToken).Methods("GET")
	a.Router.HandleFunc("/photo", a.recvPhoto).Methods("POST")
}

func (a *App) getToken(w http.ResponseWriter, r *http.Request) {
	_, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("key generation error:", err)
		return
	}

	var claims jwt.Claims
	claims.Subject = "alice@example.com"

	now := time.Now().Round(time.Second)
	claims.Issued = jwt.NewNumericTime(now)
	claims.Expires = jwt.NewNumericTime(now.Add(10 * time.Minute))

	// issue a JWT
	token, err := claims.EdDSASign(privateKey)

	respondWithJSON(w, http.StatusOK, map[string]string{"token": string(token)})
}

func (a *App) recvPhoto(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filename := r.Form.Get("filename")

	defer r.Body.Close()

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	io.Copy(f, r.Body)

	respondWithJSON(w, http.StatusOK, map[string]string{"token": "test"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
