package main

import "fmt"

type NegativeNumberError struct {
    number int
}

func (de *NegativeNumberError) Error() string {
	str := `
Do not support negative number
number: %d
	`
    return fmt.Sprintf(str, de.number)
}

func diff(x int)(result int, negative_error string) {
    if (x > 0) {
        result = -x
        return result, ""
	}else{
		data := NegativeNumberError{
			number: x,
		}
		negative_error := data.Error()
		return x, negative_error
	}
}

func main(){
	res, err := diff(-1)
	fmt.Println(res,err)
}