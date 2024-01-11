package iteration

func Repeat(value string, repeatTimes int) string {
	var repeats string
	if repeatTimes == 0 {
		repeatTimes = 5
	}

	for i := 0; i < repeatTimes; i++ {
		repeats += value
	}
	return repeats
}
