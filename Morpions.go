package main

import (
	"bufio"
	"fmt"
	"os"
)

// Constantes globales
const (
	TailleDamier = 3             // Taille du damier (3x3)
	SymboleX     = "X"           // Symbole du joueur 1
	SymboleO     = "O"           // Symbole du joueur 2
	CaseVide     = " "           // Représentation d'une case vide
)

// Variables globales
var (
	damier  [TailleDamier][TailleDamier]string // Représentation du damier 2D
	joueurX = true                             // Indique si c'est le tour du joueur X
)

func main() {
	initialiserDamier()
	jouer()
}

// Initialise le damier avec des cases vides
func initialiserDamier() {
	for i := 0; i < TailleDamier; i++ {
		for j := 0; j < TailleDamier; j++ {
			damier[i][j] = CaseVide
		}
	}
}

// Lance le jeu
func jouer() {
	for {
		afficherDamier()
		ligne, colonne := saisirCoordonnees()
		placerSymbole(ligne, colonne)
		if verifierVictoire() {
			afficherDamier()
			fmt.Printf("%s a gagné !\n", nomJoueur())
			break
		}
		if partieNulle() {
			afficherDamier()
			fmt.Println("Partie nulle !")
			break
		}
		// Change de joueur
		joueurX = !joueurX
	}
}

// Affiche le damier
func afficherDamier() {
	fmt.Println("\nDamier actuel :")
	for i := 0; i < TailleDamier; i++ {
		for j := 0; j < TailleDamier; j++ {
			fmt.Printf(" %s ", damier[i][j])
			if j < TailleDamier-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < TailleDamier-1 {
			fmt.Println("---+---+---")
		}
	}
	fmt.Println()
}

// Retourne le nom du joueur actuel
func nomJoueur() string {
	if joueurX {
		return "Joueur X"
	}
	return "Joueur O"
}

// Demande au joueur de saisir des coordonnées valides
func saisirCoordonnees() (int, int) {
	var (
		ligne, colonne int
		err            error
	)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("%s, entrez les coordonnées (ligne et colonne séparées par un espace) : ", nomJoueur())
		scanner.Scan()
		input := scanner.Text()
		_, err = fmt.Sscanf(input, "%d %d", &ligne, &colonne)
		if err != nil || !coordonneesValides(ligne, colonne) {
			fmt.Println("Coordonnées invalides. Réessayez.")
			continue
		}
		if damier[ligne-1][colonne-1] != CaseVide {
			fmt.Println("Cette case est déjà occupée. Choisissez-en une autre.")
			continue
		}
		break
	}
	return ligne - 1, colonne - 1
}

// Vérifie si les coordonnées sont valides
func coordonneesValides(ligne, colonne int) bool {
	return ligne >= 1 && ligne <= TailleDamier && colonne >= 1 && colonne <= TailleDamier
}

// Place le symbole du joueur actuel sur le damier
func placerSymbole(ligne, colonne int) {
	if joueurX {
		damier[ligne][colonne] = SymboleX
	} else {
		damier[ligne][colonne] = SymboleO
	}
}

// Vérifie si un joueur a gagné
func verifierVictoire() bool {
	symbole := SymboleX
	if !joueurX {
		symbole = SymboleO
	}

	// Vérification des lignes, colonnes et diagonales
	for i := 0; i < TailleDamier; i++ {
		// Vérification de la ligne i
		if damier[i][0] == symbole && damier[i][1] == symbole && damier[i][2] == symbole {
			return true
		}
		// Vérification de la colonne i
		if damier[0][i] == symbole && damier[1][i] == symbole && damier[2][i] == symbole {
			return true
		}
	}
	// Vérification des diagonales
	if damier[0][0] == symbole && damier[1][1] == symbole && damier[2][2] == symbole {
		return true
	}
	if damier[0][2] == symbole && damier[1][1] == symbole && damier[2][0] == symbole {
		return true
	}

	return false
}

// Vérifie si la partie est nulle
func partieNulle() bool {
	for i := 0; i < TailleDamier; i++ {
		for j := 0; j < TailleDamier; j++ {
			if damier[i][j] == CaseVide {
				return false
			}
		}
	}
	return true
}
