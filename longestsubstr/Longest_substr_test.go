package longestsubstr

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Max sub length 3:abc", args{"abcabaca"}, 3},
		{"Max sub length 6:abc123", args{"abcabc123"}, 6},
		{"Max sub length 7:abc1234", args{"abcabc1234a"}, 7},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.str); got != tt.want {
				t.Errorf(" lengthOfLongestSubstring() test case %v failed. Get %v, want %v", i, got, tt.want)
			}
		})
	}
}
