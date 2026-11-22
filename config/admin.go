package config

import "goravel/app/facades"

func init() {
	config := facades.Config()
	config.Add("admin", map[string]any{
		"username": config.Env("ADMIN_USERNAME", ""),
		"password": config.Env("ADMIN_PASSWORD", ""),
	})
}
