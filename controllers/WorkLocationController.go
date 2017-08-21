package controllers
import (
	"app/passporte/models"
	"log"
	"reflect"
	"app/passporte/viewmodels"
	"app/passporte/helpers"
	"regexp"
	"time"
	"strings"
)

type WorkLocationcontroller struct {
	BaseController
}

func (c *WorkLocationcontroller) AddWorkLocaction() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	companyTeamName := c.Ctx.Input.Param(":companyTeamName")
	storedSession := ReadSession(w, r, companyTeamName)
	companyUsers :=models.Company{}
	var keySliceForGroupAndUser []string
	workLocationViewmodel := viewmodels.AddLocationViewModel{}
	userMap := make(map[string]models.WorkLocationUser)
	groupNameAndDetails := models.WorkLocationGroup{}
	groupMemberNameForTask :=models.GroupMemberNameInWorkLocation{}
	groupMemberMap := make(map[string]models.GroupMemberNameInWorkLocation)
	groupMap := make(map[string]models.WorkLocationGroup)
	var keySliceForGroup [] string
	var MemberNameArray [] string
	group := models.Group{}
	userName :=models.WorkLocationUser{}

	WorkLocation := models.WorkLocation{}
	if r.Method == "POST" {
		groupKeySliceForWorkLocationString := c.GetString("groupArrayElement")
		UserOrGroupNameArray :=c.GetStrings("userAndGroupName")
		taskLocation := c.GetString("taskLocation")
		groupKeySliceForWorkLocation :=strings.Split(groupKeySliceForWorkLocationString, ",")
		userIdArray := c.GetStrings("selectedUserNames")
		var groupKeySlice	[]string
		for j:=0;j<len(userIdArray);j++ {
			tempName := UserOrGroupNameArray[j]
			userOrGroupRegExp := regexp.MustCompile(`\((.*?)\)`)
			userOrGroupSelection := userOrGroupRegExp.FindStringSubmatch(tempName)
			if (userOrGroupSelection[1]) == "User" {
				log.Println("cp1")
				tempName = tempName[:len(tempName) - 7]
				userName.FullName = tempName
				userName.Status = helpers.StatusActive
				log.Println("tempId", userIdArray[j])
				userMap[userIdArray[j]] = userName
				log.Println("userMap", userMap)
				WorkLocation.Info.WorkLocation = taskLocation
				WorkLocation.Info.UsersAndGroupsInWorkLocation.User = userMap
				WorkLocation.Settings.DateOfCreation = time.Now().Unix()
				WorkLocation.Settings.Status = helpers.StatusActive
				log.Println("userMap[tempId]", userMap)

			}
			if groupKeySliceForWorkLocation[0] != "" {
				for i := 0; i < len(groupKeySliceForWorkLocation); i++ {
					groupDetails, dbStatus := group.GetGroupDetailsForWorkLocation(c.AppEngineCtx, groupKeySliceForWorkLocation[i])
					switch dbStatus {
					case true:
						groupNameAndDetails.GroupName = groupDetails.Info.GroupName
						memberData := reflect.ValueOf(groupDetails.Members)
						for _, key := range memberData.MapKeys() {
							keySliceForGroup = append(keySliceForGroup, key.String())
							log.Println("status",groupDetails.Members[key.String()].Status)
							if groupDetails.Members[key.String()].Status != helpers.UserStatusDeleted{
								MemberNameArray = append(MemberNameArray, groupDetails.Members[key.String()].MemberName)
								break
							}
						}
					case false:
						log.Println(helpers.ServerConnectionError)
					}
					for i := 0; i < len(keySliceForGroup); i++ {
						groupMemberNameForTask.MemberName = MemberNameArray[i]
						groupMemberMap[keySliceForGroup[i]] = groupMemberNameForTask
					}
					groupNameAndDetails.Members = groupMemberMap
					groupMap[string(groupKeySliceForWorkLocation[i])] = groupNameAndDetails
					groupKeySlice = append(groupKeySlice, string(groupKeySliceForWorkLocation[i]))
				}
			}
		}
		WorkLocation.Info.UsersAndGroupsInWorkLocation.User = userMap
		if groupKeySliceForWorkLocation[0] !="" {
			WorkLocation.Info.UsersAndGroupsInWorkLocation.Group = groupMap
		}
		dbStatus :=WorkLocation.AddWorkLocationToDb(c.AppEngineCtx)
		switch dbStatus {
		case true:
			w.Write([]byte("true"))
		case false:
			w.Write([]byte("false"))
		}
	}else {
		usersDetail :=models.Users{}
		dbStatus ,testUser:= companyUsers.GetUsersForDropdownFromCompany(c.AppEngineCtx,companyTeamName)

		switch dbStatus {
		case true:
			dataValue := reflect.ValueOf(testUser)
			for _, key := range dataValue.MapKeys() {

				dataValue := reflect.ValueOf(testUser[key.String()].Users)
				for _, userKeys := range dataValue.MapKeys() {
					//viewModel.GroupNameArray   = append(viewModel.GroupNameArray ,testUser[key.String()].Users[userKey.String()].FullName+" (User)")
					dbStatus := usersDetail.GetActiveUsersEmailForDropDown(c.AppEngineCtx, userKeys.String(), testUser[key.String()].Users[userKeys.String()].Email, companyTeamName)
					switch dbStatus {
					case true:
						workLocationViewmodel.GroupNameArray = append(workLocationViewmodel.GroupNameArray, testUser[key.String()].Users[userKeys.String()].FullName + " (User)")
						keySliceForGroupAndUser = append(keySliceForGroupAndUser, userKeys.String())
					case false:
						log.Println(helpers.ServerConnectionError)
					}
				}
			}
			allGroups, dbStatus := models.GetAllGroupDetails(c.AppEngineCtx,companyTeamName)
			log.Println("get all groups",allGroups)
			switch dbStatus {
			case true:
				dataValue := reflect.ValueOf(allGroups)
				for _, key := range dataValue.MapKeys() {
					if allGroups[key.String()].Settings.Status =="Active"{
						var memberSlice []string

						keySliceForGroupAndUser = append(keySliceForGroupAndUser, key.String())
						workLocationViewmodel.GroupNameArray = append(workLocationViewmodel.GroupNameArray, allGroups[key.String()].Info.GroupName+" (Group)")
						log.Println("activeeeeeeeeee",workLocationViewmodel.GroupNameArray)

						// For selecting members while selecting a group in dropdown
						memberSlice = append(memberSlice, key.String())
						groupDataValue := reflect.ValueOf(allGroups[key.String()].Members)
						for _, memberKey := range groupDataValue.MapKeys()  {
							memberSlice = append(memberSlice, memberKey.String())
						}
						workLocationViewmodel.GroupMembers = append(workLocationViewmodel.GroupMembers, memberSlice)

					}


				}
				workLocationViewmodel.UserAndGroupKey=keySliceForGroupAndUser
			case false:
				log.Println(helpers.ServerConnectionError)
			}
		case false:
			log.Println(helpers.ServerConnectionError)
		}

	}
	workLocationViewmodel.AdminFirstName = storedSession.AdminFirstName
	workLocationViewmodel.AdminLastName = storedSession.AdminLastName
	workLocationViewmodel.ProfilePicture =storedSession.ProfilePicture
	workLocationViewmodel.CompanyTeamName = companyTeamName
	/*c.Layout = "layout/layout.html"*/
	c.Data["vm"] = workLocationViewmodel
	c.TplName = "template/workLocation.html"
}
