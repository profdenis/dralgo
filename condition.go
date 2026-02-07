package main

import "fmt"

func main() {
    fmt.Println("Starting flowchart...")
    fmt.Println("Vérifier la température")
    température := 25
    if température > 30 {
        fmt.Println("Activer la climatisation")
    } else {
        fmt.Println("Désactiver la climatisation")
        fmt.Println("Ending flowchart...")
    }
}
