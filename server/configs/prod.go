package configs

import "google.golang.org/api/option"

func prodConfig() map[string]any {
	opt := option.WithCredentials(nil)
	return map[string]any{
		"firebaseOptions": opt,
	}
}
