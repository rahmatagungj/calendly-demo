package core

import (
	"time"
)

type Invitee struct {
	Email    string
	Name     string
	Timezone *time.Location
}
