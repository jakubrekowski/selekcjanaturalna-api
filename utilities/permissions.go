package utilities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Permission struct {
	Name string
	Bit  int64
}

var SNPermissions []Permission

func LoadPermissions() ([]Permission, error) {
	content, err := ioutil.ReadFile("./config/permissions.json")
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile: %v", err)
	}

	var permissions []Permission
	err = json.Unmarshal(content, &permissions)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return permissions, nil
}

func MissingPermissions(neededPerms int64, userPermsInt int64) int64 {
	return neededPerms & ^userPermsInt
}

func GetPermissions(permissions []Permission, permissionInt int64) []Permission {
	var output []Permission

	for i := 0; i < len(permissions); i++ {
		if permissions[i].Bit&permissionInt != 0 {
			output = append(output, permissions[i])
		}
	}

	return output
}

func PermsInit() {
	SNPermissions, _ = LoadPermissions()
}
