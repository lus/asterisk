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
		fmt.Println("    > help: Shows this list")
		fmt.Println("    > stop: Stops this Asterisk instance")
		fmt.Println("    > say: Prints the message defined in the second+ argument(s) to the channel with the ID of the first one")
	case "stop":
		log.Println("Stopping this Asterisk instance...")
		err = session.Close()
		if err != nil {
			panic(err)
		}
		return
	case "say":
		if len(split) < 3 {
			fmt.Println("> Invalid syntax.")
			break
		}
		_, err = session.ChannelMessageSend(split[1], strings.Join(split[2:], " "))
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("> Success.")
	default:
		fmt.Println("> Command not found. Type 'help' for help.")
	}

	// Handle further commands
	Handle(reader, session)
}
