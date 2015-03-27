package main

func main() {
    robot := NewRobot(Table{0,0,4,4})
    robot.Place(0,0,"north")
    robot.Report()
    robot.Move()
    robot.Report()
    robot.Right()
    robot.Report()
    robot.Left()
    robot.Report()
}
