// Copyright The Cryostat Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	operatorv1beta2 "github.com/cryostatio/cryostat-operator/api/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CryostatInstance is an abstraction to work with Cryostat CRDs.
type CryostatInstance struct {
	// Name of the Cryostat CR.
	Name string
	// Namespace to install Cryostat into. For Cryostat, this comes from the
	// CR's namespace.
	InstallNamespace string
	// Namespaces that Cryostat should look for targets. For Cryostat, this
	// comes from spec.TargetNamespaces.
	TargetNamespaces []string
	// Namespaces that the operator has successfully set up RBAC for Cryostat to monitor targets
	// in that namespace. For Cryostat, this is a reference to status.TargetNamespaces.
	TargetNamespaceStatus *[]string
	// Reference to the Cryostat Spec properties.
	Spec *operatorv1beta2.CryostatSpec
	// Reference to the Cryostat Status properties.
	Status *operatorv1beta2.CryostatStatus
	// The actual CR instance as a generic Kubernetes object.
	Object client.Object
}

// FromCryostat creates a CryostatInstance from a Cryostat CR
func FromCryostat(cr *operatorv1beta2.Cryostat) *CryostatInstance {
	return &CryostatInstance{
		Name:                  cr.Name,
		InstallNamespace:      cr.Namespace,
		TargetNamespaces:      cr.Spec.TargetNamespaces,
		TargetNamespaceStatus: &cr.Status.TargetNamespaces,

		Spec:   &cr.Spec,
		Status: &cr.Status,

		Object: cr,
	}
}
