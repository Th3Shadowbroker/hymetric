package comparison

func IfThenElse[T any](comparison bool, trueValue T, falseValue T) T {
	if comparison {
		return trueValue
	}
	return falseValue
}
