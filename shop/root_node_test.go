package shop

import (
	"context"
	"reflect"
	"testing"

	"github.com/imega/mytheresa/domain"
	"github.com/imega/mytheresa/storage"
)

func TestRootNode_GetSKUs(t *testing.T) {
	type fields struct {
		Storage func() domain.Storage
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "storage is empty",
			fields: fields{
				Storage: func() domain.Storage {
					return storage.New()
				},
			},
			args:    args{ctx: context.Background()},
			want:    []string{},
			wantErr: false,
		},
		{
			name: "storage contains one sku",
			fields: fields{
				Storage: func() domain.Storage {
					s := storage.New()
					s.Set(
						context.Background(),
						domain.Key(domain.RootNodeKey),
						[]byte(`["00001"]`),
					)

					return s
				},
			},
			args: args{ctx: context.Background()},
			want: []string{"00001"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &RootNode{
				Storage: tt.fields.Storage(),
			}
			got, err := node.GetSKUs(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("RootNode.GetSKUs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RootNode.GetSKUs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRootNode_AddSKU(t *testing.T) {
	type fields struct {
		Storage func() domain.Storage
	}
	type args struct {
		ctx context.Context
		sku string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "storage is empty",
			fields: fields{
				Storage: func() domain.Storage {
					return storage.New()
				},
			},
			args: args{
				ctx: context.Background(),
				sku: "00001",
			},
			want: []string{"00001"},
		},
		{
			name: "storage contains one sku",
			fields: fields{
				Storage: func() domain.Storage {
					s := storage.New()
					s.Set(
						context.Background(),
						domain.Key(domain.RootNodeKey),
						[]byte(`["00001"]`),
					)

					return s
				},
			},
			args: args{
				ctx: context.Background(),
				sku: "00002",
			},
			want: []string{"00001", "00002"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &RootNode{
				Storage: tt.fields.Storage(),
			}
			if err := node.AddSKU(tt.args.ctx, tt.args.sku); (err != nil) != tt.wantErr {
				t.Errorf("RootNode.AddSKU() error = %v, wantErr %v", err, tt.wantErr)
			}
			got, err := node.GetSKUs(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("RootNode.GetSKUs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RootNode.GetSKUs() = %v, want %v", got, tt.want)
			}
		})
	}
}
