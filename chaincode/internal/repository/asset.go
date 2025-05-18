package repository

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SimpleAsset struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type SimpleRepository struct {
	Ctx contractapi.TransactionContextInterface
}

func NewSimpleRepository(ctx contractapi.TransactionContextInterface) *SimpleRepository {
	return &SimpleRepository{
		Ctx: ctx,
	}
}

func (r *SimpleRepository) SetAsset(asset SimpleAsset) error {
	if r.Ctx == nil {
		return fmt.Errorf("transaction context is nil")
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}
	return r.Ctx.GetStub().PutState(asset.Key, assetJSON)
}

func (r *SimpleRepository) GetAsset(key string) (*SimpleAsset, error) {
	if r.Ctx == nil {
		return nil, fmt.Errorf("transaction context is nil")
	}
	assetJSON, err := r.Ctx.GetStub().GetState(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("asset %s does not exist", key)
	}

	var asset SimpleAsset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}
	return &asset, nil
}

func (r *SimpleRepository) SetContext(ctx contractapi.TransactionContextInterface) {
	r.Ctx = ctx
}
