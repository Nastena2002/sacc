package handler

import (
	"github.com/Nastena2002/sacc/internal/businesslogic"
	"github.com/Nastena2002/sacc/internal/repository"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SimpleChaincode struct {
	contractapi.Contract
	simpleService *businesslogic.SimpleService
	bService      *businesslogic.Service
	repoService   *repository.Service
}

func NewService(businessService *businesslogic.Service, repoService *repository.Service) *SimpleChaincode {
	return NewAssetHandler(businessService, repoService)
}

func NewAssetHandler(businessService *businesslogic.Service, repoService *repository.Service) *SimpleChaincode {
	return &SimpleChaincode{
		simpleService: businessService.Simple(),
		bService:      businessService,
		repoService:   repoService,
	}
}

func (s *SimpleChaincode) Set(ctx contractapi.TransactionContextInterface, Key string, Value string) error {
	s.simpleService.Ctx = ctx
	s.repoService.SetContext(ctx)

	return s.simpleService.Set(Key, Value)
}

func (h *SimpleChaincode) Get(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	h.simpleService.Ctx = ctx
	h.repoService.SetContext(ctx)

	value, err := h.simpleService.Get(key)
	if err != nil {
		return "", err
	}

	return value, nil
}
