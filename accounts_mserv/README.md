Photokata account services


requirments:

golang

go get github.com/gorilla/mux

go get github.com/lib/pq

go get github.com/pascaldekloe/jwt

postgres sql

app.go to see the endpoints


Testing:

curl -v http://localhost:9000/users

curl -v http://localhost:9000/token

curl -v http://localhost:9000/user/1

curl -d '{"id":0, "firstname":"alex", "lastname":"alex"}' -H "Content-Type: application/json" -X POST http://localhost:9000/user
