package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("1234"))
}
func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$YeF/A5WzUWHB8mbXdUTk7uSilpLCyuVDuqoOws7SaA3lhMgno0Fwu", "1234"))
}
