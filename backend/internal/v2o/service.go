package v2o

import (
	"context"
	"fmt"
	"github.com/ryudokung/Project-0/backend/internal/blockchain"
	"github.com/ryudokung/Project-0/backend/internal/vehicle"

	"github.com/google/uuid"
)

type Service struct {
	vehicleRepo vehicle.Repository
	bcProvider  blockchain.Provider
}

func NewService(repo vehicle.Repository, bc blockchain.Provider) *Service {
	return &Service{
		vehicleRepo: repo,
		bcProvider:  bc,
	}
}

func (s *Service) BridgeToChain(ctx context.Context, vehicleID uuid.UUID, ownerAddress string) (string, error) {
	// 1. Fetch the virtual vehicle
	v, err := s.vehicleRepo.GetByID(ctx, vehicleID)
	if err != nil {
		return "", fmt.Errorf("failed to find vehicle: %w", err)
	}

	if v.Status == vehicle.StatusMinted {
		return "", fmt.Errorf("vehicle is already minted")
	}

	// 2. Prepare mint request
	req := blockchain.MintRequest{
		OwnerAddress: ownerAddress,
		MetadataURI:  v.ImageURL,
		Stats: map[string]int{
			"hp":      v.Stats.HP,
			"attack":  v.Stats.Attack,
			"defense": v.Stats.Defense,
		},
	}

	// 3. Call blockchain provider
	txHash, err := s.bcProvider.MintVehicle(ctx, req)
	if err != nil {
		return "", fmt.Errorf("blockchain minting failed: %w", err)
	}

	// 4. Update status in DB
	v.Status = vehicle.StatusMinted
	v.TokenID = &txHash // Using txHash as temporary TokenID for mock
	if err := s.vehicleRepo.Update(ctx, v); err != nil {
		return "", fmt.Errorf("failed to update vehicle status: %w", err)
	}

	return txHash, nil
}
