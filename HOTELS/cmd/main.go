package main

import (
	"log"
	"os"

	"github.com/ruziba3vich/hotello-hotels/cmd/app"
	"github.com/ruziba3vich/hotello-hotels/internal/pkg/config"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	config, err := config.LoadConfig()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Fatal(app.Run(config, logger))
}
