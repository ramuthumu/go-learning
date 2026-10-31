package main

import "fmt"

// 1. Defining a Struct (a blueprint for data)
// Similar to a Python class or dataclass
type User struct {
	Name  string
	Age   int
	Email string
}

// 2. Value Receiver Method (Read-only, works like a method with 'self' in Python)
// (u User) is the "receiver" - it attaches this function to the User struct.
func (u User) Greet() string {
	return fmt.Sprintf("Hi, I'm %s and I am %d years old.", u.Name, u.Age)
}

// 3. Pointer Receiver Method (Allows modifying the struct's fields!)
// Notice the * in (*User). It means we pass a reference/pointer, not a copy.
func (u *User) HaveBirthday() {
	u.Age++ // Modifies the original User's age directly
}

func main() {
	// Creating a User instance
	user1 := User{
		Name:  "Ramu",
		Age:   30,
		Email: "ramu@example.com",
	}

	// Accessing fields directly
	fmt.Println("User Name:", user1.Name)
	fmt.Println("User Email:", user1.Email)

	// Calling a method on user1
	greeting := user1.Greet()
	fmt.Println(greeting)

	// Modifying fields using a pointer receiver method
	fmt.Println("\nCelebrating birthday...")
	user1.HaveBirthday()

	fmt.Printf("New Age: %d\n", user1.Age)
	fmt.Println(user1.Greet())
}
