package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

var names map[string]bool

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if names == nil {
		names = make(map[string]bool)
	}

	if r.name == "" {
		for {
			name := ""
			for i := 0; i < 2; i++ {
				name = name + fmt.Sprintf("%c", 'A'+rand.Intn(26))
			}
			for i := 0; i < 3; i++ {
				name = name + fmt.Sprintf("%d", rand.Intn(10))
			}
			if !names[name] {
				r.name = name
				names[name] = true
				break
			}
		}
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	rand.Seed(time.Now().Unix())
	r.name = ""
}
