package line

type LineUsers struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increase'"`
	Token string `json:""`
}