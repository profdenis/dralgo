package main

import "fmt"

func main() {
    fmt.Println("Starting flowchart...")
    fmt.Println("Entrez votre nom : ")

    // Lire nom
    var tempString string
    _, _ = fmt.Scanln(&tempString)
    nom := tempString

    fmt.Println("Bonjour ", nom)
    fmt.Println("Entrez votre âge : ")

    // LireEntier age
    var tempInt int
    _, err := fmt.Scanln(&tempInt)
    if err != nil {
        fmt.Println("La valeur entrée n'est pas un entier")
        return
    }
    age := tempInt

    fmt.Println("Vous avez ", age, " ans")
    if age >= 18 {
        fmt.Println("Vous êtes majeur")
    } else {
        fmt.Println("Vous êtes mineur")
        fmt.Println("Ending flowchart...")
    }
}
