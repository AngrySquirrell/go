package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type Contact struct {
	ID    int
	Nom   string
	Email string
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func pauseBeforeContinue() {
	fmt.Print("\nAppuyez sur Entrée pour continuer...")
	fmt.Scanln()
}

func main() {
	// Définition des flags
	ajouter := flag.Bool("ajouter", false, "Ajouter un contact via les flags")
	name := flag.String("name", "", "Nom du contact")
	mail := flag.String("mail", "", "Email du contact")
	flag.Parse()

	contacts := make(map[int]Contact)
	nextID := 1

	// Si le flag --ajouter est utilisé
	if *ajouter {
		if *name == "" || *mail == "" {
			fmt.Println("Erreur : Les flags --name et --mail sont obligatoires avec --ajouter")
			fmt.Println("Exemple : go run main.go --ajouter --name=Jean --mail=Jean@exemple.com")
			return
		}
		
		contact := Contact{
			ID:    nextID,
			Nom:   *name,
			Email: *mail,
		}
		
		contacts[nextID] = contact
		fmt.Printf("Contact ajouté avec succès !\n")
		fmt.Printf("ID: %d | Nom: %s | Email: %s\n", nextID, *name, *mail)
		nextID++
		fmt.Println("\nPassage en mode interactif...")
		pauseBeforeContinue()
	}

	// Mode interactif normal
	for {
		clearScreen()
		fmt.Println("\n=== Mini-CRM ===")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister tous les contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Mettre à jour un contact")
		fmt.Println("5. Quitter")
		fmt.Print("Choisissez une option : ")

		var choix string
		fmt.Scanln(&choix)

		choixInt, err := strconv.Atoi(choix)
		if err != nil {
			fmt.Println("Erreur : veuillez entrer un nombre valide")
			pauseBeforeContinue()
			continue
		}

		switch choixInt {
		case 1:
			ajouterContact(contacts, &nextID)
			pauseBeforeContinue()
		case 2:
			listerContacts(contacts)
			pauseBeforeContinue()
		case 3:
			supprimerContact(contacts)
			pauseBeforeContinue()
		case 4:
			mettreAJourContact(contacts)
			pauseBeforeContinue()
		case 5:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Option invalide")
			pauseBeforeContinue()
		}
	}
}

func ajouterContact(contacts map[int]Contact, nextID *int) {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Nom : ")
	nom, _ := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)
	
	fmt.Print("Email : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	
	contact := Contact{
		ID:    *nextID,
		Nom:   nom,
		Email: email,
	}
	
	contacts[*nextID] = contact
	fmt.Printf("Contact ajouté avec l'ID %d\n", *nextID)
	*nextID++
}

func listerContacts(contacts map[int]Contact) {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact trouvé")
		return
	}
	
	fmt.Println("\n=== Liste des contacts ===")
	for id, contact := range contacts {
		fmt.Printf("ID: %d | Nom: %s | Email: %s\n", id, contact.Nom, contact.Email)
	}
}

func supprimerContact(contacts map[int]Contact) {
	var idStr string
	fmt.Print("ID du contact à supprimer : ")
	fmt.Scanln(&idStr)
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Erreur : ID invalide")
		return
	}
	
	_, existe := contacts[id]
	if !existe {
		fmt.Println("Contact non trouvé")
		return
	}
	
	delete(contacts, id)
	fmt.Printf("Contact avec l'ID %d supprimé\n", id)
}

func mettreAJourContact(contacts map[int]Contact) {
	var idStr string
	fmt.Print("ID du contact à mettre à jour : ")
	fmt.Scanln(&idStr)
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Erreur : ID invalide")
		return
	}
	
	contact, existe := contacts[id]
	if !existe {
		fmt.Println("Contact non trouvé")
		return
	}
	
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Printf("Nom actuel: %s - Nouveau nom : ", contact.Nom)
	nom, _ := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)
	
	fmt.Printf("Email actuel: %s - Nouvel email : ", contact.Email)
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	
	contact.Nom = nom
	contact.Email = email
	contacts[id] = contact
	
	fmt.Printf("Contact avec l'ID %d mis à jour\n", id)
}