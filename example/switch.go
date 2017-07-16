package example

import (
	"fmt"
	"time"
)

func SwitchFunc() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	now := time.Now().Weekday()
	switch now {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's the weekday")
	}
	whatIam := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println(" I am bool")
		case string:
			fmt.Println(" I am string")
		case int:
			fmt.Println(" I am int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatIam(true)
	whatIam(100)
	whatIam("hello")
	whatIam([]string{})

}
