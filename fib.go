package main

import ("fmt"; "os"; "strconv")

func fib(n int) int {
  a, b := 0, 1
  for i := 0; i < n; i++ {
    a, b = b, a+b
  }
  return a
}

func main() {
  for _, arg := range os.Args[1:] {
    int_val, err := strconv.Atoi(arg)
    if err == nil {
      fmt.Println(fib(int_val));
    }
  }
}
