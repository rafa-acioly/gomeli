package announcements

type status string

// Status represents the status of an announcement
const (
	CLOSED  status = "closed"
	PAUSED  status = "paused"
	ACTIVE  status = "active"
	ENABLED status = "enabled"
)
