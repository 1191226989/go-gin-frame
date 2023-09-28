package hash

import "go-gin-frame/config"

var _ Hash = (*hash)(nil)

type Hash interface {
	i()

	// HashidsEncode 加密
	HashidsEncode(params []int) (string, error)

	// HashidsDecode 解密
	HashidsDecode(hash string) ([]int, error)
}

type hash struct {
	secret string
	length int
}

func New(secret string, length int) Hash {
	return &hash{
		secret: secret,
		length: length,
	}
}

func (h *hash) i() {}

func HashidsEncode(params []int) (string, error) {
	cfg := config.Get().HashIds
	return New(cfg.Secret, cfg.Length).HashidsEncode(params)
}

func HashidsDecode(hash string) ([]int, error) {
	cfg := config.Get().HashIds
	return New(cfg.Secret, cfg.Length).HashidsDecode(hash)
}
