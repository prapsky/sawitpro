// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"

	"github.com/prapsky/sawitpro/entity"
)

type RepositoryInterface interface {
	Insert(ctx context.Context, input entity.User) (RegistrationOutput, error)
}
