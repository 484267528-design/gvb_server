package desens

import "strings"

func DesensitizationTel(tel string) string {
	//15852526885
	//158****6885
	if len(tel) != 11 {
		return ""
	}
	return tel[0:3] + "****" + tel[7:]
}
func DesensitizationEmail(email string) string {
	eList := strings.Split(email, "@")
	if len(eList) != 2 {
		return ""
	}
	return eList[0][:1] + "****@" + eList[1]
}
