package datastruct

type ObjectDetail struct {
	Id          string `db:"uuid" json:"id"`
	Name        string `db:"name" json:"name"`
	Size        int    `db:"size" json:"size"`
	ContentType string `db:"content_type" json:"content_type"`
	Path        string `db:"path" json:"path"`
	Url         string `db:"url" json:"url"`
}
