package hardwareid

import "testing"

func TestID(t *testing.T) {
	got, err := ID()
	if err != nil {
		t.Error(err)
	}
	if got == "" {
		t.Error("Got empty hardware id")
	}
}

func TestProtectedID(t *testing.T) {
	id, err := ID()
	if err != nil {
		t.Error(err)
	}
	hash, err := ProtectedID("app.id")
	if err != nil {
		t.Error(err)
	}
	if hash == "" {
		t.Error("Got empty hardware id hash")
	}
	if id == hash {
		t.Error("id and hashed id are the same")
	}
}
