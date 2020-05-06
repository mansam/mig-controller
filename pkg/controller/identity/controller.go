package identity

import (
	"context"
	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/konveyor/mig-controller/pkg/apis/migration/v1alpha1"
	"github.com/konveyor/mig-controller/pkg/logging"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logging.WithName("identity")

var _ reconcile.Reconciler = &Reconciler{}

type Reconciler struct {
	client.Client
	scheme *runtime.Scheme
}

func Add(mgr manager.Manager) error {
	reconciler := &Reconciler{
		Client: mgr.GetClient(),
		scheme: mgr.GetScheme(),
	}

	c, err := controller.New(
		"identity-controller",
		mgr,
		controller.Options{
			Reconciler: reconciler,
		})
	if err != nil {
		log.Trace(err)
		return err
	}
	err = c.Watch(
		&source.Kind{Type: &v1alpha1.MigIdentity{}},
		&handler.EnqueueRequestForObject{},
		&IdentityPredicate{})
	if err != nil {
		log.Trace(err)
		return err
	}
	return nil
}

func (r *Reconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	var err error

	// Reset the logger
	log.Reset()

	// Fetch the CR.
	identity := &v1alpha1.MigIdentity{}
	err = r.Get(context.TODO(), request.NamespacedName, identity)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		log.Trace(err)
		return reconcile.Result{}, err
	}

	err = r.validate(identity)
	if err != nil {
		log.Trace(err)
		return reconcile.Result{Requeue: true}, nil
	}

	if !identity.Status.HasBlockerCondition() {
		identity.Status.SetReady(true, ReadyMessage)
	}

	identity.MarkReconciled()
	err = r.Update(context.TODO(), identity)
	if err != nil {
		log.Trace(err)
		return reconcile.Result{Requeue: true}, nil
	}

	// Done
	return reconcile.Result{}, nil
}
