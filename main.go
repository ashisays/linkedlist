package main

import "fmt"

type node struct {
  val int
  next *node
}

func (n *node) insert(value int){
  temp := n
  for temp.next != nil{
    temp = temp.next
  }
  // fmt.Println(temp)
  temp.next = &node{val:value}
}

func (n *node) show(){
  temp := n
  for temp.next != nil{
    fmt.Printf("%v-->",temp.val)
    temp= temp.next
  }
}
//[1,2,3]
func (n *node) reverse() *node {
  temp := n
  var n1,temp2 *node
  n1 = &node{val:temp.val}
  for temp.next != nil{
    temp2 = n1
    temp = temp.next
    n1 = &node{val:temp.val}
    n1.next = temp2     
  }
  return n1
}

func main() {
  fmt.Println("Hello World")
  head := &node{val:1}
  head.insert(2)
  head.insert(3)
  head.show()
  fmt.Println("\n************")
  head = head.reverse()
  head.show()
}