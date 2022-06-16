// Copyright 2022 VoyTechnology
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package descriptor

type Method struct {
	Name         string
	InternalName string
	RequestType  string
	ResponseType string
	Path         string
	Method       string
	Body         string
}

type Service struct {
	Name         string
	InternalName string

	Methods map[string]map[string]*Method
}
