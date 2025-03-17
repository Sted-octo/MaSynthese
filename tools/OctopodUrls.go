package tools

import "os"

func OctopodDomainGetter() string {
	return os.Getenv("OCTOPOD_DOMAIN")
}

func OctopodUrlApiGetter() string {
	return OctopodDomainGetter() + "/api/v0"
}
