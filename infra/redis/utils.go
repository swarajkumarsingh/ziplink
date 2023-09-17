package redisUtils

func IncrementCounter() (int64, error) {
	counterKey := "my_counter"
	result, err := Incr(counterKey)
	if err != nil {
		return 0, err
	}
	return result, nil
}
