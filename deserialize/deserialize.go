package deserialize

import (
	"encoding/json"
	"regex"
)

func Deserialize(res string, p string) {
	re := regexp.MustCompile(p)
	matches := re.FindAllStringSubmatch(res, -1)
	for _, match := range matches {
		fmt.Printf("%#v\r\n", strings.Join(match[1:], ","))
	}
}
