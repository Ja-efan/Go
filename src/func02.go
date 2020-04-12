package main

func main () {
  msg := "HELLO"
  say (&msg)
  println(msg)
}

// Pass by Reference
func say (msg *string) {
  println(*msg)
  *msg = "Changed"
}
