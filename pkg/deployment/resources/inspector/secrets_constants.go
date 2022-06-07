//
// DISCLAIMER
//
// Copyright 2016-2022 ArangoDB GmbH, Cologne, Germany
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

package inspector

import (
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Secret
const (
	SecretGroup     = core.GroupName
	SecretResource  = "secrets"
	SecretKind      = "Secret"
	SecretVersionV1 = "v1"
)

func SecretGK() schema.GroupKind {
	return schema.GroupKind{
		Group: SecretGroup,
		Kind:  SecretKind,
	}
}

func SecretGKv1() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   SecretGroup,
		Kind:    SecretKind,
		Version: SecretVersionV1,
	}
}

func SecretGR() schema.GroupResource {
	return schema.GroupResource{
		Group:    SecretGroup,
		Resource: SecretResource,
	}
}

func SecretGRv1() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    SecretGroup,
		Resource: SecretResource,
		Version:  SecretVersionV1,
	}
}

func (p *secretsInspectorV1) GroupVersionKind() schema.GroupVersionKind {
	return SecretGKv1()
}

func (p *secretsInspectorV1) GroupVersionResource() schema.GroupVersionResource {
	return SecretGRv1()
}

func (p *secretsInspector) GroupKind() schema.GroupKind {
	return SecretGK()
}

func (p *secretsInspector) GroupResource() schema.GroupResource {
	return SecretGR()
}