package xrange

func XRange[T int | int8 | int16 | int32 | int64](min, max T) []T {
	var result = []T{}
	for i := min; i < max; i++ {
		result = append(result, i)
	}
	return result
}

func XRangeWithStep[T int | int8 | int16 | int32 | int64](min, max, step T) []T {
	var result = []T{}
	if step <= 0 {
		step = 1
	}
	for i := min; i < max; i += step {
		result = append(result, i)
	}
	return result
}

func Range[T int | int8 | int16 | int32 | int64](min, max T) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for i := min; i < max; i++ {
			ch <- i
		}
	}()
	return ch
}

func RangeWithStep[T int | int8 | int16 | int32 | int64](min, max, step T) <-chan T {
	ch := make(chan T)
	if step <= 0 {
		step += 1
	}
	go func() {
		defer close(ch)
		for i := min; i < max; i += step {
			ch <- i
		}
	}()
	return ch
}
