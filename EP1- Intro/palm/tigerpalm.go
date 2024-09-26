package tigerpalm

import (
	"fmt"

	"github.com/thitipa-palm/go-EP1/palm/internal/tiger"
)

func SayHelloTiger() {
	fmt.Println("Say Hi Palm")
	tiger.SayInternal()
}
