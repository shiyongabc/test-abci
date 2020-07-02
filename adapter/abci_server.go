package adapter

import (
	"github.com/shiyongabc/test-abci/abcitype"
)
type ABCIServer interface {
	Create(obj map[string]interface{}) (result string,errorMessage *abcitype.ErrorMessage)
}



