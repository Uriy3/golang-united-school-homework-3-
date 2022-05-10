package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	 num := len(input)
        for i := num - 1; i >= 0; i-- {
            result = append(result, input[i])
        }
        return result
    }
