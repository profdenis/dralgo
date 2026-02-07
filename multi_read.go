package main

import "fmt"

func main() {
    fmt.Println("Starting flowchart...")

    // Lire x
    var tempString string
    _, _ = fmt.Scanln(&tempString)
    x := tempString

    fmt.Println("Premier x: ", x)

    // Lire x
    _, _ = fmt.Scanln(&tempString)
    x = tempString

    fmt.Println("Deuxième x: ", x)

    // LireEntier y
    var tempInt int
    _, err := fmt.Scanln(&tempInt)
    if err != nil {
        fmt.Println("La valeur entrée n'est pas un entier")
        return
    }
    y := tempInt

    fmt.Println("Premier y: ", y)

    // LireEntier y
    _, err = fmt.Scanln(&tempInt)
    if err != nil {
        fmt.Println("La valeur entrée n'est pas un entier")
        return
    }
    y = tempInt

    fmt.Println("Deuxième y: ", y)
    fmt.Println("Ending flowchart...")
}
