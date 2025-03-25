package actions

type ReservationEmailData struct {
	Name     string
	Location string
	Time     string
	Category string
	Details  string
}

type EmailRequest struct {
	Type  string               `json:"type"`
	Data  ReservationEmailData `json:"data"`
	Email string               `json:"email"`
}
