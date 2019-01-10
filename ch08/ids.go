package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	fmt.Println("User id:", os.Getuid())

	var u *user.User
	u, _ = user.Current()
	fmt.Print("Group ids: ")
	groupIDs, _ := u.GroupIds()
	for _, i := range groupIDs {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
