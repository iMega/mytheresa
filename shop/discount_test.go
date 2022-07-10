package shop

import (
	"testing"

	"github.com/imega/mytheresa/domain"
	"github.com/stretchr/testify/assert"
)

func TestDiscount_Calc(t *testing.T) {
	type fields struct {
		Discounter *Discount
	}
	type args struct {
		product domain.Product
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   domain.Discount
	}{
		{
			name:   "given Category boots with discount of 30% will apply 30%",
			fields: fields{Discounter: New(getRawLP())},
			args: args{product: domain.Product{
				Category: "boots",
				Price:    domain.Money{Units: 1000},
			}},
			want: domain.Discount{
				Price: domain.Money{Units: 700},
				Value: "30%",
			},
		},
		{
			name:   "given SKU 000003 with discount of 15% will apply 15%",
			fields: fields{Discounter: New(getRawLP())},
			args: args{product: domain.Product{
				Category: "sandals",
				SKU:      "000003",
				Price:    domain.Money{Units: 1000},
			}},
			want: domain.Discount{
				Price: domain.Money{Units: 850},
				Value: "15%",
			},
		},
		{
			name:   "given Category boots and SKU 000003 with highest discount of 30% and lowest 15% will apply 30%",
			fields: fields{Discounter: New(getRawLP())},
			args: args{product: domain.Product{
				Category: "boots",
				SKU:      "000003",
				Price:    domain.Money{Units: 1000},
			}},
			want: domain.Discount{
				Price: domain.Money{Units: 700},
				Value: "30%",
			},
		},
		{
			name:   "product sandals and SKU 000001 without applicable discount",
			fields: fields{Discounter: New(getRawLP())},
			args: args{product: domain.Product{
				Category: "sandals",
				SKU:      "000001",
				Price:    domain.Money{Units: 1000},
			}},
			want: domain.Discount{
				Price: domain.Money{Units: 1000},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.Discounter.Calc(tt.args.product)
			assert.Equal(t, tt.want, got)
		})
	}
}

func getRawLP() []byte {
	return []byte(`{
        "ID": "4a191abe-9263-499c-8b72-98e81c9b32dd",
        "Tree": {
            "518e26dc-dfdb-4f99-8638-bbb907a0d24f": {
                "518e26dc-dfdb-4f99-8638-bbb907a0d24f_15f729bf-70a1-4789-842c-4651fbd6d055": "15f729bf-70a1-4789-842c-4651fbd6d055",
                "518e26dc-dfdb-4f99-8638-bbb907a0d24f_a7e72a68-b77e-4abe-819a-7ed1750a0e25": "a7e72a68-b77e-4abe-819a-7ed1750a0e25"
            },
            "a7e72a68-b77e-4abe-819a-7ed1750a0e25": {
                "a7e72a68-b77e-4abe-819a-7ed1750a0e25_ab47a03d-7a93-421e-8c80-256ddc20cade": "ab47a03d-7a93-421e-8c80-256ddc20cade"
            }
        },
        "Entities": [
            {
                "ID": "518e26dc-dfdb-4f99-8638-bbb907a0d24f",
                "Form": {
                    "type": "condition",
                    "order": 0,
                    "fields": {
                        "operator": "==",
                        "expression1": "category_name",
                        "expression2": "boots",
                        "operation_id": "15f729bf-70a1-4789-842c-4651fbd6d055",
                        "view_operation_key": "discount_percent",
                        "view_operation_value": 30,
                        "name": ""
                    }
                },
                "Type": "condition",
                "DiagramEntity": null
            },
            {
                "ID": "15f729bf-70a1-4789-842c-4651fbd6d055",
                "Form": {
                    "type": "operation",
                    "order": 0,
                    "fields": {
                        "operation": "catalog_product_price*(1 - 0.01*30)",
                        "name": "30%"
                    }
                },
                "Type": "operation",
                "DiagramEntity": null
            },
            {
                "ID": "a7e72a68-b77e-4abe-819a-7ed1750a0e25",
                "Form": {
                    "type": "condition",
                    "order": 1,
                    "fields": {
                        "operator": "==",
                        "expression1": "catalog_product_sku",
                        "expression2": "000003",
                        "operation_id": "ab47a03d-7a93-421e-8c80-256ddc20cade",
                        "view_operation_key": "discount_percent",
                        "view_operation_value": 15,
                        "name":""
                    }
                },
                "Type": "condition",
                "DiagramEntity": null
            },
            {
                "ID": "ab47a03d-7a93-421e-8c80-256ddc20cade",
                "Form": {
                    "type": "operation",
                    "order": 0,
                    "fields": {
                        "operation": "catalog_product_price*(1 - 0.01*15)",
                        "name": "15%"
                    }
                },
                "Type": "operation",
                "DiagramEntity": null
            }
        ]
    }
    `)
}
