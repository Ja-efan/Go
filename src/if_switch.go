package main

func main(){

  k := 1
  i := 2
  if k == 1{
    println("k is One")
  }

  if i == 1{
    println("i is One")
  } else if i == 2{
    println("i is two")
  } else {
    println ("Other")
  }

  var name string
  var category = 1

  switch category{
  case 1:
    name = "Paper Book"
  case 2:
    name = "eBook"
  case 3, 4:
    name = "Blog"
  default:
    name = "Other"
  }
  println(name)


}
