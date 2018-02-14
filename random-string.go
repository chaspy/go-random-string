package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    arr := []string{"taro", "ken", "akira", "hiroki"}

    fmt.Println(arr)
    shuffle(arr)
    fmt.Println(arr)
}

func shuffle(a []string) {
    rand.Seed(time.Now().UnixNano()) 
    for i := range a {
        j := rand.Intn(i + 1)
        a[i], a[j] = a[j], a[i]
    }
}
