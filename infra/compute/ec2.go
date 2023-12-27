package compute

import "fmt"

func GetResource(resource_type string) string {
	rslt := fmt.Sprintf("The resource type is %v", resource_type)
	return rslt
}
