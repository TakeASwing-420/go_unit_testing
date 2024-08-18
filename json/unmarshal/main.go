package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    jsonData := `{"name": "Alice", "age": 30}`
    var person Person
    a:= []byte(jsonData)
    fmt.Println(a)
    err := json.Unmarshal(a, &person)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(person) // {Alice 30}
}
