package efmt

import (
	"testing"
)

// go test -v.   Runs all tests beginning modules beginning with "Test"
func TestNtag(t *testing.T) {
	// if testing.Short() {
	// 	t.Skip("skipping test in short mode.")
	// }

	var nt *Ntag
	tests := []struct {
		name  string
		count int
		want  string
	}{
		{"NewNtag_Test", 1, "1"},
		{"Push____Test", 2, "1.1"},
		{"Next____Test", 2, "1.1.2"},
		{"Indent_1Test", 2, "    "},
		{"Indent_2Test", 2, "______"},
		{"Pop_____Test", 2, "1.2"},
		{"Clone___Test", 2, "1.2.2"},
	}
	for testNo, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch {
			case testNo == 0:
				nt = NewNtag()
				if got := nt.String(); got != tt.want {
					t.Errorf("Ntag.String() = \"%v\", want \"%v\"", got, tt.want)
				}
			case testNo == 1:
				nt.Push()
				if got := nt.String(); got != tt.want {
					t.Errorf("Ntag.String() = \"%v\", want \"%v\"", got, tt.want)
				}
			case testNo == 2:
				nt.Push()
				nt.Next()
				if got := nt.String(); got != tt.want {
					t.Errorf("Ntag.String() = \"%v\", want \"%v\"", got, tt.want)
				}
			case testNo == 3:
				if got := nt.Indent(); got != tt.want {
					t.Errorf("Ntag.String() = \"%v\", want \"%v\"", got, tt.want)
				}
			case testNo == 4:
				nt.SetIndent("___")
				if got := nt.Indent(); got != tt.want {
					t.Errorf("Ntag.String() = \"%v\", want \"%v\"", got, tt.want)
				}
			case testNo == 5:
				nt.Pop()
				nt.Next()
				if got := nt.String(); got != tt.want {
					t.Errorf("Ntag.String() = \"%v\", want \"%v\"", got, tt.want)
				}
			case testNo == 6:
				ct := nt.Clone()
				nt.Next()
				ct.Push()
				ct.Next()
				if got := ct.String(); got != tt.want {
					t.Errorf("ct.String() = \"%v\", want \"%v\"", got, tt.want)
				}
			}
		})
	}
}
