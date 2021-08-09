package main

import (
	"github.com/Jarover/Go2/lesson7"
)

func main() {

	lesson7.InQuery("SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?", false, []int{1, 6, 234}, 555)

}
