package unlocktest

import (
	"fmt"
	"github.com/oneclickvirt/UnlockTests/utils"
	"github.com/oneclickvirt/UnlockTests/uts"
	"github.com/oneclickvirt/defaultset"
)

func MediaTest(language string) {
	uts.GetIpv4Info(false)
	uts.GetIpv6Info(false)
	readStatus := uts.ReadSelect(language, "0")
	if !readStatus {
		return
	}
	//if language == "zh" {
	//	fmt.Println("测试时间: ", defaultset.Yellow(time.Now().Format("2006-01-02 15:04:05")))
	//} else {
	//	fmt.Println("Test time: ", defaultset.Yellow(time.Now().Format("2006-01-02 15:04:05")))
	//}
	if uts.IPV4 {
		fmt.Println(defaultset.Blue("IPV4:"))
		uts.RunTests(utils.Ipv4HttpClient, "ipv4", language)
	}
	if uts.IPV6 {
		fmt.Println(defaultset.Blue("IPV6:"))
		uts.RunTests(utils.Ipv6HttpClient, "ipv6", language)
	}
}
