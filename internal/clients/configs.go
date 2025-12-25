package clients

import "time"

type TheMealsDbClientConfig interface {
	Url() string
	Timeout() time.Duration
}
