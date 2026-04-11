def main():
    compteur = 5
    while (compteur > 0):
        print("Compte à rebours :", compteur)
        compteur = (compteur - 1)
    print("Décollage !")
    i = 0
    print("i =", i)
    i += 1
    while (i < 3):
        print("i =", i)
        i += 1

if __name__ == "__main__":
    main()
