package features

import (
	"fmt"

	"github.com/Lukaesebrot/asterisk/config"
	"github.com/Lukaesebrot/asterisk/embeds"
	"github.com/Lukaesebrot/dgc"
	"github.com/valyala/fasthttp"
)

// initializeRedeployFeature initializes the redeploy feature
func initializeRedeployFeature(router *dgc.Router) {
	// Register the 'redeploy' command
	router.RegisterCmd(&dgc.Command{
		Name:        "redeploy",
		Description: "[Bot Admin only] Redeploys the current bot instance",
		Usage:       "redeploy",
		Example:     "redeploy",
		Flags: []string{
			"bot_admin",
		},
		IgnoreCase:  true,
		RateLimiter: nil,
		Handler:     redeployCommand,
	})
}

// redeployCommand handles the 'redeploy' command
func redeployCommand(ctx *dgc.Ctx) {
	// Acquire a request object
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	// Acquire a response object
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	// Prepare a client for the request
	client := &fasthttp.Client{}
	request.SetRequestURI(config.CurrentConfig.Redeployment.WebhookURL)
	request.Header.SetMethod("GET")
	request.Header.Set("Redeployment-Token", config.CurrentConfig.Redeployment.Token)

	// Make a GET request to the redeployment webhook
	err := client.Do(request, response)
	if err != nil {
		ctx.RespondEmbed(embeds.Error(err.Error()))
		return
	}
	if response.StatusCode() != 200 {
		ctx.RespondEmbed(embeds.Error(fmt.Sprintf("Status code: %d", response.StatusCode())))
		return
	}
	ctx.RespondEmbed(embeds.Success("The bot is being redeployed."))
}
