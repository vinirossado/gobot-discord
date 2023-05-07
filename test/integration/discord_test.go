package discord_test

import (
	"testing"

	"github.com/arthur404dev/gobot-discord/discord"
	"github.com/bwmarrin/discordgo"
)

func TestInit(t *testing.T) {

	session, err := discordgo.New("Bot init discord")
	if err != nil {
		t.Fatalf("error creating discord session: %v", err)
	}
	defer session.Close()

	go discord.Init()

	testCommand := &discordgo.ApplicationCommand{
		Name:        "testcommand",
		Description: "A test command",
	}

	_, err = session.ApplicationCommandCreate("", session.State.User.ID, testCommand)
	if err != nil {
		t.Fatalf("error registering test command: %v", err)
	}

	interaction := &discordgo.InteractionCreate{
		Interaction: &discordgo.Interaction{
			Type: discordgo.InteractionApplicationCommand,
			Data: &discordgo.ApplicationCommandInteractionData{
				Name: "testcommand",
			},
		},
	}

	session.State.SessionID = "testsessionid"
	session.State.User = &discordgo.User{
		ID: session.State.User.ID,
	}

	session.State.Guilds = append(session.State.Guilds, &discordgo.Guild{
		ID: "testguildid",
	})

	session.State.Guilds[0].Members = append(session.State.Guilds[0].Members, &discordgo.Member{
		User: &discordgo.User{
			ID: session.State.User.ID,
		},
	})

	session.State.Ready = discordgo.Ready{
		User: &discordgo.User{
			ID: session.State.User.ID,
		},
	}
	session.State.SessionID = "testsessionid"

	err = session.InteractionCreate{
		Interaction: interaction,
	}

	if err != nil {
		t.Fatalf("error executing test command: %v", err)
	}

}
