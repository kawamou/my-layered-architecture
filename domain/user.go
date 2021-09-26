package domain

type User struct {
	ID   string `firestore:"id"`
	Name string `firestore:"name"`
	Age  int    `firestore:"age"`
}

func NewUser(id, name string, age int) *User {
	return &User{
		ID:   id,
		Name: name,
		Age:  age,
	}
}
