package main

import "fmt"

type Pet interface {
    Name() string
    Age() int
}

type Flyer interface {
    Fly()
}

type Bird struct {
    MyName string
    MyAge int
}

func (b Bird) Name() string {
    return b.MyName
}

func (b Bird) Age() int {
    return b.MyAge
}

func (b Bird) Fly() {
    fmt.Println(b.MyName, "is flying.")
}

func main() {
    var p Pet = Bird{"Sparrow", 3}
    // var f Flyer = Bird{"Sparrow", 3}

    fmt.Println("My bird's name is", p.Name(), "and it is", p.Age(), "years old.")
    // p.Fly()
    
}