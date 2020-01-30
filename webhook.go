package aftership

type Webhook struct {
	EventID            string `json:"event_id"`
	Event              string
	IsTrackingFirstTag bool     `json:"is_tracking_first_tag"`
	Message            Tracking `json:"msg"`
	Timestamp          int64    `json:"ts"`
}
