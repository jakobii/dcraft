package cfg

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/jakobii/dcraft/internal/app"
	"github.com/jakobii/dcraft/internal/app/whitelist"
)

// Get app server options
func Get() ([]app.ServerOption, error) {
	var port uint64
	var discordPublicKey string
	var whitelister app.Whitelister

	// ENVs
	if x, ok := os.LookupEnv("DCRAFT_DISCORD_PUB_KEY"); ok {
		discordPublicKey = x
	} else {
		return nil, fmt.Errorf("missing discord public key")
	}

	if x, ok := os.LookupEnv("DCRAFT_MC_WHITELIST"); ok {
		whitelister = whitelist.New(x)
	} else {
		return nil, fmt.Errorf("missing whitelist")
	}

	if x, ok := os.LookupEnv("DCRAFT_PORT"); ok {
		p, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			return nil, err
		}
		port = p
	}

	// FLAGS
	p := flag.Uint64("port", 80, "http port")
	flag.Parse()
	if isFlagPassed("port") {
		port = *p
	}
	if port == 0 {
		port = 80
	}

	// OPTS
	optDiscordPublicKey, err := app.WithDiscordPublicKey(discordPublicKey)
	if err != nil {
		return nil, err
	}
	return []app.ServerOption{
		app.WithPort(port),
		app.WithWhitelist(whitelister),
		optDiscordPublicKey,
	}, nil
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
