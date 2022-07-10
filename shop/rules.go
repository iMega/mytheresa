package shop

func DefaultRulesLoyaltyProgram() []byte {
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
