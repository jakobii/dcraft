package app

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

const pathDiscord = "/discord"

func (s *Server) routeDiscord(r *mux.Router) {
	r.HandleFunc(pathDiscord, s.handleDiscord).Methods(http.MethodPost)
}

func (s *Server) handleDiscord(w http.ResponseWriter, r *http.Request) {
	typ, ok := s.requireDiscordInteraction(w, r)
	if !ok {
		return
	}
	switch typ {
	case Ping:
		WriteJSON(w, http.StatusOK, InteractionResponse{Type: Pong})
	case ApplicationCommand:
		WriteJSON(w, http.StatusOK, InteractionResponse{
			Type: Pong,
			Data: &InteractionApplicationCommandCallbackData{
				Content: "Command Received 🤓",
			},
		})
	}
}

// requireDiscordInteraction will write back an error to the client
// if the Discord Interaction is not valid.
func (s *Server) requireDiscordInteraction(w http.ResponseWriter, r *http.Request) (typ InteractionType, ok bool) {
	if !verifyDiscordInteraction(s.discordPublicKey, r) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, `invalid request signature`)
		return 0, false
	}
	return s.requireDiscordInteractionType(w, r)
}

// https://discord.com/developers/docs/interactions/slash-commands#security-and-authorization
func verifyDiscordInteraction(key ed25519.PublicKey, r *http.Request) bool {
	signature := r.Header.Get("X-Signature-Ed25519")
	if signature == "" {
		return false
	}
	timestamp := r.Header.Get("X-Signature-Timestamp")
	if timestamp == "" {
		return false
	}

	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	if len(sig) != ed25519.SignatureSize {
		return false
	}

	var msg bytes.Buffer
	msg.WriteString(timestamp)

	defer r.Body.Close()
	var body bytes.Buffer

	// at the end of the function, copy the original body back into the request
	defer func() {
		r.Body = ioutil.NopCloser(&body)
	}()

	// copy body into buffers
	_, err = io.Copy(&msg, io.TeeReader(r.Body, &body))
	if err != nil {
		return false
	}

	return ed25519.Verify(key, msg.Bytes(), sig)
}

func (s *Server) requireDiscordInteractionType(w http.ResponseWriter, r *http.Request) (InteractionType, bool) {
	body := readRequestBody(r)
	typ, err := unmarshalDiscordInteractionType(body)
	if err != nil {
		http.Error(w, fmt.Sprintf(`missing interaction type: %v`, err), http.StatusBadRequest)
		return 0, false
	}
	return typ, true
}

func unmarshalDiscordInteractionType(body []byte) (InteractionType, error) {
	x := struct {
		Type InteractionType
	}{}
	err := json.Unmarshal(body, &x)
	if err != nil {
		return 0, err
	}
	switch t := x.Type; t {
	case Ping, ApplicationCommand:
		return t, nil
	default:
		return 0, fmt.Errorf("invalid interaction type: %v", t)
	}
}

// InteractionType ...
type InteractionType int

// InteractionTypes
// https://discord.com/developers/docs/interactions/slash-commands#interaction-interactiontype
const (
	Ping               = 1
	ApplicationCommand = 2
)

// InteractionResponseType ...
type InteractionResponseType int

// Discord InteractionResponseTypes
// https://discord.com/developers/docs/interactions/slash-commands#interaction-response-interactionresponsetype
const (
	// ACK a Ping
	Pong InteractionResponseType = 1
	// ACK a command without sending a message, eating the user's input
	Acknowledge InteractionResponseType = 2
	// respond with a message, eating the user's input
	ChannelMessage InteractionResponseType = 3
	// respond with a message, showing the user's input
	ChannelMessageWithSource InteractionResponseType = 4
	// ACK a command without sending a message, showing the user's input
	AcknowledgeWithSource InteractionResponseType = 5
)

// InteractionResponse ...
// https://discord.com/developers/docs/interactions/slash-commands#interaction-response
type InteractionResponse struct {
	Type InteractionResponseType                    `json:"type"`
	Data *InteractionApplicationCommandCallbackData `json:"data,omitempty"`
}

// InteractionApplicationCommandCallbackData ...
// https://discord.com/developers/docs/interactions/slash-commands#interaction-response-interactionapplicationcommandcallbackdata
type InteractionApplicationCommandCallbackData struct {
	TTS             *bool            `json:"tts,omitempty"`              // is the response TTS
	Content         string           `json:"content"`                    // message content
	Embeds          []Embed          `json:"embeds,omitempty"`           // supports up to 10 embeds
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"` // allowed mentions object
}

// Embed ...
// https://discord.com/developers/docs/resources/channel#embed-object-embed-structure
type Embed struct {
	// TODO
}

// AllowedMentions ...
// https://discord.com/developers/docs/resources/channel#embed-object-embed-structure
type AllowedMentions struct {
	// TODO
}
