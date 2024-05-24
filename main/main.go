package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

var (
	word             string
	unknown          string = ""
	trys                    = 0
	run                     = true
	gameOver                = false
	germanCharacters        = []rune{
		'e', 'n', 'i', 's', 'r', 'a', 't', 'd', 'h', 'u', 'l', 'c', 'g', 'm', 'o', 'b', 'w', 'f', 'k', 'z', 'p', 'v', 'ß', 'j', 'y', 'x', 'q',
	}
	helperPointer   = 0
	deathCharacters []rune
)

func main() {
	newWord()

	for run {
		update()
	}

	if gameOver {
		printGameOver()
	} else {
		printWin()
	}
}

func newWord() {
	content, err := os.ReadFile("wörter")
	if err != nil {
		log.Fatal(err)
	}

	words := strings.Split(string(content), "\n")
	word = words[rand.Intn(len(words))]
	word = strings.TrimRight(word, string(rune(13)))

	word = strings.Replace(word, "ü", "ue", -1)
	word = strings.Replace(word, "ä", "ae", -1)
	word = strings.Replace(word, "ö", "oe", -1)

	unknown = ""
	for i := 0; i < len(word); i++ {
		unknown += "_ "
	}

	fmt.Println(word)
}

func update() {
	PrintGameField()

	if trys >= len(assets)-1 {
		run = false
		gameOver = true
		return
	}

	if !strings.Contains(unknown, "_") {
		run = false
		return
	}

	var guess rune
	if len(os.Args) > 1 && os.Args[1] == "--easy" && helperPointer < 10 {
		guess = germanCharacters[helperPointer]
		helperPointer++
	} else {
		_, err := fmt.Scanf("%c\n", &guess)
		if err != nil {
			return
		}
	}
	deathCharacters = append(deathCharacters, guess)

	dashes := []rune(unknown)

	if strings.Contains(strings.ToLower(word), strings.ToLower(string(guess))) {
		for i, character := range word {
			if strings.ToLower(string(character)) == strings.ToLower(string(guess)) {
				dashes[2*i] = character
			}
		}
	} else {
		trys++
	}

	unknown = string(dashes)
}
