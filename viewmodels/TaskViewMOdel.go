package viewmodels
import (
	"app/passporte/models"
)
type TaskViewModel  struct {
	Task		[]models.Task
	Key		[]string
	ProjectName	[]string

}