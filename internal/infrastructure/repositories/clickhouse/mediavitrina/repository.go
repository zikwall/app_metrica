package mediavitrina

import clickhousebuffer "github.com/zikwall/clickhouse-buffer/v4"

type Repository struct {
	writer clickhousebuffer.Writer
}

func New(writer clickhousebuffer.Writer) *Repository {
	return &Repository{
		writer: writer,
	}
}
