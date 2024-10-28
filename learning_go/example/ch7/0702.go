package main

import (
	"fmt"
	"time"
)

type Person struct {
	LastName    string
	FirstName   string
	Age         int
	LastUpdated time.Time
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s :年齢%d: 更新%v", p.LastName, p.FirstName, p.Age, p.LastUpdated)
}

func (p *Person) Aged() {
	p.Age++
	p.LastUpdated = time.Now()
}

func main() {
	p := Person{
		LastName:    "Hoge",
		FirstName:   "Hogeo",
		Age:         120,
		LastUpdated: time.Now(),
	}

	output := p.String()
	fmt.Println(output)

	p.Aged()

	fmt.Println(p.String())
}
