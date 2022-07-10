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

package shop

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Knetic/govaluate"
	"github.com/imega/mytheresa/domain"
	loyaltyprogram "github.com/imega/mytheresa/shop/loyalty-program"
	"github.com/tkanos/go-dtree"
)

type Discount struct {
	WithDiscount30 bool
	WithDiscount15 bool
	LP             loyaltyprogram.LoyaltyProgram
}

func New(rules []byte) *Discount {
	lp := loyaltyprogram.LoyaltyProgram{}
	if err := json.Unmarshal(rules, &lp); err != nil {
		return nil
	}

	return &Discount{LP: lp}
}

func (d *Discount) Calc(product domain.Product) domain.Discount {
	vars := map[string]interface{}{
		"category_name":         product.Category,
		"catalog_product_sku":   product.SKU,
		"catalog_product_price": product.Price.Units,
	}

	noDiscount := domain.Discount{
		Price: domain.Money{
			Units:    product.Price.Units,
			Currency: product.Price.Currency,
		},
	}

	node, err := calc(d.LP, vars)
	if err != nil {
		return noDiscount
	}

	v, ok := node.Value.(uint64)
	if !ok {
		return noDiscount
	}

	if v == 0 {
		return noDiscount
	}

	return domain.Discount{
		Price: domain.Money{
			Units:    v,
			Currency: product.Price.Currency,
		},
		Value: node.Name,
	}
}

func calc(
	lp loyaltyprogram.LoyaltyProgram,
	vars map[string]interface{},
) (*dtree.Tree, error) {
	tree, err := lp.DTree()
	if err != nil {
		return nil, fmt.Errorf("failed to create tree, %s", err)
	}

	varsJSON, err := json.Marshal(vars)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal vars, %s", err)
	}

	node, err := tree.ResolveJSON(varsJSON, func(o *dtree.TreeOptions) {
		o.Operators = make(map[string]dtree.Operator)
		o.Operators["expression"] = expression
		o.Operators["request2value"] = request2value
	})
	if err != nil {
		return nil, fmt.Errorf("failed to resolve tree, %s", err)
	}

	return node, nil
}

func expression(
	requests map[string]interface{},
	node *dtree.Tree,
) (*dtree.Tree, error) {
	expression, err := govaluate.NewEvaluableExpression(node.Value.(string))
	if err != nil {
		return nil, fmt.Errorf("failed to initial expression, %s", err)
	}

	result, err := expression.Evaluate(requests)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate expression, %s", err)
	}

	s := fmt.Sprintf("%.0f", result)
	val, _ := strconv.ParseUint(s, 10, 64)

	requests[node.Key] = val
	requests["name"] = node.Name

	return node, nil
}

func request2value(
	requests map[string]interface{},
	node *dtree.Tree,
) (*dtree.Tree, error) {
	node.Value = requests[node.Key]
	node.Name = requests["name"].(string)

	return node, nil
}
