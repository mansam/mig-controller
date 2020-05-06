package v1alpha1

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/konveyor/mig-controller/pkg/settings"
)

const (
	TouchAnnotation = "openshift.io/touch"
	VeleroNamespace = "openshift-migration"
)

// Migration application CR.
type MigResource interface {
	// Get a map containing the correlation label.
	// Correlation labels are used to track any resource
	// created by the controller.  The includes the application
	// (global) label and the resource label.
	GetCorrelationLabels() map[string]string
	// Get the resource correlation label.
	// The label is used to track resources created by the
	// controller that is related to this resource.
	GetCorrelationLabel() (string, string)
	// Get the resource namespace.
	GetNamespace() string
	// Get the resource name.
	GetName() string
	// Mark the resource as having been reconciled.
	// Updates the ObservedDigest.
	// Update the touch annotation. This ensures that the resource
	// is changed on every reconcile. This needed to support
	// remote watch event propagation on OCP4.
	MarkReconciled()
	// Get whether the resource has been reconciled.
	HasReconciled() bool
	// Get whether the resource lives in the sandbox namespace.
	InSandbox() bool
	// Get whether the resource lives in the privileged namespace.
	InPrivileged() bool
}

// Plan
func (r *MigPlan) GetCorrelationLabels() map[string]string {
	key, value := r.GetCorrelationLabel()
	return map[string]string{
		PartOfLabel: Application,
		key:         value,
	}
}

func (r *MigPlan) GetCorrelationLabel() (string, string) {
	return CorrelationLabel(r, r.UID)
}

func (r *MigPlan) GetNamespace() string {
	return r.Namespace
}

func (r *MigPlan) GetName() string {
	return r.Name
}

func (r *MigPlan) MarkReconciled() {
	uuid, _ := uuid.NewUUID()
	if r.Annotations == nil {
		r.Annotations = map[string]string{}
	}
	r.Annotations[TouchAnnotation] = uuid.String()
	r.Status.ObservedDigest = digest(r.Spec)
}

func (r *MigPlan) HasReconciled() bool {
	return r.Status.ObservedDigest == digest(r.Spec)
}

func (r *MigPlan) InSandbox() bool {
	return r.GetNamespace() == settings.Settings.Namespace.Sandbox
}

func (r *MigPlan) InPrivileged() bool {
	return r.GetNamespace() == settings.Settings.Namespace.Privileged
}

// Storage
func (r *MigStorage) GetCorrelationLabels() map[string]string {
	key, value := r.GetCorrelationLabel()
	return map[string]string{
		PartOfLabel: Application,
		key:         value,
	}
}

func (r *MigStorage) GetCorrelationLabel() (string, string) {
	return CorrelationLabel(r, r.UID)
}

func (r *MigStorage) GetNamespace() string {
	return r.Namespace
}

func (r *MigStorage) GetName() string {
	return r.Name
}

func (r *MigStorage) MarkReconciled() {
	uuid, _ := uuid.NewUUID()
	if r.Annotations == nil {
		r.Annotations = map[string]string{}
	}
	r.Annotations[TouchAnnotation] = uuid.String()
	r.Status.ObservedDigest = digest(r.Spec)
}

func (r *MigStorage) HasReconciled() bool {
	return r.Status.ObservedDigest == digest(r.Spec)
}

func (r *MigStorage) InSandbox() bool {
	return r.GetNamespace() == settings.Settings.Namespace.Sandbox
}

func (r *MigStorage) InPrivileged() bool {
	return r.GetNamespace() == settings.Settings.Namespace.Privileged
}

// Cluster
func (r *MigCluster) GetCorrelationLabels() map[string]string {
	key, value := r.GetCorrelationLabel()
	return map[string]string{
		PartOfLabel: Application,
		key:         value,
	}
}

func (r *MigCluster) GetCorrelationLabel() (string, string) {
	return CorrelationLabel(r, r.UID)
}

func (r *MigCluster) GetNamespace() string {
	return r.Namespace
}

func (r *MigCluster) GetName() string {
	return r.Name
}

func (r *MigCluster) MarkReconciled() {
	uuid, _ := uuid.NewUUID()
	if r.Annotations == nil {
		r.Annotations = map[string]string{}
	}
	r.Annotations[TouchAnnotation] = uuid.String()
	r.Status.ObservedDigest = digest(r.Spec)
}

func (r *MigCluster) HasReconciled() bool {
	return r.Status.ObservedDigest == digest(r.Spec)
}

func (r *MigCluster) InSandbox() bool {
	return r.GetNamespace() == settings.Settings.Namespace.Sandbox
}

func (r *MigCluster) InPrivileged() bool {
	return r.GetNamespace() == settings.Settings.Namespace.Privileged
}

// Migration
func (r *MigMigration) GetCorrelationLabels() map[string]string {
	key, value := r.GetCorrelationLabel()
	return map[string]string{
		PartOfLabel: Application,
		key:         value,
	}
}

func (r *MigMigration) GetCorrelationLabel() (string, string) {
	return CorrelationLabel(r, r.UID)
}

func (r *MigMigration) GetNamespace() string {
	return r.Namespace
}

func (r *MigMigration) GetName() string {
	return r.Name
}

func (r *MigMigration) MarkReconciled() {
	uuid, _ := uuid.NewUUID()
	if r.Annotations == nil {
		r.Annotations = map[string]string{}
	}
	r.Annotations[TouchAnnotation] = uuid.String()
	r.Status.ObservedDigest = digest(r.Spec)
}

func (r *MigMigration) HasReconciled() bool {
	return r.Status.ObservedDigest == digest(r.Spec)
}

func (r *MigMigration) InSandbox() bool {
	return r.GetNamespace() == settings.Settings.Namespace.Sandbox
}

func (r *MigMigration) InPrivileged() bool {
	return r.GetNamespace() == settings.Settings.Namespace.Privileged
}

// Cluster
func (r *MigIdentity) GetCorrelationLabels() map[string]string {
	key, value := r.GetCorrelationLabel()
	return map[string]string{
		PartOfLabel: Application,
		key:         value,
	}
}

func (r *MigIdentity) GetCorrelationLabel() (string, string) {
	return CorrelationLabel(r, r.UID)
}

func (r *MigIdentity) GetNamespace() string {
	return r.Namespace
}

func (r *MigIdentity) GetName() string {
	return r.Name
}

func (r *MigIdentity) MarkReconciled() {
	uuid, _ := uuid.NewUUID()
	if r.Annotations == nil {
		r.Annotations = map[string]string{}
	}
	r.Annotations[TouchAnnotation] = uuid.String()
	r.Status.ObservedDigest = digest(r.Spec)
}

func (r *MigIdentity) HasReconciled() bool {
	return r.Status.ObservedDigest == digest(r.Spec)
}

func (r *MigIdentity) InSandbox() bool {
	return r.GetNamespace() == settings.Settings.Namespace.Sandbox
}

func (r *MigIdentity) InPrivileged() bool {
	return r.GetNamespace() == settings.Settings.Namespace.Privileged
}

//
// Generate a sha256 hex-digest for an object.
func digest(object interface{}) string {
	j, _ := json.Marshal(object)
	hash := sha256.New()
	hash.Write(j)
	digest := hex.EncodeToString(hash.Sum(nil))
	return digest
}
