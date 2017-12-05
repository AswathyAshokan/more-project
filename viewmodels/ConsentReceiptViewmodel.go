
package viewmodels
type ConsentReceipt struct {
GroupKey               []string
GroupMembers        []string
CompanyTeamName        string
CompanyPlan        string
AdminFirstName        string
AdminLastName        string
ProfilePicture        string
NotificationArray    [][]string
NotificationNumber       int
}

type LoadConsent struct {
Values                  [][]string
Keys            []string
InnerContent            []ConsentStruct
CompanyTeamName         string
CompanyPlan        string
AdminFirstName        string
AdminLastName        string
ProfilePicture        string
NotificationArray    [][]string
NotificationNumber       int

}
type ConsentStruct struct {
Description           string
AcceptedUsers         []string
RejectedUsers         []string
PendingUsers          []string
InstructionKey        string
AcceptedDate           string


}
type EditConsentReceipt struct {
GroupKey                   []string
GroupMembers            []string
ReceiptName                string
ConsentKey               []string
ConsentMembers            []string
CompanyTeamName            string
CompanyPlan            string
AdminFirstName            string
AdminLastName            string
ProfilePicture            string
PageType                string
UserNameToEdit                 []string
InstructionArrayToEdit      []string
ConsentId               string
UsersForEdit             []string
SelectedUsersKey         []string
NotificationArray        [][]string
NotificationNumber           int
}