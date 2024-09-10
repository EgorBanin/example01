package main

import (
	"activeobject/after"
	"activeobject/before"
	"activeobject/common"
	"fmt"
)

func p[T any](v T) *T {
	return &v
}

func main() {
	cfg := common.Config{
		GayEnabled:     false,
		AgeOfConsent:   18,
		DickPicAllowed: false,
	}
	rq := common.CreateProfileRequest{
		Name:             "Lexx",
		Gender:           "m",
		Photo:            "userpic.jpg",
		Age:              39,
		Purposes:         []string{"sex", "marriage", "friendship"},
		LookingForGender: []string{"m", "f"},
		LookingForAge: common.AgeRange{
			From: p(18),
			To:   p(30),
		},
	}
	rsBefore, err := before.Handle(cfg, rq)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", rsBefore)
	}

	rsAfter, err := after.Handle(cfg, rq)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("%v\n", rsAfter)
	}

}
