package role

import (
	"testing"
)

func Test_role_valueRange(t *testing.T) { // MEMO: 範囲外の値がないことを確認
	for roleName, roleNum := range NameList {
		if roleNum < 0 || roleNum >= maxRole {
			t.Errorf("Role '%s' has an out-of-range value: %d", roleName, roleNum)
		}
	}
}

func Test_role_allRolesMapped(t *testing.T) { // MEMO: NameList の完全性を確認
	mappedRoles := make(map[int]bool)
	for _, roleNum := range NameList {
		mappedRoles[roleNum] = true
	}

	for i := 0; i < maxRole; i++ {
		if !mappedRoles[i] {
			t.Errorf("Role with value %d is not mapped in NameList", i)
		}
	}
}

func Test_role_nameListValidity(t *testing.T) { // MEMO: NameList に無効なエントリがないことを確認
	for roleName, roleNum := range NameList {
		if roleNum < 0 || roleNum >= maxRole {
			t.Errorf("NameList contains invalid mapping: %s -> %d", roleName, roleNum)
		}
	}
}
