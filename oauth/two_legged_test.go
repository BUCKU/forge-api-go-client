package oauth_test

import (
	"testing"
	"os"
	"github.com/apprentice3d/forge-api-go-client/oauth"
)

func TestAuthenticate(t *testing.T) {

	clientID :=os.Getenv("FORGE_CLIENT_ID")
	clientSecret :=os.Getenv("FORGE_CLIENT_SECRET")

	if len(clientID) == 0 || len(clientSecret) == 0 {
		t.Fatalf("Could not get from env the Forge secrets")
	}

	t.Run("Valid Forge Secrets", func(t *testing.T) {
		authenticator := oauth.NewTwoLeggedClient(clientID, clientSecret)

		bearer, err := authenticator.Authenticate("data:read")

		if err != nil {
			t.Error(err.Error())
		}

		if len(bearer.AccessToken) == 0 {
			t.Errorf("Wrong bearer content: %v", bearer)
		}
	})

	t.Run("Invalid Forge Secrets", func(t *testing.T) {
		authenticator := oauth.NewTwoLeggedClient("", clientSecret)

		bearer, err := authenticator.Authenticate("data:read")

		if err == nil {
			t.Errorf("Expected to fail due to wrong credentials, but got %v", bearer)
		}

		if len(bearer.AccessToken) != 0 {
			t.Errorf("expected to not receive a token, but received: %s", bearer.AccessToken)
		}
	})


	t.Run("Invalid scope", func(t *testing.T) {
		authenticator := oauth.NewTwoLeggedClient(clientID, clientSecret)

		bearer, err := authenticator.Authenticate("data:improvise")

		if err == nil {
			t.Errorf("Expected to fail due to wrong scope, but got %v\n", bearer)
		}

		if len(bearer.AccessToken) != 0 {
			t.Errorf("expected to not receive a token, but received: %s", bearer.AccessToken)
		}
	})

	t.Run("Invalid or unreachable host", func(t *testing.T) {
		authenticator := oauth.NewTwoLeggedClient(clientID, clientSecret)
		authenticator.Host = "http://localhost"

		bearer, err := authenticator.Authenticate("data:read")

		if err == nil {
			t.Errorf("Expected to fail due to wrong host, but got %v\n", bearer)
		}

		if len(bearer.AccessToken) != 0 {
			t.Errorf("expected to not receive a token, but received: %s", bearer.AccessToken)
		}
	})

}