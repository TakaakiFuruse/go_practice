////////////////// OOP 1 ///////////////////////////////////////
// package main

// import (
// 	"fmt"
// )

// // def initialise name and email
// type user struct {
// 	name  string
// 	email string
// }

// // def notify; p u.name and...
// func (u user) notify() {
// 	fmt.Printf("User: Sending User Email To %s<%s>\n",
// 		u.name,
// 		u.email)
// }

// // changeEmail!(mail) ; u.mail = mail
// func (u *user) changeEmail(email string) {
// 	u.email = email
// }

// // driver code
// func main() {
// 	// Values of type user can be used to call methods
// 	// declared with a value receiver.
// 	bill := user{"Bill", "bill@email.com"}
// 	fmt.Println(bill)
// 	bill.notify()
// 	bill.changeEmail("bill@gmail.com")

// 	// & => call destructive method

// 	lisa := &user{"Lisa", "lisa@email.com"}
// 	fmt.Println(lisa)
// 	lisa.notify()
// 	lisa.changeEmail("lisa@gmail.com")
// }

////////////////// OOP2 ///////////////////////////////////
package main

import (
	"fmt"
)

// notifier is an interface that defined notification type behavior.
// attr_accessor: notify
type notifier interface {
	notify()
}

// user defines a user in the program.
// def initialize ; @name = name, @email...
type user struct {
	name  string
	email string
}

// notify implements a method that can be called via
// a value of type user.

func (u user) notify() {
	fmt.Printf("user: Sending user Email To %s<%s>\n",
		u.name,
		u.email)
}

// main is the entry point for the application.
func main() {
	// Create a value of type user.
	u := &user{"Jill", "jill@email.com"}

	// Pass a pointer of of type user to the function.
	sendNotification(u)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(notify notifier) {
	notify.notify()
}
