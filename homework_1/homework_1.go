/* Напишите функцию, производящую следующие вычисления.
S = 2 × x + y ^ 2 − 3/z, где x — int16, y — uint8, a z — float32.
Тип S должен быть во float32. */

package main

import (
	"fmt"
)

func main() {
	var (
		x int16   = 3
		y uint8   = 4
		z float32 = 2.4
	)

	S := 2*float32(x) + float32(y*y) - 3/z
	fmt.Println(S)

}
