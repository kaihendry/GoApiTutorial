// main.go

package main

import "os"

func main() {
	a := App{}
	a.Initialize("root", "", "rest_api_example")

	a.Run(":" + os.Getenv("PORT"))
}
