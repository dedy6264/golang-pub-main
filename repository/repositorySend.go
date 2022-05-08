package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"pub/config"
	"pub/model"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/lib/pq"
)

func SingleSend(req model.ReqInquiry) model.RespGlobal {
	var result model.RespGlobal
	var isiResult model.RespReqSend
	db := config.DbConn()
	var wa, text string
	var idClient, idMessage, status int
	err := db.QueryRow(`
	SELECT 	
	status,
	message_id,
	client_id
	FROM send
	WHERE id = $1`, req.IDSend).
		Scan(
			&status,
			&idMessage,
			&idClient)
	if err == sql.ErrNoRows {
		result.Status = "31"
		result.StatusDateTime = time.Now()
		result.StatusDesc = "Account Transaction Number Not Found"
		return result
	}
	if err != nil {
		result.Status = "81"
		result.StatusDateTime = time.Now()
		result.StatusDesc = err.Error()
		return result
	}
	if status == 1 {
		result.Status = "81"
		result.StatusDateTime = time.Now()
		result.StatusDesc = "Sorry guys, this message has been sent"
		return result
	}
	err = db.QueryRow(`
	SELECT 	
	client_phone
	FROM client
	WHERE id = $1`, idClient).
		Scan(
			&wa)
	if err == sql.ErrNoRows {
		result.Status = "31"
		result.StatusDateTime = time.Now()
		result.StatusDesc = "Account Transaction Number Not Found"
		return result
	}
	if err != nil {
		result.Status = "81"
		result.StatusDateTime = time.Now()
		result.StatusDesc = err.Error()
		return result
	}
	err = db.QueryRow(`
	SELECT 	
	message
	FROM message
	WHERE id = $1`, idMessage).
		Scan(
			&text)
	if err == sql.ErrNoRows {
		result.Status = "31"
		result.StatusDateTime = time.Now()
		result.StatusDesc = "Account Transaction Number Not Found"
		return result
	}
	if err != nil {
		result.Status = "81"
		result.StatusDateTime = time.Now()
		result.StatusDesc = err.Error()
		return result
	}
	///// db transaction start; save to send DB
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println("DB Transaction Begin ", err.Error())
		result.Status = "81"
		result.StatusDesc = "Error DB Transaction " + err.Error()
		result.StatusDateTime = time.Now()
		return result
	}
	//Insert Data DB Transaction Merchant
	status = 1
	_, err = tx.ExecContext(ctx, `UPDATE send SET
	status = $1,
	updated_at = $2 WHERE id = $3`,
		status,
		time.Now(), req.IDSend)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("update drivers: unable to rollback :", rollbackErr)
		}
		if err, ok := err.(*pq.Error); ok {
			fmt.Println("UpdateMerchantSaving : DB Transaction Rollback", err)
			result.Status = "99"
			result.StatusDesc = strings.Replace(err.Detail, "Key (merchant_", "(", -1)
			result.StatusDateTime = time.Now()
			return result
		}
	}
	randnumb := rand.Intn(3-0) + 1
	index := strconv.Itoa(randnumb)
	publish := model.User{
		No:    wa,
		Text:  text,
		Index: index,
	}
	//send to publisher
	var redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	payload, err := json.Marshal(publish)
	if err != nil {
		// panic(err)
		result.Status = "81"
		result.StatusDesc = "Error Marshal " + err.Error()
		result.StatusDateTime = time.Now()
		return result
	}

	if err := redisClient.Publish(ctx, "send-user-dat", payload).Err(); err != nil {
		// panic(err)
		result.Status = "81"
		result.StatusDesc = "Error publisher " + err.Error()
		result.StatusDateTime = time.Now()
		return result
	}
	fmt.Println("Send from channel:", index, "to number :", wa)

	if err := tx.Commit(); err != nil {
		fmt.Println("DB Transaction Begin ", err.Error())
	}

	/////
	isiResult.No = wa
	isiResult.Text = text
	isiResult.Index = strconv.Itoa(idMessage) + strconv.Itoa(idMessage)
	result.Status = "00"
	result.StatusDateTime = time.Now()
	result.StatusDesc = "Success"
	result.Result = isiResult
	return result
}
