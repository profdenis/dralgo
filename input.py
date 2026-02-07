def main():
    print("Starting flowchart...")
    print("Entrez votre nom : ")
    nom = input()
    print("Bonjour ", nom)
    print("Entrez votre âge : ")

    # LireEntier age
    try:
        age = int(input())
    except ValueError:
        print("La valeur entrée n'est pas un entier")
        import sys; sys.exit(1)

    print("Vous avez ", age, " ans")
    if age >= 18:
        print("Vous êtes majeur")
    else:
        print("Vous êtes mineur")
        print("Ending flowchart...")

if __name__ == "__main__":
    main()
