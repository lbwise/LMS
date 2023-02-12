package routes

import (
	"encoding/json"
	"io"
	"errors"
	"github.com/gin-gonic/gin"

	lg "github.com/lbwise/LMS/log"
)

func getData(obj *any, c *gin.Context) (any, error) {
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		lg.Logger.Println(err)
		return nil, errors.New("Unable to read given data")
	}
	json.Unmarshal(data, obj)
	return data, nil
}