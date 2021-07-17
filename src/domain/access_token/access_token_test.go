package access_token

import "testing"

func TestGetAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("new access token should not be expired")
	}

	if at.AccessToken != "" {
		t.Error("new access token should not have access token id")
	}

	if at.UserId != 0 {
		t.Error("new access token should not have an associated user id")
	}

}
