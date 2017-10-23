package main

import (
  "fmt"
  "sort"
)

func main() {
  var colors [3]string
  colors[0] = "Red"
  colors[1] = "Green"
  colors[2] = "Blue"
  fmt.Println(colors)
  fmt.Println(colors[1])

  var numbers = [5]int{5,3,1,2,4}
  fmt.Println(numbers)

  fmt.Println("Number of colors:", len(colors))
  fmt.Println("Number of numbers:", len(numbers))

  //Slices
  var colors1 = []string{"Red", "Green", "Blue"}
  fmt.Println(colors1)

  colors1 = append(colors1, "Purple")
  fmt.Println(colors1)

  colors1 = append(colors1[1:len(colors1)]) //removes first item of the slice
  fmt.Println(colors1)

  colors1 = append(colors1[:len(colors1)-1]) //removes last item of the slice
  fmt.Println(colors1)

  numbers1 := make([]int, 5, 5)
  numbers1[0] = 1
  numbers1[1] = 7
  numbers1[2] = 32
  numbers1[3] = 12
  numbers1[4] = 156
  fmt.Println(numbers1)

  numbers1 = append(numbers1, 235)
  fmt.Println(numbers1)
  fmt.Println(cap(numbers1))
  sort.Ints(numbers1)
  fmt.Println(numbers1)
}
