package features

import (
	"time"

	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/asterisk/middlewares"
	"github.com/Lukaesebrot/asterisk/static"
	"github.com/Lukaesebrot/asterisk/users"
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

// Initialize initializes all features
func Initialize(session *discordgo.Session) {
	// Define and initialize the router
	router := dgc.Create(&dgc.Router{
		Prefixes: []string{
			"$",
			"<@!" + static.Self.ID + ">",
			"<@" + static.Self.ID + ">",
			"as!",
			"你",
		},
		IgnorePrefixCase: true,
		BotsAllowed:      false,
		PingHandler:      infoCommand,
	})
	router.Initialize(session)

	// Create the general rate limiter
	generalRateLimiter := dgc.NewRateLimiter(5*time.Second, 1*time.Second, func(ctx *dgc.Ctx) {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error("Hey! Don't spam!"))
	})

	// Create the hashing rate limiter
	hashingRateLimiter := dgc.NewRateLimiter(1*time.Minute, 5*time.Second, func(ctx *dgc.Ctx) {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error("You may only create a hash once a minute."))
	})

	// Create the bug report rate limiter
	bugReportRateLimiter := dgc.NewRateLimiter(10*time.Minute, 10*time.Second, func(ctx *dgc.Ctx) {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error("Please wait at least 10 minutes between two bug reports."))
	})

	// Create the feature request rate limiter
	featureRequestRateLimiter := dgc.NewRateLimiter(10*time.Minute, 10*time.Second, func(ctx *dgc.Ctx) {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, embeds.Error("Please wait at least 10 minutes between two feature requests."))
	})

	// Initialize the default help command
	router.RegisterDefaultHelpCommand(session, generalRateLimiter)

	// Initialize all the features
	initializeSettingsFeature(router, generalRateLimiter)
	initializeInfoFeature(router, generalRateLimiter)
	initializeReminderFeature(router, generalRateLimiter)
	initializeRandomFeature(router, generalRateLimiter)
	initializeMathFeature(router, generalRateLimiter)
	initializeLaTeXFeature(router, generalRateLimiter)
	initializeHashFeature(router, hashingRateLimiter)
	initializeGoogleFeature(router, hashingRateLimiter)
	initializeBugFeature(router, bugReportRateLimiter, session)
	initializeRequestFeature(router, featureRequestRateLimiter, session)
	initializeDebugFeature(router, generalRateLimiter)
	initializeSayFeature(router, generalRateLimiter)
	initializeStarboardFeature(session)

	// Register all the middlewares
	router.AddMiddleware("*", middlewares.InjectGuildObject)
	router.AddMiddleware("*", middlewares.CheckCommandChannel)
	router.AddMiddleware("*", middlewares.InjectUserObject)
	router.AddMiddleware("bot_admin", middlewares.CheckInternalPermissions("BOT_ADMINISTRATOR", users.PermissionAdministrator))
	router.AddMiddleware("bot_mod", middlewares.CheckInternalPermissions("BOT_MODERATOR", users.PermissionModerator, users.PermissionAdministrator))
	router.AddMiddleware("guild_admin", middlewares.CheckGuildPermissions("ADMINISTRATOR", discordgo.PermissionAdministrator))
}
