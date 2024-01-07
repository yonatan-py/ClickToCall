package configs

import "google.golang.org/api/option"

func devConfig() map[string]any {
	return map[string]any{
		"firebaseOptions": option.WithCredentialsFile("./click-to-call.json"),
	}
}
