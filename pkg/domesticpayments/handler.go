package domesticpayments

import (
	"github.com/kaanaktas/openbanking-paymentinitiation/api"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterHandler(e *echo.Echo, service Service) {
	e.GET("/:aspspId/domestic-payments/:cid", initiatePayment(service))
}

func initiatePayment(s Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		rid := c.Response().Header().Get(echo.HeaderXRequestID)
		aspspId := c.Param("aspspId")
		if aspspId == "" {
			return c.JSON(http.StatusBadRequest, api.JsonResponse(rid, "aspspId can't be empty"))
		}

		cid := c.Param("cid")
		if cid == "" {
			return c.JSON(http.StatusBadRequest, api.JsonResponse(rid, "cid can't be empty"))
		}

		res, err := s.DomesticPayment(cid, aspspId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, api.JsonResponse(rid, err.Error()))
		}

		return c.JSON(http.StatusOK, res)
	}
}
