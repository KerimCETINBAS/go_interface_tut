package main

import (
	"fmt"

	"github.com/kerimcetinbas/gotut/services"
)

func main() {
	us := services.UserService()

	us.AddUser("hasan")
	us.AddUser("mehmet")

	fmt.Println(us.GetUsers())
}

// private protected public internal
