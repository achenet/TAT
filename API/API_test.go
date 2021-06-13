package API

import "testing"

func TestSendRequest(t *testing.T) {

	err := SendRequest()
	if err != nil {
		t.Error("Error sending request", err.Error())
	}
}
