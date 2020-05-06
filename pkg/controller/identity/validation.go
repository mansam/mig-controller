package identity

import (
	cnd "github.com/konveyor/controller/pkg/condition"
	"github.com/konveyor/mig-controller/pkg/apis/migration/v1alpha1"
)

//
// Types
const ()

//
// Categories
const (
	Advisory = cnd.Advisory
	Critical = cnd.Critical
	Error    = cnd.Error
	Warn     = cnd.Warn
)

// Reasons
const (
	NotSet   = "NotSet"
	NotFound = "NotFound"
)

// Statuses
const (
	True  = cnd.True
	False = cnd.False
)

// Messages
const (
	ReadyMessage = "The identity is ready."
)

// Validate the plan resource.
func (r Reconciler) validate(plan *v1alpha1.MigIdentity) error {

	return nil
}
