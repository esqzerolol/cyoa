package storytelling

import (
	"errors"
	"fmt"
)

var ErrDeadEndReached = errors.New("dead end reached")
var ErrNoIntro = errors.New("no intro located")

func clearScreen() {
    fmt.Print("\033[H\033[2J")
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
		return "", ErrDeadEndReached
	} else {
		return nextArc, nil
	}
}

func (s *StoryArc) printTitle() {
	fmt.Println(s.Title)
	fmt.Println()
}

func (s *StoryArc) printStory() {
	for _, curParagraph := range s.Story {
		fmt.Println(curParagraph)
	}
	fmt.Println()
}

func (s *StoryArc) printOptions() (string, error) {
	if len(s.Options) == 0 {
		return "", ErrDeadEndReached
	}
	for key, curOption := range s.Options {
		fmt.Println(curOption.OptionText)
		fmt.Printf("[%d]: %s", (key+1), curOption.NextArc)
		fmt.Println()
	}

	return s.Options[chooseNumber(len(s.Options))].NextArc, nil
	// return s.Options[0].NextArc, nil
}

func chooseNumber(maxNum int) (chosen int) {
	for {
		fmt.Println("Escolha o numero do proximo arco que deseja ir:")
		_, err := fmt.Scan(&chosen)
		if err != nil || chosen <= 0 || chosen > maxNum {
			fmt.Println("Invalid input, please type a valid number.")
			continue
		}
		clearScreen()
		return (chosen-1)
	}
}

func StartAdventure(story map[string]StoryArc) (lastArc string, err error) {
	lastArc = "intro"
	if val, ok := story[lastArc]; ok {
		for {
			nextArc, err := val.decisionMenu()
			if err != nil {
				return lastArc, ErrDeadEndReached
			}
			val = story[nextArc]
			lastArc = nextArc
		}
	} 
	return "", nil
}
