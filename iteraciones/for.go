package iteraciones

import (
	"fmt"
)

func Iterar() {

	for i := 0; i < 100; i += 5 {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
}
