package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gophercises/cyoa/storytelling"
)

var currentStory map[string]storytelling.StoryArc

func main() {

	file, err := os.Open("test.json")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo gopher.json | codigo de erro: ", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&currentStory)
	if err != nil {
		fmt.Println("Erro no decode do arquivo gopher.json | codigo do erro: ", err)
	}

	switch lastArc, err := storytelling.StartAdventure(currentStory); err {
	case storytelling.ErrNoIntro:
		fmt.Println("No intro detected, please insert a starting point for your journey!")
	case storytelling.ErrDeadEndReached:
		fmt.Println("Parabéns, aventureiro! 🎉")
		fmt.Println("Você chegou ao final desta história no arco de", lastArc, ".")
		fmt.Println("Obrigado por jogar!")
		fmt.Println("Sinta-se à vontade para embarcar novamente, escolha caminhos diferentes e descubra novas aventuras.")
	default:
	}

	storytelling.WaitForExit()

}
