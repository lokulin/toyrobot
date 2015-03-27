package main

import "fmt"

type Robot struct {
  table Table
  x, y int
  direction Direction
  placed bool
}

type Direction struct {
  x, y int
}

func (r *Robot) move() {
  if r.placed {
    r.x,r.y = r.x+1,r.y+1
    fmt.Println("Robot is: ", r)
  } else {
    fmt.Println("not placed")
  }
}

func (r *Robot) left() {

}

func (r *Robot) right() {

}

func (r *Robot) place(x int, y int, direction string) {

}
