package abci

import (
	"github.com/shiyongabc/test-abci/abcitype"
)
// MysqlAPI
type ABCIServerImpl struct {
	useInformationSchema bool
}

// NewMysqlAPI create new MysqlAPI instance
func NewABCIServerImpl() (api *ABCIServerImpl) {
	api = &ABCIServerImpl{}
	return

}

// Create by Table name and obj map
func (api *ABCIServerImpl) Create(obj map[string]interface{}) (result string,errorMessage *abcitype.ErrorMessage) {

	return "",nil
}




