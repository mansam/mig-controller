package identity

import "sigs.k8s.io/controller-runtime/pkg/predicate"

type IdentityPredicate struct {
	predicate.Funcs
}
