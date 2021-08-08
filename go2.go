package main

import (

	"Go2/lesson7"
)

func main() {
	//lesson4.WithWorkers()
	//lesson4.WithSigterm()

	lesson7.InQuery("SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?",false, []int{1, 6, 234}, 555)

}
