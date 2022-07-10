// Copyright © 2020 Dmitry Stoletov <info@imega.ru>
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

package loyaltyprogram

type ID string

type NodeType string

type Entity struct {
	ID            ID
	Type          NodeType
	DiagramEntity DiagramEntity
	Form          Form
}

type DiagramEntity map[string]interface{}

type ShortForm struct {
	Entities []SFEntity `json:"entities"`
}

type SFEntity struct {
	ID   ID   `json:"id"`
	Form Form `json:"form"`
}

type Form struct {
	Type   NodeType               `json:"type"`
	Order  int                    `json:"order"`
	Fields map[string]interface{} `json:"fields"`
}

const (
	Condition NodeType = "condition"
	Operation NodeType = "operation"
)
