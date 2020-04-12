package main

func main() {
  total := sum (1,7,3,5,9)
  println(total)

  count, total2 := sum2(2,4,6,8,10)
  println(count, total2)
}

func sum (nums ...int) int {
  s := 0
  for _,n := range nums{
    s += n
  }
  return s
}

func sum2 (nums ...int) (int, int) {
  s := 0
  count := 0
  for _,n := range nums{
    s += n
    count ++
  }
  return count, s
}
