package web

import (
	"fmt"
	"github.com/jackparsonss/grpc-monorepo/web-server/internal/application/domain"
	"github.com/jackparsonss/grpc-monorepo/web-server/internal/ports"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Adapter struct {
	api  ports.APIPort
	port int
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a *Adapter) Run() {
	e := echo.New()

	g := e.Group("/inventory")
	g.POST("", a.addStock)
	g.GET("/:id", a.getStockItem)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", a.port)))
}

func (a *Adapter) addStock(c echo.Context) error {
	i := domain.InventoryItem{}
	if err := c.Bind(&i); err != nil {
		return err
	}

	id, err := a.api.AddNewStock(&i)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	return c.JSON(http.StatusOK, id)
}

func (a *Adapter) getStockItem(c echo.Context) error {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	item, err := a.api.GetStockItem(int64(intId))
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	return c.JSON(http.StatusOK, item)
}
