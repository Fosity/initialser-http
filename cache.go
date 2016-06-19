package main

type Cache interface {
	Get(key string) []byte
	Set(key string, data []byte)
}
type AvatarCache struct {

}

