package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// Check returns a message indicating that the service is awake and listening
func Check(c echo.Context) error {
	return c.JSON(http.StatusOK, "Uh, we had a slight weapons malfunction, but uh... everything's perfectly all right now. We're fine. We're all fine here now, thank you. How are you?")
}
