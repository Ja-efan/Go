package main

func main() {
  // basic for
  sum := 0
  for i := 1 ; i<=100 ; i++ {
    sum += i
  }
  println(sum)

  // for likes while
  n := 1
  for n < 100 {
    n *= 2
  }
  println(n)

  // for to infinite loop
  //for {
  //  println("Infinite loop")
  //}

  names := [] string {"홍길동","이순신","강감찬"}

  for index,name := range names{
    println(index, name)
  }

}
