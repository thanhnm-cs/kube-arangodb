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

package metric_descriptions

import (
	"github.com/arangodb/kube-arangodb/pkg/util/metrics"
)

var (
	arangodbOperatorAgencyIndex = metrics.NewDescription("arangodb_operator_agency_index", "Current index of the agency cache", []string{`namespace`, `name`}, nil)
)

func init() {
	registerDescription(arangodbOperatorAgencyIndex)
}

func NewArangodbOperatorAgencyIndexGaugeFactory() metrics.FactoryGauge[ArangodbOperatorAgencyIndexInput] {
	return metrics.NewFactoryGauge[ArangodbOperatorAgencyIndexInput]()
}

func NewArangodbOperatorAgencyIndexInput(namespace string, name string) ArangodbOperatorAgencyIndexInput {
	return ArangodbOperatorAgencyIndexInput{
		Namespace: namespace,
		Name:      name,
	}
}

type ArangodbOperatorAgencyIndexInput struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

func (i ArangodbOperatorAgencyIndexInput) Gauge(value float64) metrics.Metric {
	return ArangodbOperatorAgencyIndexGauge(value, i.Namespace, i.Name)
}

func (i ArangodbOperatorAgencyIndexInput) Desc() metrics.Description {
	return ArangodbOperatorAgencyIndex()
}

func ArangodbOperatorAgencyIndex() metrics.Description {
	return arangodbOperatorAgencyIndex
}

func ArangodbOperatorAgencyIndexGauge(value float64, namespace string, name string) metrics.Metric {
	return ArangodbOperatorAgencyIndex().Gauge(value, namespace, name)
}
