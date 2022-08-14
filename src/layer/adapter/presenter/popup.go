package presenter

import (
	"github.com/demo/layer/adapter"
	"github.com/demo/layer/usecase/popup"
	"github.com/junyaU/evbus"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Popup struct {
	bus adapter.MessagePublisher
}

func NewPopup(bus evbus.Bus) *Popup {
	return &Popup{
		bus: bus,
	}
}

func (p *Popup) Create(c echo.Context) error {
	var req struct {
		OwnerId       string `validate:"required" form:"ownerId"`
		StartTime     string `validate:"required" form:"startTime"`
		EndingTime    string `validate:"required" form:"endingTime"`
		MainName      string `validate:"required,min=1,max=20" form:"mainName"`
		ShoulderName  string `validate:"required,min=5,max=50" form:"shoulderName"`
		PromotionText string `validate:"required,min=100,max=500" form:"promotionText"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	hostCommand := popup.NewHostCommand(req.OwnerId, req.StartTime, req.EndingTime, req.MainName, req.ShoulderName, req.PromotionText)
	if err := p.bus.Publish("popup:host", hostCommand); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (p *Popup) UpdateMerchandises(c echo.Context) error {
	var req struct {
		PopupId        string   `validate:"required" form:"popupId"`
		OwnerId        string   `validate:"required" form:"ownerId"`
		MerchandiseIds []string `validate:"required" form:"endingTime"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	hostCommand := popup.NewExhibitedCommand(req.PopupId, req.OwnerId, req.MerchandiseIds)
	if err := p.bus.Publish("popup:exhibit", hostCommand); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)

}
