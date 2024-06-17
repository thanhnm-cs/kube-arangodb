//
// DISCLAIMER
//
// Copyright 2016-2024 ArangoDB GmbH, Cologne, Germany
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
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

package v1

import (
	"k8s.io/apimachinery/pkg/types"

	"github.com/arangodb/kube-arangodb/pkg/util"
)

type ArangoMemberSpec struct {
	// Group define Member Groups.
	// +doc/type: string
	Group ServerGroup `json:"group,omitempty"`

	ID string `json:"id,omitempty"`

	// DeploymentUID define Deployment UID.
	DeploymentUID types.UID `json:"deploymentUID,omitempty"`

	// Overrides define Member Overrides (Override values from ServerGroup).
	Overrides *ArangoMemberSpecOverrides `json:"overrides,omitempty"`

	// Template keeps template which is gonna be applied on the Pod.
	Template *ArangoMemberPodTemplate `json:"template,omitempty"`

	// DeletionPriority define Deletion Priority.
	// Higher value means higher priority. Default is 0.
	// Example: set 1 for Coordinator which should be deleted first and scale down coordinators by one.
	DeletionPriority *int `json:"deletion_priority,omitempty"`
}

func (a *ArangoMemberSpec) GetDeletionPriority() int {
	if a == nil {
		return 0
	}

	return util.TypeOrDefault[int](a.DeletionPriority, 0)
}
