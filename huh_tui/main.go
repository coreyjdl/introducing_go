package main

import (
    "fmt"
    "github.com/charmbracelet/huh"
)

func main() {
    var name string

    form := huh.NewForm(
        huh.NewGroup(
            huh.NewInput().
                Title("What’s your name?").
                Value(&name),
        ),
    )

    if err := form.Run(); err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Hello, %s!\n", name)
}
