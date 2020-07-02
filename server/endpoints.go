package server

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/shiyongabc/test-abci/abcitype"
	"github.com/tendermint/tendermint/abci/types"
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
		userId:=c.QueryParam("userId")
		operateKey:=c.QueryParam("operateKey")
		if errorMessage!=nil{
			return c.JSON( http.StatusBadRequest, errorMessage)
		}
		//api.Create(payload)
		println("tx=",userId)
		bytett,_:=json.Marshal(payload)
		txParam:=userId+"-"+operateKey+"="+string(bytett)
		println("tx=",txParam)
		var req types.RequestDeliverTx
		//byteData , _ := json.Marshal(txParam)
		req.Tx=[]byte(txParam)
		//types.ResponseDeliverTx

		res:=app.DeliverTx(req)
		var reqCh types.RequestCheckTx
		app.CheckTx(reqCh)
		app.Commit()

		return c.JSON( http.StatusOK, res)
	}

}

func endpointTableGet(app *Application) func(c echo.Context) error {
//api.Create()
	return func(c echo.Context) error {
		var queryReq types.RequestQuery
		print(c.QueryParam("data"))
		queryReq.Data=[]byte(c.QueryParam("data"))
		print(queryReq.Data)
		queryReq.Prove=true
		res:=app.Query(queryReq)
		return c.JSON( http.StatusOK, res)
	}

}
