package main

import (
    "fmt"
    "encoding/json"
)

type Response1 struct {
    Page   int
    Fruits []string
}

type Response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {
    // Encoding
    // encoding basic data types to JSON
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.7)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("go")
    fmt.Println(string(strB))

    // encoding slices and maps to JSON
    slcD := []string{"apple", "peach", "banana"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "banana": 10}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    mapTest := map[string][]string{"fruits": []string{"apple", "banana"}}
    mapTestB, _ := json.Marshal(mapTest)
    fmt.Println(string(mapTestB))

    mapTest2 := map[string]interface{}{}
    mapTest2["string"] = "test"
    mapTest2["bool"] = true
    mapTest2["int"] = 1
    mapTest2["float"] = 1.5
    mapTest2["array"] = []string{"one", "two", "three"}
    mapTest2["map"] = map[string]int{"obj1": 1, "obj2": 2, "obj3": 3}
    mapTest2B, _ := json.Marshal(mapTest2)
    fmt.Println(string(mapTest2B))

    // encoding structs
    res1D := &Response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    res2D := &Response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

    // Decoding
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

    var dat map[string]interface{}

    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    num := dat["num"].(float64)
    fmt.Println(num)
}
