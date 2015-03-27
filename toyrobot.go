package main

func main() {
    robot := NewRobot(Table{0,0,4,4})
    robot.place(0,0,"north")
    robot.move()
    robot.right()
    robot.left()
    robot.report()
}
