package effectdate

import "testing"

var tests = []struct {
	Name        string
	Input       string
	Delay       int
	OpenDays    bool
	Result      string
	ShouldPanic bool
}{
	{"Simple test", "2022-12-05", 2, false, "2022-12-07", false},
	{"Other test", "2022-12-05", 4, false, "2022-12-09", false},
	{"Invalid iso date", "20220-12-05", 2, false, "", true},
	{"Other invalid iso date", "05/12/2022", 2, false, "", true},
	{"Get effect without openDays doesn't care about week ends", "2022-12-05", 7, false, "2022-12-12", false},
	{"Get effect without openDays doesn't care about week ends", "2022-12-05", 12, false, "2022-12-17", false},
	{"Get effect with openDays DOES care about week ends", "2022-12-05", 7, true, "2022-12-14", false},
	{"Get effect with openDays DOES care about week ends", "2022-12-05", 14, true, "2022-12-23", false},
	{"Get effect with openDays DOES care about holidays", "2022-07-11", 3, true, "2022-07-15", false},
	{"Get effect with openDays DOES care about holidays AND week ends", "2022-07-11", 15, true, "2022-08-02", false},
	{"holidays on week ends counts once", "2022-12-23", 2, true, "2022-12-27", false},
	{"Get effect works with past dates", "2022-12-07", -2, true, "2022-12-05", false},
	{"Getting past dates can take care about week ends and holidays", "2022-08-02", -15, true, "2022-07-11", false},
}

func TestEffectDate(t *testing.T) {
	for _, test := range tests {
		result, err := GetEffectDate(test.Input, test.Delay, test.OpenDays)
		if err == nil && test.ShouldPanic {
			t.Fatalf("testing %s should throw an error", test.Name)
		} else if err != nil && !test.ShouldPanic {
			t.Fatalf("testing %s throw an error but should return %s", test.Name, test.Result)
		} else if test.Result != result {
			t.Fatalf("testing %s return %s. Want %s", test.Name, result, test.Result)
		}
	}

}
