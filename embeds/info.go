package embeds

import (
	"runtime"
	"strconv"
	"time"

	"github.com/Lukaesebrot/asterisk/static"
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
	"github.com/c2h5oh/datasize"
)

// Info generates an information embed
func Info(ctx *dgc.Ctx) *discordgo.MessageEmbed {
	// Read the memstats
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	heapInUse := datasize.ByteSize(memStats.HeapInuse)
	stackInUse := datasize.ByteSize(memStats.StackInuse)

	return &discordgo.MessageEmbed{
		Title:     "Information",
		Timestamp: time.Now().Format(time.RFC3339),
		Color:     0x0893d8,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "Application",
				Value: "Mode: `" + static.Mode + "`" +
					"\nVersion: `" + static.Version + "`" +
					"\nUptime: `" + time.Since(static.StartupTime).Round(time.Second).String() + "`",
			},
			{
				Name: "Discord",
				Value: "Guilds: `" + strconv.Itoa(len(ctx.Session.State.Guilds)) + "`" +
					"\nAPI latency: `" + ctx.Session.HeartbeatLatency().Round(time.Millisecond).String() + "`",
			},
			{
				Name: "System",
				Value: "OS: `" + runtime.GOOS + "`" +
					"\nArchitecture: `" + runtime.GOARCH + "`" +
					"\nCurrent Goroutines: `" + strconv.Itoa(runtime.NumGoroutine()) + "`" +
					"\nOccupied heap: `" + heapInUse.HumanReadable() + "`" +
					"\nOccupied stack: `" + stackInUse.HumanReadable() + "`",
			},
			{
				Name: "General",
				Value: "Developer(s): `Lukaesebrot#8001`" +
					"\nGitHub repository: [here](http://github.com/Lukaesebrot/asterisk)" +
					"\nInvite me: [here](https://discord.com/api/oauth2/authorize?client_id=" + static.Self.ID + "&permissions=76864&scope=bot)" +
					"\nSupport guild: [here](https://discord.gg/ddz9b86)",
			},
		},
	}
}
