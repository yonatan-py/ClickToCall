package configs

func GetConfig(env string) map[string]any {
	if env == "prod" {
		return prodConfig()
	}
	return devConfig()
}
