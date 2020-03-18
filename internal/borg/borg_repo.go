package borg

import (
	"context"
	"fmt"
	"github.com/bboortz/goborg/pkg/appcontext"
	"time"
)

var currentId string

var BorgRepo Borgs

// Give us some seed data
func init() {
	// RepoCreateBorg(Borg{Name: "Host meetup"})
}

func RepoFindBorg(id string) Borg {
	for _, b := range BorgRepo {
		if b.Id == id {
			return b
		}
	}
	// return empty Borg if not found
	return Borg{}
}

//this is bad, I don't think it passes race condtions
func RepoAddBorg(ctx context.Context, b Borg) Borg {
	logger := appcontext.Logger(ctx)

	for i, ib := range BorgRepo {
		if ib.Id == b.Id {
			logger.Debug("Borg overwritten: " + b.Id)
			b.LastSeen = time.Now()
			BorgRepo[i] = b
			return b
		}
	}

	b.LastSeen = time.Now()
	BorgRepo = append(BorgRepo, b)
	return b
}

func RepoDestroyBorg(id string) error {
	for i, b := range BorgRepo {
		if b.Id == id {
			BorgRepo = append(BorgRepo[:i], BorgRepo[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Borg with id of %s to delete", id)
}
