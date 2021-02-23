package secret

func Handshake(val uint) []string {
	events := []string{"wink", "double blink", "close your eyes", "jump"}
	n := len(events)
	reverse := ((val & (1 << n)) != 0)

	res := []string{}
	for i := 0; i < len(events); i++ {
		if reverse {
			if val&((1<<(n-1))>>i) != 0 {
				res = append(res, events[n-1-i])
			}
		} else {
			if val&(1<<i) != 0 {
				res = append(res, events[i])
			}
		}
	}
	return res
}
