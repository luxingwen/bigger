package request

import (
	"bigger/models"
	"encoding/json"
	"github.com/astaxie/beego/httplib"
)

func Post(address string, m interface{}) ([]byte, error) {
	var data []byte
	if v, ok := m.([]byte); ok {
		data = v
	} else {
		b, err := json.Marshal(m)
		if err != nil {
			return nil, err
		}
		data = b
	}
	req := httplib.Post(address)
	req.Header("X-Bmob-Application-Id", models.BmobAppKey)
	req.Header("X-Bmob-REST-API-Key", models.BmobResterKey)
	req.Header("Content-Type", "application/json")
	req.Body(data)
	return req.Bytes()
}
