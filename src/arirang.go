package main

func main() {
  // Raw String Literal. 복수라인.
  rawLiteral := `아리랑\n
아리랑\n
  아라리요`

  // Interpreted String rawLiteral
  interLiteral := "아리랑아리랑\n아라리요"
  // interLiteral := "아리랑아리랑 \n"+
  //                    "아라리요"

  println(rawLiteral)
  println()
  println(interLiteral)
}
