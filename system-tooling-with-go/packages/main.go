package main

import (
	"fmt"
	"packages/core"
	"packages/utils"
)

func main() {
	friendName := core.Names["friend"]
	friendSays := utils.Greeting(friendName, "Oh hello there!")
	fmt.Println(friendSays)

	enemyName := core.Names["enemy"]
	enemySays := utils.Greeting(enemyName, friendName)
	fmt.Println(enemySays)
}
