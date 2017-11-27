package main

import "fmt"

type s_mesure struct {
	pieds, pouces int
}

func main() {
	var (option int
	mesure1 s_mesure
	mesure2 s_mesure
	interchanger s_mesure)

	for  quitter := false; quitter == false; {
		imprimerMenu()
		fmt.Scanln(&option)

		switch option {
		case 0:
			quitter = true

		case 1:
			mesure1 = entrerMesure("Entrer une première mesure en pied pouces : ")

		case 2:
			mesure2 = entrerMesure("Entrer une deuxième mesure en pied pouces : ")

		case 3:
			fmt.Print("La première mesure est ", mesure1.pieds, " pieds et ", mesure1.pouces, " pouces\n")

		case 4:
			fmt.Print("La première mesure est ", mesure2.pieds, " pieds et ", mesure2.pouces, " pouces.\n")

		case 5:
			interchanger = mesure1
			mesure1 = mesure2
			mesure2 = interchanger

		case 6:
			//fmt.Print("La première mesure en centimètres est : ", mesure1.pieds*12*2.54+mesure1.pouces*2.54)

		case 7:
			//fmt.Print("La première mesure en mètres est : ", (mesure1.pieds*12*2.54+mesure1.pouces*2.54)/100)

		case 8:
			fmt.Println(ajouterUnPouce(mesure1))
		case 9:
			fmt.Println(ajouterUnPied(mesure1))
		default:
			fmt.Println("Entrer un option valide.")
		}
	}
}

func imprimerMenu() {
	fmt.Print("\n1.Entrer une donnée en pieds pouces.\n2.Entrer une seconde mesure en pieds pouces\n3.Afficher la première mesure\n4.Afficher la seconde mesure.\n5.Interchanger les mesures.\n6.Afficher la première mesure en centimètre(s).\n7.Afficher la seconde mesure en mètre(s).\n8.Ajouter un pouces à la première mesure.\n9.Ajouter un pied à la seconde mesure.\n0.Quitter.\n\nVotre choix >>> ")
}

func entrerMesure(texteAAfficher string) s_mesure {
	var (donneeEntrerParUtil int
	mesureSaisie s_mesure)
	fmt.Print(texteAAfficher)
	fmt.Scanln(&donneeEntrerParUtil)
	mesureSaisie.pieds = donneeEntrerParUtil /12
	mesureSaisie.pouces = donneeEntrerParUtil %12
	return mesureSaisie
}

func ajouterUnPouce(mesure s_mesure) s_mesure {
	if mesure.pouces == 11 {
		mesure.pouces = 0
		mesure.pieds = mesure.pieds + 1
	} else {
		mesure.pouces = mesure.pouces + 1
	}
	return mesure
}

func ajouterUnPied(mesure s_mesure) s_mesure {
	mesure.pieds = mesure.pieds + 1
	return mesure
}