Photokata backend

requirments:
golang
go get github.com/gorilla/mux github.com/lib/pq
postgres sql

app.go to see the endpoints

Testing:
curl -v http://localhost:9000/users
curl -v http://localhost:9000/user/1
curl -d '{"id":0, "firstname":"alex", "lastname":"alex"}' -H "Content-Type: application/json" -X POST http://localhost:9000/user
