package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestAffichage(t *testing.T) {
	// Capture de la sortie de la fonction affichage
	var buf bytes.Buffer
	fmt.Print = func(a ...interface{}) (n int, err error) {
		return fmt.Fprint(&buf, a...)
	}

	// Appel de la fonction affichage
	Affichage()

	expected := " 1  2  3 \n 4  5  6 \n 7  8  9 \n" // Ce que devrait afficher le damier initial
	if buf.String() != expected {
		t.Errorf("Affichage() = %v, want %v", buf.String(), expected)
	}
}

func TestGagner(t *testing.T) {
	// Tester la fonction gagner
	// Assurer que le joueur 1 gagne
	tableauMorpion = [9]string{"X", "X", "X", "4", "5", "6", "7", "8", "9"}
	if !Gagner() {
		t.Error("Expected player 1 to win")
	}

	// Tester pour un cas de victoire du joueur 2
	tableauMorpion = [9]string{"X", "O", "X", "O", "O", "X", "X", "8", "9"}
	if Gagner() {
		t.Error("Expected no winner, but player 2 wins")
	}
}

func TestPartieNulle(t *testing.T) {
	// Tester la fonction partieNulle pour une partie nulle
	tableauMorpion = [9]string{"X", "O", "X", "O", "X", "O", "O", "X", "X"}
	if !PartieNulle() {
		t.Error("Expected the game to be a draw")
	}
}

func TestRemplirCase(t *testing.T) {
	// Tester la fonction remplirCase
	tableauMorpion = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	RemplirCase(0) // Joueur 1 prend la première case
	if tableauMorpion[0] != "X" {
		t.Errorf("Expected 'X' in position 0, got %v", tableauMorpion[0])
	}

	RemplirCase(1) // Joueur 2 prend la deuxième case
	if tableauMorpion[1] != "O" {
		t.Errorf("Expected 'O' in position 1, got %v", tableauMorpion[1])
	}
}
