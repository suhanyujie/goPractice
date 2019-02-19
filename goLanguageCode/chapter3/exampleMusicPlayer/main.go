package main

import (
	"bufio"
	"fmt"
	"os"
	"practice/goLanguageCode/chapter3/exampleMusicPlayer/library"
	"practice/goLanguageCode/chapter3/exampleMusicPlayer/mp"
	"strconv"
	"strings"
)

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			m1, err := lib.Get(i)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(i+1, ":", m1.Name, m1.Artist, m1.Source, m1.Type)
		}
	case "add":
		if len(tokens) != 6 {
			fmt.Println("Usage: lib add <name><artist><source><type>")
		} else {
			id++
			lib.Add(&library.MusicEntry{strconv.Itoa(id), tokens[2],tokens[3],tokens[4],tokens[5]})
		}
	default:
		fmt.Println(`
	Enter following commands to control the player:
	lib list -- View the existing music lib
	lib add <name><artist><source><type> -- Add a music to the music lib
	lib remove <name> -- Remove the specified music from the lib
	play <name> -- Play the specified music
	`)
	}

}

func handlePlayCommand(tokens []string)  {
	if len(tokens) != 2 {
		fmt.Println("Usage: play <name>")
		return
	}
	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music ",tokens[1],"does not exist.")
		return
	}
	mp.Play(e.Source,e.Type)
}

func main() {
	fmt.Println(`
	Enter following commands to control the player:
	lib list -- View the existing music lib
	lib add <name><artist><source><type> -- Add a music to the music lib
	lib remove <name> -- Remove the specified music from the lib
	play <name> -- Play the specified music
	`)
	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter command ->")
		rawLine,_,_ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line," ")
		if tokens[0] == "lib" {
			if len(tokens)<2 {
				fmt.Println(`
	Enter following commands to control the player:
	lib list -- View the existing music lib
	lib add <name><artist><source><type> -- Add a music to the music lib
	lib remove <name> -- Remove the specified music from the lib
	play <name> -- Play the specified music
	`)
				continue
			}
			handleLibCommands(tokens)
		} else if(tokens[0] == "play") {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized command: ",tokens[0])
		}
	}
}
