package viewmodels

import (
	"app/passporte/models"
)
type Group struct {
	Groups       []models.Group
	GroupName    string
	GroupMembers string
	GroupKey     []string

}
