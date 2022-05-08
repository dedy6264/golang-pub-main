package usecase

import (
	"fmt"
	"pub/model"
	"pub/repository"
	"time"
)

func SingleSend(req model.ReqInquiry) model.RespGlobal {
	var result model.RespGlobal
	t := time.Now()
	hasil := repository.SingleSend(req)
	fmt.Println("start request:", req.IDSend)

	if hasil.Status != "00" {
		result.Status = hasil.Status
		result.StatusDateTime = t
		result.StatusDesc = hasil.StatusDesc
		return result
	}
	return hasil
}
