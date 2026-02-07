def main():
    print("Starting flowchart...")
    x = input()
    print("Premier x: ", x)
    x = input()
    print("Deuxième x: ", x)

    # LireEntier y
    try:
        y = int(input())
    except ValueError:
        print("La valeur entrée n'est pas un entier")
        import sys; sys.exit(1)

    print("Premier y: ", y)

    # LireEntier y
    try:
        y = int(input())
    except ValueError:
        print("La valeur entrée n'est pas un entier")
        import sys; sys.exit(1)

    print("Deuxième y: ", y)
    print("Ending flowchart...")

if __name__ == "__main__":
    main()
