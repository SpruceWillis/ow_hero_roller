package gifProvider

type GifProvider interface {
	GetGifUrl(string) (string, error)
	EmbedMessage(string) string
}
