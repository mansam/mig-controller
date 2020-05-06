/*
Copyright 2019 Red Hat Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"errors"
	"github.com/konveyor/mig-controller/pkg/auth"
	kapi "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

// MigIdentitySpec defines the desired state of MigIdentity
type MigIdentitySpec struct {
	MigClusterRef     *kapi.ObjectReference `json:"migClusterRef"`
	IdentitySecretRef *kapi.ObjectReference `json:"identitySecretRef"`
}

// MigIdentityStatus defines the observed state of MigIdentity
type MigIdentityStatus struct {
	Conditions
	Incompatible   `json:",inline"`
	ObservedDigest string `json:"observedDigest,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MigPlan is the Schema for the migplans API
// +k8s:openapi-gen=true
type MigIdentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MigIdentitySpec   `json:"spec,omitempty"`
	Status MigIdentityStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MigPlanList contains a list of MigPlan
type MigIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MigPlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MigIdentity{}, &MigIdentityList{})
}

func (r *MigIdentity) GetSecret(client k8sclient.Client) (*kapi.Secret, error) {
	return GetSecret(client, r.Spec.IdentitySecretRef)
}

func (r *MigIdentity) GetToken(client k8sclient.Client) (string, error) {
	secret, err := r.GetSecret(client)
	if err != nil {
		return "", err
	}
	if secret == nil {
		return "", errors.New("identity secret not found")
	}
	if secret.Data == nil {
		return "", errors.New("identity secret misconfigured")
	}
	if secret.Data["token"] == nil {
		return "", errors.New("identity secret misconfigured")
	}
	return string(secret.Data["token"]), nil
}

func (r *MigIdentity) GetCluster(client k8sclient.Client) (*MigCluster, error) {
	return GetCluster(client, r.Spec.MigClusterRef)
}

func (r *MigIdentity) BuildIdentity(client k8sclient.Client) (*auth.Identity, error) {
	token, err := r.GetToken(client)
	if err != nil {
		return nil, err
	}

	cluster, err := r.GetCluster(client)
	if err != nil {
		return nil, err
	}
	if cluster == nil || !cluster.Status.IsReady() {
		return nil, errors.New("cluster is not in a ready state")
	}
	restCfg, err := cluster.BuildRestConfig(client)
	if err != nil {
		return nil, err
	}
	clusterIdentity := &auth.Identity{
		Token: token,
		RestCfg: *restCfg,
	}
	err = clusterIdentity.BuildClient()
	if err != nil {
		return nil, err
	}
	return clusterIdentity, nil
}