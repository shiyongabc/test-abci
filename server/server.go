package server

import (
	"github.com/labstack/echo"
)

// APIServer is a http server could access mysql api
type APIServer struct {
	*echo.Echo               // echo web server
	app *Application // database api adapter
}

// New create a new APIServer instance
func New(app *Application) *APIServer {
	server := &APIServer{}
	server.Echo = echo.New()

	
	server.app = app

	//c, err := redis.Dial("tcp", redisHost)
	//if err != nil {
	//	fmt.Println("Connect to redis error", err)
	//	c=nil
	//}else{
	//	fmt.Println("Connect to redis success")
	//}

	mountEndpoints(server.Echo, app)
	return server
}

// Start server
func (server *APIServer) Start(address string) *APIServer {
	//server.StartMetadata()
	server.Logger.Infof("server start at %s", address)
	server.Logger.Fatal(server.Echo.Start(address))
	return server
}

