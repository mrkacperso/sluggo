package sluggo

import (
	"reflect"
	"testing"
)

func TestGetSlug(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "unchanged", args: args{text:"foo"}, want: "foo"},
		{name: "to mixed case", args: args{text:"FoO"}, want: "foo"},
		{name: "with space", args: args{text:"foo bar"}, want: "foo_bar"},
		{name: "multiple spaces", args: args{text:"f  oo"}, want: "f_oo"},
		{name: "Special characters", args: args{text:"f*$oo&śŚb@r"}, want: "f-oo-b-r"},
		{name: "With dash", args: args{text:"foo-bar"}, want: "foo-bar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSlug(tt.args.text); got != tt.want {
				t.Errorf("GetSlug() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSlugWithOpts(t *testing.T) {
	type args struct {
		text string
		opts *SlugOptions
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSlugWithOpts(tt.args.text, tt.args.opts); got != tt.want {
				t.Errorf("GetSlugWithOpts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDefaultOptions(t *testing.T) {
	tests := []struct {
		name string
		want *SlugOptions
	}{
		{name: "test defaults", want: &SlugOptions{
			TrimSpaces:     true,
			MaxLength:      0,
			Case:           Lowercase,
			SpaceSymbol:    "_",
			CharSymbol:     "-",
			PreserveLength: false,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultOptions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}