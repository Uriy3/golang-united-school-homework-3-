package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	var sum float32
        result = 0.0
        for , val := range input {
            sum += val
        }
        result = sum / 15.0
        return result	
