package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("This is a simple fmt print!")
	log.Println("This is a simple log print!")

	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.Println("This is another simple log print!")

	// log.Panic("Something went wrong!")

	// log.Fatal("Something went very wrong!")

	logger.Info("This is an info log!")
	logger.Warning("This is an warning log!")
	logger.Error("This is an error log!")

	logger.SetLevel(logger.WarningLevel)

}
