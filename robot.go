package main

import (
  "strings"
  "fmt"
  "sort"
)

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
  return &Robot{ compass: []Direction{Direction{"north",0, 1},
                          Direction{"east", 1, 0},
                          Direction{"south",0,-1},
                          Direction{"west",-1, 0}},
                 placed: false,
                 table: table }
}

func (r *Robot) Move() {
  if r.placed {
      r.x += r.compass[0].x
      r.y += r.compass[0].y
  }
}

func (r *Robot) Left() {
  if r.placed {
      r.compass = append(r.compass[len(r.compass)-1:],r.compass[:len(r.compass)-1]...)
  }
}

func (r *Robot) Right() {
  if r.placed {
    r.compass = append(r.compass[1:],r.compass[:1]...)
  }
}

func (r *Robot) Place(xx int, yy int, d string) {
  if r.table.Contains(xx, yy) {
    r.x, r.y = xx, yy

    r.placed = true
  }
}

func (r *Robot) Report() {
  if r.placed {
    fmt.Printf("%d,%d,%s\n", r.x, r.y, strings.ToUpper(r.compass[0].name))
  }
  fmt.Println(CompassIndex(r.compass, "north"))
}

func CompassIndex(c []Direction, s string) int {
  fmt.Println(c)
  i := sort.Search(len(c), func(i int) bool { fmt.Println(c)
    fmt.Println(s)
    return c[i].name == s })
  return i
}
   
