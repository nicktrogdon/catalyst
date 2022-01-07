package service

import (
	"context"

	"github.com/arangodb/go-driver"

	"github.com/SecurityBrewery/catalyst/bus"
	"github.com/SecurityBrewery/catalyst/database"
	"github.com/SecurityBrewery/catalyst/database/busdb"
	"github.com/SecurityBrewery/catalyst/generated/model"
	"github.com/SecurityBrewery/catalyst/storage"
)

type Service struct {
	bus      *bus.Bus
	database *database.Database
	settings *model.Settings
	storage  *storage.Storage
}

func New(bus *bus.Bus, database *database.Database, storage *storage.Storage, settings *model.Settings) (*Service, error) {
	return &Service{database: database, bus: bus, settings: settings, storage: storage}, nil
}

func (s *Service) publishRequest(ctx context.Context, err error, function string, ids []driver.DocumentID) {
	if err != nil {
		return
	}
	if ids != nil {
		userID := "unknown"
		user, ok := busdb.UserFromContext(ctx)
		if ok {
			userID = user.ID
		}

		go s.bus.PublishRequest(userID, function, ids)
	}
}
