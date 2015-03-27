package main

import "fmt"

type Robot struct {
  table Table
  x, y int
  compass []Direction
  placed bool
}

type Direction struct {
  name string
  x, y int
}

func NewRobot(table Table) *Robot {
  return &Robot{ compass: []Direction{Direction{"NORTH",0, 1},
                          Direction{"EAST", 1, 0},
                          Direction{"SOUTH",0,-1},
                          Direction{"WEST",-1, 0}},
                 placed: false,
                 table: table }
}

func (r *Robot) LookingAt() (x int, y int) {
  x = r.x + r.compass[0].x
  y = r.y + r.compass[0].y
  return x, y
}

func (r *Robot) Move() {
  if r.placed && r.table.Contains(r.LookingAt()) {
      r.x, r.y = r.LookingAt()
  }
}

func (r *Robot) Left() {
  if r.placed {
      r.compass = append(r.compass[len(r.compass)-1:],r.compass[:len(r.compass)-1]...)
  }
}

func (r *Robot) Right() {
  if r.placed { r.compass = append(r.compass[1:],r.compass[:1]...) }
}

func (r *Robot) Place(xx int, yy int, d string) {
  if r.table.Contains(xx, yy) {
    r.x, r.y = xx, yy
    r.placed = true
    i := 0
    for _,c := range r.compass { if c.name != d { r.Right(); i++ } }
    if i == len(r.compass) { r.placed = false }
  }
}

func (r *Robot) Report() {
  if r.placed {
    fmt.Printf("%d,%d,%s\n", r.x, r.y, r.compass[0].name)
  }
}
