def main():
    print("Entrez un nombre :")
    try:
        x = int(input())
    except ValueError:
        print("Erreur: un entier est attendu")
        exit(1)
    if (x < 0):
        print("Négatif")
    elif (x > 1000000):
        print("Trop grand")
    else:
        print("Bien")

if __name__ == "__main__":
    main()
