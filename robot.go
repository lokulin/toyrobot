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
  return &Robot{ compass: []Direction{Direction{"north",0, 1},
                          Direction{"east", 1, 0},
                          Direction{"south",0,-1},
                          Direction{"west",-1, 0}},
                 placed: false }
}

func (r *Robot) move() {
    if r.placed {
        r.x += r.compass[0].x
        r.y += r.compass[0].y
    }
}

func (r *Robot) left() {
  if r.placed {
      r.compass = append(r.compass[len(r.compass)-1:],r.compass[:len(r.compass)-1]...)
  }
}

func (r *Robot) right() {
  if r.placed {
    r.compass = append(r.compass[1:],r.compass[:1]...)
  }
}

func (r *Robot) place(xx int, yy int, d string) {
  r.x, r.y = xx, yy
  r.placed = true
}

func (r *Robot) report() {
  if r.placed {
    fmt.Println(r.x, r.y, r.compass[0].name)
  }
}
