package pipeline

import (
	"context"
)

// Stager interface
type Stager interface {
	Execute(context.Context)
}
