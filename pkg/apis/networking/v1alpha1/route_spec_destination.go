//
// DISCLAIMER
//
// Copyright 2024 ArangoDB GmbH, Cologne, Germany
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

package v1alpha1

import shared "github.com/arangodb/kube-arangodb/pkg/apis/shared"

type ArangoRouteSpecDestination struct {
	// Service defines service upstream reference
	Service *ArangoRouteSpecDestinationService `json:"service,omitempty"`

	// Schema defines HTTP/S schema used for connection
	Schema *ArangoRouteSpecDestinationSchema `json:"schema,omitempty"`

	// TLS defines TLS Configuration
	TLS *ArangoRouteSpecDestinationTLS `json:"tls,omitempty"`
}

func (s *ArangoRouteSpecDestination) Validate() error {
	if s == nil {
		s = &ArangoRouteSpecDestination{}
	}

	if err := shared.WithErrors(
		shared.ValidateOptionalInterfacePath("service", s.Service),
		shared.ValidateOptionalInterfacePath("schema", s.Schema),
		shared.ValidateOptionalInterfacePath("tls", s.TLS),
	); err != nil {
		return err
	}

	return nil
}
