package businesslogic

import (
	"fmt"

	"github.com/Nastena2002/sacc/chaincode/internal/repository"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SimpleService struct {
	Repository *repository.SimpleRepository
	Ctx        contractapi.TransactionContextInterface
}

func NewSimpleService(repo *repository.SimpleRepository, ctx contractapi.TransactionContextInterface) *SimpleService {
	return &SimpleService{
		Repository: repo,
		Ctx:        ctx,
	}
}

type Service struct {
	simpleService *SimpleService
	repoService   *repository.Service
	ctx           contractapi.TransactionContextInterface
}

func NewService(repoService *repository.Service) *Service {
	return &Service{
		simpleService: NewSimpleService(repoService.SimpleRepo(), nil),
		repoService:   repoService,
		ctx:           nil,
	}
}

func (s *Service) Simple() *SimpleService {
	return s.simpleService
}

func (s *SimpleService) Get(key string) (string, error) {
	s.Repository.SetContext(s.Ctx)

	asset, err := s.Repository.GetAsset(key)
	if err != nil {
		return "false", fmt.Errorf("failed to get account %s: %v", key, err)
	}

	return asset.Value, nil
}

func (s *SimpleService) Set(key string, value string) error {
	s.Repository.SetContext(s.Ctx)

	asset := repository.SimpleAsset{
		Key:   key,
		Value: value,
	}
	err := s.Repository.SetAsset(asset)
	if err != nil {
		return fmt.Errorf("failed to create account A: %v", err)
	}

	return nil
}
