package main

import (
	"net/http"
	"pub/config"
	"pub/model"
	"pub/usecase"

	"github.com/labstack/echo/v4"
)

// var redisClient = redis.NewClient(&redis.Options{
// 	Addr: "localhost:6379",
// })

// var ctx = context.Background()

func main() {
	config.SetEnvFileToVariable()

	config.SetConnectionDB()
	defer config.CloseConnectionDB()
	// app := fiber.New()
	e := echo.New()
	ec := e.Group("/whatsapp/v1/")
	ec.POST("singleSend", func(c echo.Context) error {
		u := new(model.ReqInquiry)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
		result := usecase.SingleSend(*u)
		return c.JSON(http.StatusOK, result)
	})
	// ec.POST("cek", func(c echo.Context) error {
	// 	var user model.User
	// 	var respon model.Respon
	// 	if err := c.Bind(user); err != nil {
	// 		panic(err)
	// 	}

	// 	payload, err := json.Marshal(user)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	if err := redisClient.Publish(ctx, "send-user-dat", payload).Err(); err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("from : ", user.No, "to sender :", user.Index)
	// 	respon.Status = "00"
	// 	respon.No = user.No
	// 	return c.JSON(http.StatusOK, respon)

	// })

	// app.Listen(":3000")
	e.Logger.Fatal(e.Start(":8000"))

}
