// main.go

package main

import "os"

func main() {
	a := App{}
	a.Initialize()

	a.Run(":" + os.Getenv("PORT"))
}
