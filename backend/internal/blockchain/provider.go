package blockchain

import (
	"context"
	"math/big"
)

type MintRequest struct {
	OwnerAddress string
	MetadataURI  string
	Stats        map[string]int
}

type Provider interface {
	MintMech(ctx context.Context, req MintRequest) (string, error) // Returns Transaction Hash or Token ID
	GetBalance(ctx context.Context, address string) (*big.Int, error)
}

type BaseProvider struct {
	// Add RPC URL, Private Key, Contract Address here
}

func NewBaseProvider() *BaseProvider {
	return &BaseProvider{}
}

func (p *BaseProvider) MintMech(ctx context.Context, req MintRequest) (string, error) {
	// Placeholder for actual Ethers/Web3 implementation
	// In a real scenario, we'd use go-ethereum to call the contract
	return "0x-mock-tx-hash", nil
}

func (p *BaseProvider) GetBalance(ctx context.Context, address string) (*big.Int, error) {
	return big.NewInt(0), nil
}
