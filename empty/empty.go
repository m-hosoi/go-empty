package empty

// CheckEmpty is empty
type CheckEmpty interface {
	IsEmpty() bool
}

// IsNotEmpty .
func IsNotEmpty(chk interface{}) bool {
	return !IsEmpty(chk)
}

// IsEmpty .
func IsEmpty(chk interface{}) bool {
	switch v := chk.(type) {
	case nil:
		return true
	case bool:
		if !v {
			return true
		}
	case int64:
		if v == 0 {
			return true
		}
	case int32:
		if v == 0 {
			return true
		}
	case int:
		if v == 0 {
			return true
		}
	case float64:
		if v == 0 {
			return true
		}
	case float32:
		if v == 0 {
			return true
		}
	case string:
		if v == "" {
			return true
		}
	case CheckEmpty:
		if v.IsEmpty() {
			return true
		}
	case *CheckEmpty:
		if (*v).IsEmpty() {
			return true
		}
	}
	return false
}

// IIf is chk ? a : b
func IIf(chk, a, b interface{}) interface{} {
	if IsEmpty(chk) {
		return b
	}
	return a
}
