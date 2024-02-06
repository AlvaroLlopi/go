package users

import (
	"fmt"
	"time"

	"github.com/alvarollopi/go/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Alvaro", time.Now(), true)
	fmt.Println(u)
}
