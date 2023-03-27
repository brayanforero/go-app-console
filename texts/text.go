package texts

import "fmt"

func GetAllTranslatedText(lang string) map[string]string {

	es := map[string]string{
		"header": "Ingrese su nombre",
		"body":   "Lista de usuarios",
		"alert":  "¿Quieres continuar?",
	}
	en := map[string]string{
		"header": "Enter your name",
		"body":   "List of users",
		"alert":  "Do you want continue?",
	}

	if lang == "es" {
		return es
	}

	return en
}

func PrintHeader() {
	fmt.Println("Please chooose a languaje")
	fmt.Println("1. Type 'en' for English")
	fmt.Println("2. Type 'es' for Spanish")
}
