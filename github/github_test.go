package github

import (
	"bytes"
	"io"
	"testing"
)

func TestNewRequest_GetBody(t *testing.T) {
	c := NewClient(nil)
	inBody := &User{Login: String("l")}
	req, err := c.NewRequest("POST", "user", inBody)
	if err != nil {
		t.Fatalf("NewRequest returned unexpected error: %v", err)
	}

	if req.GetBody == nil {
		t.Fatal("req.GetBody is nil, expected it to be populated")
	}

	// Read body the first time
	body1, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatalf("failed to read req.Body: %v", err)
	}

	// Get a new body reader
	newBodyReader, err := req.GetBody()
	if err != nil {
		t.Fatalf("req.GetBody() returned error: %v", err)
	}
	defer newBodyReader.Close()

	// Read body the second time
	body2, err := io.ReadAll(newBodyReader)
	if err != nil {
		t.Fatalf("failed to read body from GetBody: %v", err)
	}

	if !bytes.Equal(body1, body2) {
		t.Errorf("expected bodies to be equal, got %q and %q", body1, body2)
	}
}
