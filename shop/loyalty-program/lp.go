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

func (lp *LoyaltyProgram) getIDTopEntity() ID {
	tmp := make(map[ID]struct{})

	for _, item := range lp.Tree {
		for k, v := range item {
			tmp[k] = struct{}{}
			tmp[v] = struct{}{}
		}
	}

	for key := range lp.Tree {
		if _, ok := tmp[key]; !ok {
			return key
		}
	}

	return ""
}

func (lp *LoyaltyProgram) getEntity(id ID) (Entity, error) {
	for _, e := range lp.Entities {
		if e.ID == id {
			return e, nil
		}
	}

	return Entity{}, ErrNodeNotExist
}

func (lp *LoyaltyProgram) nextNode(nodeID ID, nodeType NodeType) (Entity, error) {
	for _, v := range lp.Tree[nodeID] {
		entity, err := lp.getEntity(v)
		if err != nil {
			return Entity{}, fmt.Errorf("failed getting node, %w", err)
		}

		if entity.Type == nodeType {
			return entity, nil
		}
	}

	return Entity{}, ErrNodeNotExist
}

// nolint: funlen,forcetypeassert,gomnd
func (lp *LoyaltyProgram) DTree() (*dtree.Tree, error) {
	tree := []dtree.Tree{}
	curID := lp.getIDTopEntity()

	entity, err := lp.getEntity(curID)
	if err != nil {
		return nil, fmt.Errorf("failed getting node, %w", err)
	}

	tree = append(
		tree,
		dtree.Tree{
			ID:    1,
			Order: 1,
		},
		dtree.Tree{
			ID:       2,
			Order:    2,
			ParentID: 1,
			Key:      entity.Form.Fields["expression1"].(string),
			Operator: entity.Form.Fields["operator"].(string),
			Value:    entity.Form.Fields["expression2"],
		},
	)

	num := 2
	nType := Operation

	for node, err := lp.nextNode(curID, Operation); err == nil; node, err = lp.nextNode(curID, nType) {
		num++

		if nType == Operation {
			tree = append(tree, dtree.Tree{
				ID:       num,
				Order:    num,
				ParentID: num - 1,
				Key:      "result",
				Operator: "expression",
				Value:    node.Form.Fields["operation"],
				Name:     node.Form.Fields["name"].(string),
			})

			num++

			tree = append(tree, dtree.Tree{
				ID:       num,
				Order:    num,
				ParentID: num - 1,
				Key:      "result",
				Operator: "request2value",
			})
		}

		if nType == Condition {
			curID = node.ID

			tree = append(tree, dtree.Tree{
				ID:       num,
				Order:    num,
				ParentID: 1,
				Key:      node.Form.Fields["expression1"].(string),
				Operator: node.Form.Fields["operator"].(string),
				Value:    node.Form.Fields["expression2"],
			})
		}

		if nType == Operation {
			nType = Condition
		} else {
			nType = Operation
		}
	}

	return dtree.CreateTree(tree), nil
}
