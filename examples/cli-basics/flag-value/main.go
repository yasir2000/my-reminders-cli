package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

//custom type aas value interface
type IdsFlag []string

func (ids IdsFlag) String() string {
	return strings.Join(ids, ",")
}

func (ids *IdsFlag) Set(id string) error {
	*ids = append(*ids, id)
	return nil
}

type person struct {
	name string
	born time.Time
}

func (p person) String() string {
	return fmt.Sprintf("My name is : %s, and I am %s", p.name, p.born.String())
}

func (p *person) Set(name string) error {
	p.name = name
	p.born = time.Now()
	return nil
}

func main() {
	var ids IdsFlag
	var p person

	flag.Var(&ids, "id", "the id appended to the list")
	flag.Var(&p, "name", "the name of the person")

	flag.Parse()
	fmt.Println(ids)
	fmt.Println(p)

}
