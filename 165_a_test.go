package main

import (
	"reflect"
	"testing"
)

func TestSumOfNumbers(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			name:    "1",
			args:    args{slice: []int{1, 2, 3, 4}},
			wantSum: 10,
		},
		{
			name:    "2",
			args:    args{slice: []int{-1, 2, -3, 4}},
			wantSum: 2,
		},
		{
			name:    "3",
			args:    args{slice: []int{1, -2, 3, -4}},
			wantSum: -2,
		},
		{
			name:    "4",
			args:    args{slice: []int{}},
			wantSum: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := SumOfNumbers(tt.args.slice); gotSum != tt.wantSum {
				t.Errorf("SumOfNumbers() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestIsValidDataInSlice(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{slice: []string{"1", "2", "3", "4"}},
			want: true,
		},
		{
			name: "2",
			args: args{slice: []string{"-1", "2", "-3", "4"}},
			want: true,
		},
		{
			name: "3",
			args: args{slice: []string{"1", "2", "3", "4r"}},
			want: false,
		},
		{
			name: "4",
			args: args{slice: []string{"1", "2", "3", "ad"}},
			want: false,
		},
		{
			name: "5",
			args: args{slice: []string{}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidDataInSlice(tt.args.slice); got != tt.want {
				t.Errorf("IsValidDataInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertSlice(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{slice: []string{"1", "2", "3", "4"}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "2",
			args: args{slice: []string{"4", "3", "2", "1"}},
			want: []int{4, 3, 2, 1},
		},
		{
			name: "3",
			args: args{slice: []string{"-4", "3", "-2", "1"}},
			want: []int{-4, 3, -2, 1},
		},
		{
			name: "4",
			args: args{slice: []string{}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertSlice(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
