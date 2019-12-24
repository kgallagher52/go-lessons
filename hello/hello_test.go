package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

// 1. In directory run 'go mod init example.com
// 2. Then run go test
/* 3.
output PASS
ok      example.com     0.291s
*/
