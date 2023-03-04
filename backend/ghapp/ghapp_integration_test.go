package ghapp

import (
	"context"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"log"
	"net/http"
	"testing"
)

type mockFileToCommit struct {
}

func (m mockFileToCommit) Path() string {
	return "test.txt"
}

func (m mockFileToCommit) CommitMessage() string {
	return "testing"
}

func (m mockFileToCommit) AsBytes() ([]byte, error) {
	return []byte("test"), nil
}

func TestGitHubIntegration(t *testing.T) {

	transport, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, 0, 0, "../key.pem")
	if err != nil {
		log.Fatalf("Error initializing application: %s", err.Error())
		return
	}
	ghApp := New(transport, "test")
	_, err = ghApp.MakePullRequest(context.Background(), mockFileToCommit{})
	if err != nil {
		t.Fail()
		log.Printf("an error occured: %s", err.Error())
	}
}
