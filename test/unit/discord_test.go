package discord_test

import (
	"os"
	"testing"

	"github.com/arthur404dev/gobot-discord/discord"
	"github.com/sirupsen/logrus"
)

func TestInit(t *testing.T) {

	os.Setenv("BOT_TOKEN", "test-token")

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	file, err := os.OpenFile("bot.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Errorf("error opening log file: %v", err)
	} else {
		logger.SetOutput(file)
	}

	discord.Init()

	// if logger.LastLog().Level != "info" || logger.LastLog().Messag != "bot is running. Press CTRL + C to exit." {
	// 	t.Errorf("Init() did not log the appropriate message")
	// }

}
