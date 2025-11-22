package main

import (
	"context"
	"fmt"
	"network-testing-util/internal/app"
	"network-testing-util/internal/app/http/echo"
	"network-testing-util/internal/app/http/request"
)

func main() {
	fmt.Println("=====================================")
	fmt.Println("🚀 Network Testing Utility Server")
	fmt.Println("=====================================")

	ctx := context.Background()

	echoApp := echo.NewApp()
	echoApp.PrintCommonInfo()

	requestApp := request.NewApp()
	requestApp.PrintCommonInfo()

	mainApp := app.NewApp([]app.Executable{
		echoApp,
		requestApp,
	})

	mainApp.Run(ctx)
}
