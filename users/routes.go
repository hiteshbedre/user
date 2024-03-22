package main

import (
 "fmt"
 "github.com/gobuffalo/buffalo"
)

func fetchArticle(router *fiber.App) {
  app = buffalo.New(buffalo.Options{
			Env:         "env",
			SessionName: "_toodo_session",
	  })

  app.GET("/users/register", handler.getRegisteredUsers)
}

