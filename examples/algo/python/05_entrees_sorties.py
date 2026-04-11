def main():
    print("Quel est votre nom ?")
    nom = input()
    print("Quel est votre âge ?")
    try:
        age = int(input())
    except ValueError:
        print("Erreur: un entier est attendu")
        exit(1)
    print("Bonjour", nom, "!")
    if (age >= 18):
        print("Vous êtes majeur.")
    else:
        print("Vous êtes mineur.")

if __name__ == "__main__":
    main()
