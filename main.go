package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var DeadEndReached = errors.New("Dead end reached")
var NoIntro = errors.New("No intro located")

func clearScreen() {
    fmt.Print("\033[H\033[2J")
}

type StoryMethods interface {
	printTitle()
	printStory()
	printOptions()
	decisionMenu()
}

type StoryArc struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []OptionsPair `json:"options"`
}

type OptionsPair struct {
	OptionText string `json:"text"`
	NextArc    string `json:"arc"`
}

func (s *StoryArc) decisionMenu() (string, error) {
	s.printTitle()
	s.printStory()
	nextArc, err := s.printOptions()
	if err != nil {
		return "", DeadEndReached
	} else {
		return nextArc, nil
	}
}

func (s *StoryArc) printTitle() {
	fmt.Println(s.Title, "\n")

}

func (s *StoryArc) printStory() {
	for _, curParagraph := range s.Story {
		fmt.Println(curParagraph)
	}
	fmt.Println()
}

func (s *StoryArc) printOptions() (string, error) {
	if len(s.Options) == 0 {
		return "", DeadEndReached
	}
	for key, curOption := range s.Options {
		fmt.Println(curOption.OptionText)
		fmt.Printf("[%d] Next Arc: %s", (key+1), curOption.NextArc)
		fmt.Println()
	}

	return s.Options[chooseNumber(len(s.Options))].NextArc, nil
	// return s.Options[0].NextArc, nil
}

func chooseNumber(maxNum int) (chosen int) {
	for {
		fmt.Println("Type the number for the next arc you want to go:")
		_, err := fmt.Scan(&chosen)
		if err != nil || chosen <= 0 || chosen > maxNum {
			fmt.Println("Invalid input, please type a valid number.")
			continue
		}
		clearScreen()
		return (chosen-1)
	}
}

func startAdventure(story map[string]StoryArc) (lastArc string, err error) {
	lastArc = "intro"
	if val, ok := story[lastArc]; ok {
		for {
			nextArc, err := val.decisionMenu()
			if err != nil {
				return lastArc, DeadEndReached
			}
			val = story[nextArc]
			lastArc = nextArc
		}
	} 
	return "", nil
}

var currentStory map[string]StoryArc

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

	switch lastArc, err := startAdventure(currentStory); err {
	case NoIntro:
		fmt.Println("No intro detected, please insert a starting point for your journey!")
	case DeadEndReached:
		fmt.Println("Parabéns, aventureiro! 🎉")
		fmt.Println("Você chegou ao final desta história no arco de", lastArc, ".")
		fmt.Println("Obrigado por jogar!")
		fmt.Println("Sinta-se à vontade para embarcar novamente — escolha caminhos diferentes e descubra novas aventuras.")
	default:
	}

}
