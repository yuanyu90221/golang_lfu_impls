package lfu

import (
	"reflect"
	"testing"
)

func TestRunLFU(t *testing.T) {
	type args struct {
		actions []string
		value   [][]int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example",
			args: args{actions: []string{"LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"},
				value: [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {3}, {4, 4}, {1}, {3}, {4}},
			},
			want: []string{"null", "null", "null", "1", "null", "-1", "3", "null", "-1", "3", "4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunLFU(tt.args.actions, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunLFU() = %v, want %v", got, tt.want)
			}
		})
	}
}
