package faiconn

type AIConn interface {
	SendMessage(message string) (string, error)

	GetModel() string

	GetProvider() string

	Close() error
}

// AI配置
type AIConfig struct {
	URL      string
	APIKey   string
	Model    string
	Provider string
}
