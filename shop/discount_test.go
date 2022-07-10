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
			fields: fields{Discounter: New(DefaultRulesLoyaltyProgram())},
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
			fields: fields{Discounter: New(DefaultRulesLoyaltyProgram())},
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
			fields: fields{Discounter: New(DefaultRulesLoyaltyProgram())},
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
			fields: fields{Discounter: New(DefaultRulesLoyaltyProgram())},
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
