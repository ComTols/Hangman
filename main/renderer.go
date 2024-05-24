package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"strings"
)

var (
	assets         = getAssets()
	clearFunctions = getClears()
)

func getClears() map[string]func() {
	functions := make(map[string]func()) //Initialize it
	functions["linux"] = func() {
		cmd := exec.Command("functions") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	functions["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	return functions
}

func getAssets() []string {
	content, err := os.ReadFile("assetset")
	if err != nil {
		log.Fatal(err)
	}

	var assetset = string(content)
	var assets = strings.Split(assetset, "t")
	return assets
}

func PrintGameField() {
	clearFunctions[runtime.GOOS]()

	if slices.Contains(os.Args, "--tipp") {
		fmt.Print("You tried: ")
		for _, c := range deathCharacters {
			fmt.Printf("%c ", c)
		}
		fmt.Println()
	}

	fmt.Println(assets[trys])

	fmt.Println("\n" + unknown)
}

func printGameOver() {
	fmt.Printf("You failed!!!\nThe word was: %s\nTry again...", word)
}

func printWin() {
	fmt.Printf("You win!\nCongratulation...")
}
