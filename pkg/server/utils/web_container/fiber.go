package web_container

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type FiberWebContainer struct {
	*fiber.App
}

func NewFiberApp() WebContainer {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true, // 区分大小写
		StrictRouting: true, // 严格路由
		AppName:       "go fiber template",
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
	})
	app.Use(recover.New())
	app.Use(logger.New())

	return &FiberWebContainer{app}
}

func (f *FiberWebContainer) RunWebContainer(addr string) error {
	return f.Listen(addr)
}

func InitApp(app *fiber.App) {
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Next:   nextLogger,
		Format: "[INFO-${locals:requestid}]${time} pid: ${pid} status:${status} - ${method} path: ${path} queryParams: ${queryParams} body: ${body}\n resBody: ${resBody}\n error: ${error}\n",
	}))
}

func nextLogger(c *fiber.Ctx) bool {
	// return true 会跳过本次中间件执行
	if c.Path() == "/v1/upload" {
		return true
	}
	return false
}
