package embeds

import (
	"fmt"
	"time"

	"github.com/Lukaesebrot/asterisk/nodes/todos"
	"github.com/bwmarrin/discordgo"
)

// ToDos generates a ToDo list embed
func ToDos(userToDos []todos.ToDo) *discordgo.MessageEmbed {
	// Define the ToDos string
	toDosString := "No ToDo objects"
	if len(userToDos) > 0 {
		toDosString = ""
		counter := 1
		for _, toDo := range userToDos {
			if counter != 1 {
				toDosString += "\n"
			}
			toDosString += fmt.Sprintf("`%d.` Content: `%s`", counter, toDo.Content)
			counter++
		}
	}

	return &discordgo.MessageEmbed{
		Title:     "ToDo List",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0xffff00,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "ToDos",
				Value: toDosString,
			},
		},
	}
}
