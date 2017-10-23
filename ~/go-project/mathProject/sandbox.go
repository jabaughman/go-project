package main
import (
    "fmt"
    "math/big"
    "math"
  )

func main(){
  fmt.Println("Hello, World!")
  i1, i2, i3 := 34, 54, 78
  intsum := i1 + i2 + i3
  fmt.Println("Integer sum: ", intsum)

  f1, f2, f3 := 23.5, 65.1, 76.3
  floatsum := f1 + f2 + f3
  fmt.Println("Float sum: ", floatsum)

  var b1, b2, b3, bigSum big.Float
  b1.SetFloat64(23.5)
  b2.SetFloat64(65.1)
  b3.SetFloat64(76.3)

  bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
  fmt.Printf("BigSum = %.10g\n", &bigSum)

  circleRadius := 15.5
  circumference := circleRadius * math.Pi

  fmt.Printf("Circumference: %.2f\n", circumference)
}
