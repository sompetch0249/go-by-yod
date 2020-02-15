package fizzbuzz

import "testing"

func TestFizzBuzzSayOrigin(t *testing.T) {
	t.Run("given 1 say 1", func(t *testing.T) {
		var given = 1
		var want = "1"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
	t.Run("given 2 say 2", func(t *testing.T) {
		var given = 2
		var want = "2"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
	t.Run("given 4 say 4", func(t *testing.T) {
		var given = 4
		var want = "4"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})
	t.Run("given 7 say 7", func(t *testing.T) {
		var given = 7
		var want = "7"

		var get = Say(given)
		if want != get {
			t.Errorf("given %v want %q but get %q", given, want, get)
		}
	})

}
