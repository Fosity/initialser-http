package main

type Cache interface {
	Get(key string) []byte
	Set(key string, data []byte)
}

type Config struct {
	MaxItems int
	MaxBytes int64
}

type AvatarCache struct {

}

