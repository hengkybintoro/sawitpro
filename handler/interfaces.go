// This file contains the interfaces for the handler layer.
// The service layer is responsible for interacting with the service layer.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package handler

import (
	"github.com/labstack/echo/v4"
)

type ServerInterface interface {
	AddEstate(c echo.Context) error
	AddTree(c echo.Context) error
	GetEstateStats(c echo.Context) error
	GetDronePlan(c echo.Context) error
}
