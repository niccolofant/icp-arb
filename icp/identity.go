package icp

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aviate-labs/agent-go/identity"
)

func LoadIntentity(path string) (identity.Identity, error) {
	if path == "" {
		wd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("failed to get working directory: %w", err)
		}

		path = filepath.Join(wd, "identity.pem")
	}

	pem, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read identity file: %w", err)
	}

	iid, err := identity.NewSecp256k1IdentityFromPEMWithoutParameters(pem)
	if err != nil {
		return nil, fmt.Errorf("failed to create identity from pem: %w", err)
	}

	return iid, nil
}
