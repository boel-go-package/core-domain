package config

type AuditTrailModel struct {
	Version  int    `bun:"version,notnull,default:0" json:"version"`
	CreateBy string `bun:"create_by,notnull" json:"create_by"`
	CreateAt string `bun:"create_at,nullzero,notnull,default:current_timestamp" json:"create_at"`
	UpdateBy string `bun:"update_by" json:"update_by"`
	UpdateAt string `bun:"update_at,nullzero,default:current_timestamp" json:"update_at"`
	// DeleteBy string    `bun:"delete_by" json:"delete_by"`
	// DelateAt time.Time `bun:"delete_at,nullzero,default:current_timestamp" json:"delete_at"`
	// Status     int       `bun:"status,default:1" json:"status"`
}
