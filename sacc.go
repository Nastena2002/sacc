package main

import (
	"github.com/Nastena2002/sacc/internal"
	businesslogic "github.com/Nastena2002/sacc/internal/businessLogic"
	"github.com/Nastena2002/sacc/internal/handler"
	"github.com/Nastena2002/sacc/internal/repository"
	"go.uber.org/fx"
)

func main_sacc() {
	fx.New(
		fx.Provide(handler.NewService),
		fx.Provide(businesslogic.NewService),
		fx.Provide(repository.NewService),
		fx.Provide(internal.NewApp),

		fx.Invoke(func(*internal.App) {}),
	).Run()
}
