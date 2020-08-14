package main

import "os"
import "github.com/go-martini/martini"

func main() {
    m := martini.Classic()
    m.Get("/", func() string {
        return "Hello world!"
    })
    m.RunOnAddr("0.0.0.0:" + os.Getenv("PORT"))
}
