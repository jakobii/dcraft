package app

import (
	"crypto/ed25519"
	"encoding/hex"
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Start is the main entry point to the app
func Start() {
	var port uint64
	var discordPublicKey ed25519.PublicKey

	if x, ok := os.LookupEnv("DCRAFT_DISCORD_PUB_KEY"); ok {
		k, err := hex.DecodeString(x)
		if err != nil {
			log.Fatal("invalid pub key", err)
		}
		discordPublicKey = ed25519.PublicKey(k)
	}

	if x, ok := os.LookupEnv("DCRAFT_PORT"); ok {
		p, err := strconv.ParseUint(x, 10, 64)
		if err != nil {
			log.Fatal("invalid port", err)
		}
		port = p
	}

	p := flag.Uint64("port", 80, "http port")
	flag.Parse()
	if isFlagPassed("port") {
		port = *p
	}
	if port == 0 {
		port = 80
	}
	log.Println("starting server on port", port)

	s := Server{
		discordPublicKey: discordPublicKey,
	}
	http.ListenAndServe(
		":"+strconv.FormatUint(port, 10),
		s.Handler(),
	)
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
