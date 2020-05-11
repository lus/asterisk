package concommands

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Handle handles incoming console commands
func Handle(reader *bufio.Reader, session *discordgo.Session) {
	// Read the input line
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.ReplaceAll(line, "\n", "")
	line = strings.ReplaceAll(line, "\r", "")
	split := strings.Split(line, " ")

	// Handle the executed command
	switch strings.ToLower(split[0]) {
	case "help":
		fmt.Println("help: Shows this list")
		fmt.Println("stop: Stops this Asterisk instance")
	case "stop":
		log.Println("Stopping this Asterisk instance...")
		err = session.Close()
		if err != nil {
			panic(err)
		}
		return
	default:
		fmt.Println("Command not found. Type 'help' for help.")
	}

	// Handle further commands
	Handle(reader, session)
}
