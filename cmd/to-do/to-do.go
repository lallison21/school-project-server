package main

import (
	"fmt"

	"github.com/lallison21/to-do/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// в вермии до 1.21 импортировали как сторонюю библиотеку. В 1.21 использовать log/slog
	// TODO: init logger: slog

	// TODO: init storage: postgresql

	// TODO: init router: chi, "chi render"

	// TODO: run server:
}
