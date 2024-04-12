// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/module_include_list/module_include_list.go - DO NOT EDIT.

package include

import (
	// Import packages to perform 'func InitializeModule()' when in-use.
	m0 "github.com/elastic/beats/v7/metricbeat/autodiscover/builder/hints"
	m1 "github.com/elastic/beats/v7/metricbeat/autodiscover/appender/kubernetes/token"
	m2 "github.com/elastic/beats/v7/metricbeat/processor/add_kubernetes_metadata"

	// Import packages that perform 'func init()'.
)

// InitializeModules initialize all of the modules.
func InitializeModule() {
	m0.InitializeModule()
	m1.InitializeModule()
	m2.InitializeModule()
}
