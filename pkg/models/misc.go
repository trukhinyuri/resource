package models

import (
	rstypes "git.containerum.net/ch/json-types/resource-service"
	"github.com/sirupsen/logrus"
)

// PermCheck checks permissions needed for operation
// Owner can do all actions
func PermCheck(perm, needed rstypes.PermissionStatus) bool {
	switch perm {
	case rstypes.PermissionStatusRead:
		if needed == rstypes.PermissionStatusReadDelete {
			return false
		}
		fallthrough
	case rstypes.PermissionStatusReadDelete:
		if needed == rstypes.PermissionStatusWrite {
			return false
		}
		fallthrough
	case rstypes.PermissionStatusWrite:
		if needed == rstypes.PermissionStatusOwner {
			return false
		}
		fallthrough
	case rstypes.PermissionStatusOwner:
		return true
	}
	logrus.Errorf("unreachable code in PermCheck")
	return false
}