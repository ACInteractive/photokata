Photokata photo services


requirments:

golang

go get github.com/gorilla/mux

go get github.com/lib/pq

go get github.com/pascaldekloe/jwt

postgres sql

app.go to see the endpoints


Testing:

curl -d "filename=test.png" --request POST --data-binary "@canvas1.png" http://localhost:9001/photo