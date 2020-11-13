package auth

import "time"

type EpochGetter interface {
	Epoch() int64
}

type epochGetterImpl struct{}

func NewEpochGetter() EpochGetter {
	return &epochGetterImpl{}
}

func (e *epochGetterImpl) Epoch() int64 {
	return time.Now().Unix()
}
