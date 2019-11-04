package go_shadowsocks

import (
	"reflect"
	"testing"
)

func TestRandPassword(t *testing.T) {
	tests := []struct {
		name string
		want *Password
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandPassword(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RandPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}