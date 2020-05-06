package identity

import (
	"github.com/konveyor/controller/pkg/condition"
	"github.com/konveyor/mig-controller/pkg/apis/migration/v1alpha1"
)

//
// Types
const ()

//
// Categories
const (
	Advisory = condition.Advisory
	Critical = condition.Critical
	Error    = condition.Error
	Warn     = condition.Warn
)

// Reasons
const (
	NotSet   = "NotSet"
	NotFound = "NotFound"
)

// Statuses
const (
	True  = condition.True
	False = condition.False
)

// Messages
const (
	ReadyMessage = "The identity is ready."
)

// Validate the plan resource.
func (r Reconciler) validate(identity *v1alpha1.MigIdentity) error {
	id, err := identity.BuildIdentity(r.Client)
	if err != nil {
		log.Trace(err)
		return err
	}

	return nil
}

func (r Reconciler) validateToken(identity *v1alpha1.MigIdentity) error {
	_, err := identity.GetToken(r.Client)
	if err != nil {
		return err
	}
	return nil
}
