package main

import (
	"errors"
	"fmt"
)

func main() {
	sum1, _ := sumOfAny(10, 20, 30, 40)
	fmt.Println("Sum:", sum1)

	sum2, _ := sumOfAny(10.23, 20.89, 30.54, 40.86)
	fmt.Println("Sum2:", sum2)

	sum3, _ := sumOfAny(10.23, 20, 30, 40.86, 50, 60, 10.34, 11, 23)
	fmt.Println("Sum2:", sum3)

	var num1 int = 1
	var num2 int8 = -2
	var num3 int16 = -3
	var num4 int32 = -4
	var num5 int64 = -5
	var num6 uint8 = 1
	var num7 uint16 = 2
	var num8 uint32 = 3
	var num9 uint64 = 4
	var num10 float32 = 1.1
	var num11 float64 = 2.1

	slice1 := []any{num1, num2, num3, num4, num5, num6, num7, num8, num9, num10, num11}

	sum4, _ := sumOfAny(slice1...)
	fmt.Printf("sum4: %0.2f\n", sum4)

	if sum5, err := sumOfAny(10, 20.1, "hello", true); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sum5", sum5)
	}

}

// int, int8,int16,int32,int64,uint8,uint16,uint32,uint64, float32,float64
// sumOfAny(10,10.5,200,200.45,10,11,12,34.5,.5,.6,.7)
func sumOfAny(vals ...any) (any, error) {
	var sum float64
	for _, val := range vals {
		switch val.(type) {
		case float64:
			sum = sum + val.(float64)
		case float32:
			sum = sum + float64(val.(float32))
		case int:
			sum = sum + float64(val.(int))
		case uint8:
			sum = sum + float64(val.(uint8))
		case uint16:
			sum = sum + float64(val.(uint16))
		case uint32:
			sum = sum + float64(val.(uint32))
		case uint64:
			sum = sum + float64(val.(uint64))
		case int8:
			sum = sum + float64(val.(int8))
		case int16:
			sum = sum + float64(val.(int16))
		case int32:
			sum = sum + float64(val.(int32))
		case int64:
			sum = sum + float64(val.(int64))
		default:
			return 0, errors.New("invalid type")
		}
	}
	return sum, nil
}
