package collatzconjecture
import "errors"

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
    	return 0, errors.New("A value less than or equal to 0 entered")
    }
	count := 0
	for n != 1 {
        if n % 2 == 1{
            n = n * 3 + 1
        } else {
            n /= 2
        }
        count++
    }
    return count, nil
}
