package main

import "fmt"

func main() {
    robot := Robot{}
    robot.table = Table{0,0,4,4}
    robot.placed = true
    fmt.Println("Robot is: ", robot)
    robot.move
    fmt.Println("Robot is: ", robot)
}
