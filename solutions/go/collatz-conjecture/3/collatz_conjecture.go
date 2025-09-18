package collatzconjecture
import "errors"

var ErrNonPositive = errors.New("input must be a positive integer")

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
    	return 0, ErrNonPositive
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
