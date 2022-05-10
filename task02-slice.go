package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	keys := make([]int, len(input))
        i := 0
        for k := range input {
            keys[i] = k
            i++
        }
        sort.Ints(keys)
        for , k := range keys {
            result = append(result, input[k])
        }
        return result
    }
