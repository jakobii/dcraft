package cfg

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jakobii/dcraft/internal/app"
	"github.com/jakobii/dcraft/internal/app/auth"
	"github.com/jakobii/dcraft/internal/app/whitelist"
	"gopkg.in/yaml.v2"
)

// Parse app server options
func Parse() ([]app.ServerOption, error) {
	var port uint64
	var discordPublicKey string
	var whitelister app.Whitelister
	var super string
	var creds []auth.Credential

	// ENVs
	if x, ok := os.LookupEnv("DCRAFT_SUPER"); ok {
		super = x
	}
	if path, ok := os.LookupEnv("DCRAFT_CREDENTIALS"); ok {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("fail to read credentials file: %v", err)
		}
		switch filepath.Ext(path) {
		case ".yaml", ".yml":
			err := yaml.Unmarshal(b, &creds)
			if err != nil {
				return nil, fmt.Errorf("fail to read credentials file: %v", err)
			}
		case ".json":
			err := json.Unmarshal(b, &creds)
			if err != nil {
				return nil, fmt.Errorf("fail to read credentials file: %v", err)
			}
		}
	}

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

	// INITS
	authCache := auth.NewCache(auth.WithSuper(super))
	authCache.PutCredential(creds...)

	// OPTS
	optDiscordPublicKey, err := app.WithDiscordPublicKey(discordPublicKey)
	if err != nil {
		return nil, err
	}
	return []app.ServerOption{
		app.WithAuthenticator(authCache),
		app.WithAuthorizer(authCache),
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
