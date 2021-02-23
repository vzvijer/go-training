package flatten

func Flatten(in interface{}) []interface{} {
	var val interface{}
	out := []interface{}{}
	a, ok := in.([]interface{})
	if ok {
		var elem interface{}
		for _, elem = range a {

			val, ok = elem.(int)
			if ok {
				out = append(out, val)
			} else {
				val, ok = elem.([]interface{})
				if ok {
					for _, val = range Flatten(val) {
						out = append(out, val)
					}
				}
			}
		}
	}
	return out
}
