package viewmodels
type DashBoardViewModel struct {
	CompletedTask 		float32
	PendingTask		float32
	AcceptedUsers		float32
	RejectedUsers		float32
	PendingUsers		float32
	JobNameArray		[]string
	Key			[]string
	JobArrayLength		int
	JobCustomerNameArray 	[]string
	TaskDetailArray		[][]string
	CompanyTeamName         string
	CompanyPlan		string
	AdminFirstName		string
	AdminLastName		string
	ProfilePicture		string

}