package main

type Table struct {
   x,y,xx,yy int
}

func (t Table) Contains(x int, y int) bool {
  if x >= t.x && x <= t.xx && y >= t.y && y <= t.yy { 
    return true 
  } else { 
    return false
  }
}
