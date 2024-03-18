//
// DISCLAIMER
//
// Copyright 2023-2024 ArangoDB GmbH, Cologne, Germany
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

package crds

import (
	_ "embed"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	"github.com/arangodb/go-driver"

	"github.com/arangodb/kube-arangodb/pkg/util"
)

const (
	SchedulerProfileVersion = driver.Version("1.0.0")
)

func SchedulerProfileWithOptions(opts ...func(*CRDOptions)) *apiextensions.CustomResourceDefinition {
	return getCRD(schedulerProfileCRD, schedulerProfileCRDSchemas, opts...)
}

func SchedulerProfileDefinitionWithOptions(opts ...func(*CRDOptions)) Definition {
	return Definition{
		Version: SchedulerProfileVersion,
		CRD:     SchedulerProfileWithOptions(opts...),
	}
}

var schedulerProfileCRD = util.NewYamlLoader[apiextensions.CustomResourceDefinition](schedulerProfile)
var schedulerProfileCRDSchemas = util.NewYamlLoader[crdSchemas](schedulerProfileSchemaRaw)

//go:embed scheduler-profile.yaml
var schedulerProfile []byte

//go:embed scheduler-profile.schema.generated.yaml
var schedulerProfileSchemaRaw []byte
