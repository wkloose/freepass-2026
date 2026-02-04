package main

import (
	"github.com/Hisyam/freepass-2026/initializers"
)

func init() {
	initializers.LoadEnv()     
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {

}