package main

type Table struct {
   x,y,xx,yy int
}

func (t Table) Contains(x int, y int) Table {
  return t
}