// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package go-restful provides functions to trace the emicklei/go-restful/v3
// package (https://github.com/emicklei/go-restful).
//
// Instrumentation of an incoming request is achieved via a go-restful
// FilterFunc which may be applied at the container level, at the
// webservice level or at a route level
package restful // import "go.opentelemetry.io/contrib/instrumentation/emicklei/go-restful"
