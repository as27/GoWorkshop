package main

import "fmt"

type Celsius int

func (c Celsius) String() string {
	return fmt.Sprintf("%d C", c)
}

func main() {
	c := Celsius(24)
	fmt.Println(c.String())
}
