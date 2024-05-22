package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func CreateStack(actions ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(actions) - 1; i >= 0; i-- {
			action := actions[i]
			next = action(next)
		}

		return next
	}
}
