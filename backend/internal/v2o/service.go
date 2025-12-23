package v2o

import (
	"context"
	"fmt"
	"github.com/ryudokung/Project-0/backend/internal/blockchain"
	"github.com/ryudokung/Project-0/backend/internal/mech"

	"github.com/google/uuid"
)

type Service struct {
	mechRepo   mech.Repository
	bcProvider blockchain.Provider
}

func NewService(repo mech.Repository, bc blockchain.Provider) *Service {
	return &Service{
		mechRepo:   repo,
		bcProvider: bc,
	}
}

func (s *Service) BridgeToChain(ctx context.Context, mechID uuid.UUID, ownerAddress string) (string, error) {
	// 1. Fetch the virtual mech
	m, err := s.mechRepo.GetByID(ctx, mechID)
	if err != nil {
		return "", fmt.Errorf("failed to find mech: %w", err)
	}

	if m.Status == mech.StatusMinted {
		return "", fmt.Errorf("mech is already minted")
	}

	// 2. Prepare mint request
	req := blockchain.MintRequest{
		OwnerAddress: ownerAddress,
		MetadataURI:  m.ImageURL,
		Stats: map[string]int{
			"hp":      m.Stats.HP,
			"attack":  m.Stats.Attack,
			"defense": m.Stats.Defense,
		},
	}

	// 3. Call blockchain provider
	txHash, err := s.bcProvider.MintMech(ctx, req)
	if err != nil {
		return "", fmt.Errorf("blockchain minting failed: %w", err)
	}

	// 4. Update status in DB
	m.Status = mech.StatusMinted
	m.TokenID = &txHash // Using txHash as temporary TokenID for mock
	if err := s.mechRepo.Update(ctx, m); err != nil {
		return "", fmt.Errorf("failed to update mech status: %w", err)
	}

	return txHash, nil
}
