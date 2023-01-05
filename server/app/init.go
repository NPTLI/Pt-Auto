package app

import (
	"context"
	"io/ioutil"
	"pt-auto/public"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

func Config() {
	data, _ := ioutil.ReadFile("env.yaml")
	err := yaml.Unmarshal(data, &public.Configs)
	if err != nil {
		panic("decode error")
	}
}

func Redis() {
	public.Redis = redis.NewClient(&redis.Options{
		Addr:     public.Configs.Redis.Addr,
		Password: public.Configs.Redis.Password,
		DB:       public.Configs.Redis.DB,
	})

	_, err := public.Redis.Ping(context.Background()).Result()
	if err != nil {
		panic("redis error:" + err.Error())
	}
}

func ZapLog() {
	public.Log = zap.NewExample()
}

func Schedule() {
	// s := gocron.NewScheduler(time.UTC)

	// _, err := s.Every("10m").Do("")
	// if err != nil {
	// 	public.Log.Error(err.Error())
	// }
}

func Web() {
	app := fiber.New()

	// app.All("/api/getListZh", getListZh())

	err := app.Listen(":666")
	if err != nil {
		public.Log.Error(err.Error())
	}
}
