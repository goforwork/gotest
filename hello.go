package main

import (
    "fmt"
    // "reflect"
)
 
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    a := s{"tony","male"}
    
    // res := reflect.TypeOf(a)
    // res := reflect.ValueOf(a)

    // a.f()

    fmt.Printf("%v",res);

}


type s struct{
    name string
    sex  string
}

func (s *s) f() {
    fmt.Println(s.name)
}