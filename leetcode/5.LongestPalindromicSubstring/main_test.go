package main

import (
	"fmt"
	"testing"
)

var table = []struct{
	text string
	palind bool
	longest int
}{
	{"bob", true, 3},
	{"bobo", false, 3},
	{"hello", false, 2},
	{"tenet", true, 5},
	{"babad", false, 3},
	{"cbbd", false, 2},
	{"a", true, 1},
	{"ac", false, 1},
	{
		"cmmrracelnclsbtdmuxtfiyahrvxuwreyorosyqapfpnsntommsujibzwhgugwtvxsdsltiiyymiofbslwbwevmjrsbbssicnxptvwmsmiifypoujftxylpyvirfueagprfyyydxeiftathaygmolkcwoaavmdmjsuwoibtuqoewaexihispsshwnsurjopdwttlzyqdbkypvjsbuidsdnpgklhwfnqdvlffcysnxeywvwvblatmxbflnuykhfhjptenhcxqinomlwalvlezefqybpuepbnymzlruuirpiatqgjgcnfmrlzshauoxuoqopcikogfwpssjdeplytcapmujyvgtfmmolnuadpwblgmcaututcrwsqrlpaaqobjfnhudmsulztbdkxpfejavastxejtpbqfftdtcdhvtpbzfuqcwkxtldtjycreimiujtxudtmokcoebhodbkgkgxjzrgyuqhozqtidltodlkziyofdeszwiobkwesdijxbbagguxvofvtphqxgvidqfkljufgubjmjllxoanbizwtedykwmneaosopynzlzvrlqcmyaahdcagfatlhwtgqxsyxwjhexfiplwtrtydjzrsysrcwszlntwrpgfedhgjzhztffqnjotlfudvczwfkhuwmdzcqgrmfttwaxocohtuscdxwfvhcymjpkqexusdaccw",
		false,
		5,
	},
}

func TestIsPalindrome(t *testing.T) {
	for _, tt := range table {
		t.Run(fmt.Sprintf("Is %s palindrome?", tt.text), func(t *testing.T) {
			if IsPalindrome(tt.text, 0, len(tt.text)) != tt.palind {
				t.Errorf("%s: wanted %v, got %v", tt.text, tt.palind, !tt.palind)
			}
		})
	}
}

func TestLongestPalindrome(t *testing.T) {
	for _, tt := range table {
		t.Run(fmt.Sprintf("Longest palindrome of %s", tt.text), func(t *testing.T) {
			memo := make(map[string]string)
			longest := LongestPalindrome(tt.text, 0, len(tt.text), memo)
			if len(longest) != tt.longest {
				t.Errorf("got %s (%d), wanted (%d)",
					longest, len(longest), tt.longest)
			}
		})
	}
}
