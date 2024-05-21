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
	arangodbOperatorEngineAssertions = metrics.NewDescription("arangodb_operator_engine_assertions", "Number of assertions invoked during Operator runtime", []string{`key`}, nil)
)

func init() {
	registerDescription(arangodbOperatorEngineAssertions)
}

func ArangodbOperatorEngineAssertions() metrics.Description {
	return arangodbOperatorEngineAssertions
}

func ArangodbOperatorEngineAssertionsCounter(value float64, key string) metrics.Metric {
	return ArangodbOperatorEngineAssertions().Counter(value, key)
}
