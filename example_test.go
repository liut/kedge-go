package kedge

import (
	"fmt"
)

func ExampleKedgeSession() {
	c := New()
	ts, err := c.Session()
	if err == nil {
		fmt.Println(ts.PeerPort, ts.Version)
	}
	// Output: 6881 1.2.14.0
}

func ExampleKedgeStats() {
	c := New()
	st, err := c.Stats()
	if err == nil {
		if st.Uptime > 0 {
			fmt.Println("OK")
		}
	}
	// Output: OK
}
