package sort_integers

import (
	"reflect"
	"testing"
)

func TestSortInt64(t *testing.T) {
	type args struct {
		arr []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "Sorting a array of big integers",
			args: args{
				arr: []int64{1000000007, 1000000003, 1000000006, 1000000001, 1000000004, 1000000002, 1000000005},
			},
			want: []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005, 1000000006, 1000000007},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortInt64(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
