package awsat

func String(v string) *string {
	return &v
}

func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

func Int64(v int64) *int64 {
	return &v
}

func Int64AsIntValue(v *int64) int {
	if v != nil {
		return int(*v)
	}
	return 0
}

func Bool(v bool) *bool {
	return &v
}

func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}
