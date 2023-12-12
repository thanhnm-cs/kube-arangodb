//
// DISCLAIMER
//
// Copyright 2023 ArangoDB GmbH, Cologne, Germany
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

import shared "github.com/arangodb/kube-arangodb/pkg/apis/shared"

type PodTemplate struct {
	// Scheduling keeps the scheduling information
	*Scheduling `json:",inline"`

	// ContainerNamespace keeps the Container layer Kernel namespace configuration
	*ContainerNamespace `json:",inline"`

	// SecurityPod keeps the security settings for Pod
	*SecurityPod `json:",inline"`
}

func (a *PodTemplate) GetSecurityPod() *SecurityPod {
	if a == nil {
		return nil
	}

	return a.SecurityPod
}

func (a *PodTemplate) GetScheduling() *Scheduling {
	if a == nil {
		return nil
	}

	return a.Scheduling
}

func (a *PodTemplate) GetContainerNamespace() *ContainerNamespace {
	if a == nil {
		return nil
	}

	return a.ContainerNamespace
}

func (a *PodTemplate) Validate() error {
	if a == nil {
		return nil
	}
	return shared.WithErrors(
		a.GetScheduling().Validate(),
		a.GetContainerNamespace().Validate(),
		a.GetSecurityPod().Validate(),
	)
}
