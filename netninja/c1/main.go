package main

impor (
	"fmt"
)

type User struct {
	Id             string   `json:"id"`
	Email          string   `json:"email"`
	HashedPassword []byte   `json:"hashed_password"`
	DateOfBirth    string   `json:"date_of_birth"`
	BirthPlace     string   `json:"birth_place"`
	Friends        []string `json:"friends"`
}

func main() {
	fmt.Print("Hello World!")
}
