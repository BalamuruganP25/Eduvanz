package messages

type Register struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Date       string `json:"date"`
	Profession string `json:"profession"`
	Locality   string `json:"locality"`
	GuestCount int    `json:"guestcount"`
	Address    string `json:"address"`
}

type MeetUpList struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Date       string `json:"date"`
	Profession string `json:"profession"`
	Locality   string `json:"locality"`
	GuestCount int    `json:"guestcount"`
	Address    string `json:"address"`
}
