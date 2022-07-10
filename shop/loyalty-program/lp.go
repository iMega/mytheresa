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

import (
	"errors"
	"fmt"

	"github.com/tkanos/go-dtree"
)

var (
	ErrTreeNodeNotExist = errors.New("tree node not exists")
	ErrNodeNotExist     = errors.New("node not exists")
)

type LoyaltyProgram struct {
	ID       string
	Entities []Entity
	Tree     map[ID]map[ID]ID
}

func (lp *LoyaltyProgram) GetIdTopEntity() ID {
	h := make(map[ID]struct{})
	for _, item := range lp.Tree {
		for k, v := range item {
			h[k] = struct{}{}
			h[v] = struct{}{}
		}
	}

	for key := range lp.Tree {
		if _, ok := h[key]; !ok {
			return ID(key)
		}
	}

	return ""
}

func (lp *LoyaltyProgram) GetEntity(id ID) (Entity, error) {
	for _, e := range lp.Entities {
		if e.ID == id {
			return e, nil
		}
	}

	return Entity{}, ErrNodeNotExist
}

func (lp *LoyaltyProgram) nextNode(nodeID ID, t NodeType) (Entity, error) {
	for _, v := range lp.Tree[nodeID] {
		e, err := lp.GetEntity(v)
		if err != nil {
			return Entity{}, fmt.Errorf("failed getting node, %s", err)
		}

		if e.Type == t {
			return e, nil
		}
	}

	return Entity{}, ErrNodeNotExist
}

func (lp *LoyaltyProgram) DTree() (*dtree.Tree, error) {
	t := []dtree.Tree{}
	curID := lp.GetIdTopEntity()

	e, err := lp.GetEntity(curID)
	if err != nil {
		return nil, fmt.Errorf("failed getting node, %s", err)
	}

	t = append(
		t,
		dtree.Tree{
			ID:    1,
			Order: 1,
		},
		dtree.Tree{
			ID:       2,
			Order:    2,
			ParentID: 1,
			Key:      e.Form.Fields["expression1"].(string),
			Operator: e.Form.Fields["operator"].(string),
			Value:    e.Form.Fields["expression2"],
		},
	)

	i := 2
	nType := Operation
	for n, err := lp.nextNode(curID, Operation); err == nil; n, err = lp.nextNode(curID, nType) {
		i++
		if nType == Operation {
			t = append(t, dtree.Tree{
				ID:       i,
				Order:    i,
				ParentID: i - 1,
				Key:      "result",
				Operator: "expression",
				Value:    n.Form.Fields["operation"],
				Name:     n.Form.Fields["name"].(string),
			})

			i++
			t = append(t, dtree.Tree{
				ID:       i,
				Order:    i,
				ParentID: i - 1,
				Key:      "result",
				Operator: "request2value",
			})
		}

		if nType == Condition {
			curID = n.ID

			t = append(t, dtree.Tree{
				ID:       i,
				Order:    i,
				ParentID: 1,
				Key:      n.Form.Fields["expression1"].(string),
				Operator: n.Form.Fields["operator"].(string),
				Value:    n.Form.Fields["expression2"],
			})
		}

		if nType == Operation {
			nType = Condition
		} else {
			nType = Operation
		}
	}

	return dtree.CreateTree(t), nil
}
