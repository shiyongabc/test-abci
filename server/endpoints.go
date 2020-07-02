package server

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/shiyongabc/test-abci/abcitype"
	"github.com/tendermint/tendermint/abci/abcitype"
	"net/http"
)

// mountEndpoints to echo server
func mountEndpoints(s *echo.Echo, app *Application) {

//s.GET("/api/swagger-ui.html", endpointSwaggerUI).Name = "Swagger UI"

s.GET("/api/blockChain/query", endpointTableGet(app)).Name = "Retrive Some Records"
s.POST("/api/blockChain/trade", endpointTableCreate(app)).Name = "Retrive Some Records"


}
func bodyMapOf(c echo.Context) (jsonMap map[string]interface{}, errorMessage *abcitype.ErrorMessage) {
	jsonMap = make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if (err != nil) {
		errorMessage = &abcitype.ErrorMessage{abcitype.ERR_PARAMETER, err.Error()}
	}
	return
}

func endpointTableCreate(app *Application) func(c echo.Context) error {
	//api.Create()
	return func(c echo.Context) error {
		payload, errorMessage := bodyMapOf(c)
		if errorMessage!=nil{
			return c.JSON( http.StatusBadRequest, errorMessage)
		}
		//api.Create(payload)
		var req abcitype.RequestDeliverTx
		req
		app.DeliverTx(req)
		return c.JSON( http.StatusOK, "blockchain trade start!")
	}

}

func endpointTableGet(app *Application) func(c echo.Context) error {
//api.Create()
	return func(c echo.Context) error {
		app.Query()
		return c.JSON( http.StatusOK, "blockchain trade start!")
	}

}
