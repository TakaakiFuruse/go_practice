package main

import "fmt"

type user struct {
	name  string
	email string
	level string
}

func (u *user) notify() {
	fmt.Printf("you are user %s<%s>\n", u.name, u.email)
}

func (u *user) auth() {
	fmt.Printf("your level is %s\n", u.level)
}

type admin struct {
	user
}

func main() {
	ad := admin{
		user: user{
			name:  "Mr.User",
			email: "User@User.com",
			level: "super",
		},
	}
	u := user{
		name:  "Mr.normal",
		email: "n@n.com",
		level: "user",
	}
	showAuth(&ad)
	showAuth(&u)
}

type auth interface {
	auth()
}

func showAuth(a auth) {
	a.auth()
}

type notifier interface {
	notify()
}

func sendNotifier(n notifier) {
	n.notify()
}

// package main

// import "fmt"

// type notifier interface {
// 	notify()
// }

// type user struct {
// 	name  string
// 	email string
// }

// func (u *user) notify() {
// 	fmt.Printf("Seding email to %s<%s>\n", u.name, u.email)

// }

// type admin struct {
// 	name  string
// 	email string
// }

// func (a *admin) notify() {
// 	fmt.Printf("sending admin email to %s<%s>\n", a.name, a.email)
// }

// func sendNotification(n notifier) {
// 	n.notify()
// }

// func main() {
// 	bill := user{"bill", "test@test.com"}
// 	sendNotification(&bill)

// 	lisa := admin{"lisa", "lisa@lisa.com"}
// 	sendNotification(&lisa)
// }

// package main

// import "fmt"

// type duration int

// func (d *duration) pretty() string {
// 	return fmt.Sprintf("Duration %d\n", *d)

// }

// func main() {
// 	fmt.Println(duration(42).pretty())
// }

// package main

// import "fmt"

// type notifier interface {
// 	notify()
// }

// type user struct {
// 	name  string
// 	email string
// }

// func (u user) notify() {
// 	fmt.Printf("sending user email to %s <%s>\n", u.name, u.email)
// }

// func main() {
// 	u := user{"Bill", "bill@bill.com"}
// 	u.notify()
// 	sendNotification(u)
// }

// func sendNotification(n notifier) {
// 	n.notify()
// }
