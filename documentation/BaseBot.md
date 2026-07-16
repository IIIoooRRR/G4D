# func cdn-url
the creation of a bot via is presented .yaml config with the addition of my personal team
to add links to cdn photos according to the criteria of rarity, fandom, name (parsed from the path that a certain channel is sent)
```go
func main() {
cfg := g4d.MustLoadCfg("internal.locale.yaml")
logger, err := zap.NewDevelopment()
if err != nil {
    logger.Error("Error initializing logger", zap.Error(err))
}
bot, err := cfg.NewBot(logger, context.Background())
if err != nil {
    logger.Error("Error initializing bot", zap.Error(err))
}
err = bot.Run()
if err != nil {
    logger.Error("Error bot run", zap.Error(err))
}
	bot.AddCommands([]g4d.CommandTemplate{
		{_const.EventMessageCreate, "upload CDN", getCDN},
	})
}

func GetCDN(event *gw.RawEvent) error {
d := parse.Event[shema.GetMessage](event)
if uploadChannels[d.ChannelID] != true {
    return nil
}
cardUrl := d.Attachments[0].URL
content := strings.Split(d.Content, "/")
rarity := content[6]
fandom := content[7]
name := content[8]
logger.Info(zap.String("rarity:", rarity), zap.String("fd:", fandom), zap.String("name:" name), zap.String(url: "cardUrl"))
return services.Link.Add(rarity, fandom, name, cardUrl)
}

```
