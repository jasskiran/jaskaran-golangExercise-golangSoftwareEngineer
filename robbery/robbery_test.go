package robbery

import "testing"

func TestRobbery(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "RobberyShouldReturnValueOfCentreHouse",
			args: args{nums: []int{2,3,4}},
			want: 3,
		},
		{
			name: "RobberyShouldReturnMaximumValueFromAlternateHouses",
			args: args{nums: []int{1,2,3,4}},
			want: 6,
		},
		{
			name: "RobberyShouldReturn0",
			args: args{nums: []int{1,2}},
			want: 0,
		},
		{
			name: "RobberyShouldReturn0",
			args: args{nums: []int{1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Robbery(tt.args.nums); got != tt.want {
				t.Errorf("Robbery() = %v, want %v", got, tt.want)
			}
		})
	}
}