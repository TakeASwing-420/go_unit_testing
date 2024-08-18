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
    person := Person{Name: "Alice", Age: 30}
    jsonData, err := json.Marshal(person)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(jsonData)) // {"name":"Alice","age":30}
}
