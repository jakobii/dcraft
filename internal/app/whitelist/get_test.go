package whitelist

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/jakobii/dcraft/internal/fileproxy"
	"github.com/jakobii/dcraft/internal/minecraft"
)

func TestGet(t *testing.T) {
	a := []minecraft.User{
		{UUID: uuid.New(), Name: "Alex"},
		{UUID: uuid.New(), Name: "Bob"},
		{UUID: uuid.New(), Name: "Chris"},
		{UUID: uuid.New(), Name: "Dave"},
	}

	j, err := json.Marshal(a)
	check(t, err)

	list := Whitelist{
		Conenter: fileproxy.NewMockContenter(j),
	}
	b, err := list.Get(context.TODO())
	check(t, err)

	for _, user := range a {
		if !hasUser(b, user) {
			t.Fatal(wantGot(a, b))
		}
	}

}

func wantGot(want, got interface{}) string {
	return fmt.Sprintf("want: %v, got: %v", want, got)
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func hasUser(users []minecraft.User, user minecraft.User) bool {
	for _, x := range users {
		if x.UUID == user.UUID && x.Name == user.Name {
			return true
		}
	}
	return false
}
