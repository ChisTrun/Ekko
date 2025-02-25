package checker

import (
	"context"
	"fmt"
)

func CheckRole(ctx context.Context, roleId int32, roleIds []string) error {
	if len(roleIds) != 1 {
		return fmt.Errorf("invalid role id")
	}

	if roleIds[0] != fmt.Sprintf("%v", roleId) {
		return fmt.Errorf("permission denied")
	}

	return nil
}
