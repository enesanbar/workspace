package get

import (
	"context"
	"errors"

	"github.com/enesanbar/workspace/golang/projects/acme/internal/logging"
	"github.com/enesanbar/workspace/golang/projects/acme/internal/modules/data"
)

var (
	// error thrown when the requested person is not in the database
	errPersonNotFound = errors.New("person not found")
)

// Config is the configuration for Getter
type Config interface {
	Logger() logging.Logger
	DataDSN() string
}

//go:generate mockery -name=myLoader -case underscore -testonly -inpkg
type myLoader interface {
	Load(ctx context.Context, ID int) (*data.Person, error)
}

// Getter will attempt to load a person.
// It can return an error caused by the data layer or when the requested person is not found
type Getter struct {
	cfg  Config
	data myLoader
}

// NewGetter creates and initializes a Getter
func NewGetter(cfg Config) *Getter {
	return &Getter{
		cfg: cfg,
	}
}

// Do will perform the get
func (g *Getter) Do(ctx context.Context, ID int) (*data.Person, error) {
	// load person from the data layer
	person, err := g.getLoader().Load(ctx, ID)
	if err != nil {
		if err == data.ErrNotFound {
			// By converting the error we are hiding the implementation details from our users.
			return nil, errPersonNotFound
		}
		return nil, err
	}

	return person, err
}

func (g *Getter) getLoader() myLoader {
	if g.data == nil {
		g.data = data.NewDAO(g.cfg)
	}

	return g.data
}
