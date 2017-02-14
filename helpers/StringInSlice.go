/*Created by Aswathy A*/
package helpers

/*Function for checking whether the string is present in a slice.
Return true if it is present and false if not*/
func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
