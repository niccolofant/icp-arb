package icp

import (
	"fmt"
	"net/http"

	"github.com/aviate-labs/agent-go"
	"github.com/aviate-labs/agent-go/identity"
)

type Agent struct {
	raw *agent.Agent
}

func NewAgent(id identity.Identity, httpClient *http.Client) (*Agent, error) {
	if httpClient == nil {
		return nil, fmt.Errorf("nil httpClient")
	}

	agent, err := agent.New(agent.Config{
		ClientConfig: []agent.ClientOption{
			agent.WithHttpClient(httpClient),
		},
		Identity: id,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create icp agent: %w", err)
	}

	return &Agent{raw: agent}, nil
}

func (a *Agent) Sender() Principal {
	return NewPrincipal(a.raw.Sender())
}

func (a Agent) Raw() *agent.Agent {
	return a.raw
}
