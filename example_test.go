package hardwareid_test

import (
	"fmt"
	"log"

	"github.com/sandipmavani/hardwareid"
)

func Example() {
	id, err := hardwareid.ID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}

func Example_protected() {
	appID := "Corp.SomeApp"
	id, err := hardwareid.ProtectedID(appID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}
