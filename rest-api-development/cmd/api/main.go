package main

import "log"

const (
	postgresDNS = "postgres://root:secret@localhost:5455/blog?sslmode=disable"
)

func main() {
	log.Println("Hello World!")
}
