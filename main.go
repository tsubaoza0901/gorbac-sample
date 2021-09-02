package main

import (
	"fmt"

	"github.com/mikespook/gorbac"
)

func main() {
	// Role Names
	mediaManager := "MediaManager"
	systemMaintenanceManager := "SystemMaintenanceManager"
	systemAdmin := "SystemAdmin"

	// Get some new permissions
	pA := gorbac.NewStdPermission("permission-a")
	pB := gorbac.NewStdPermission("permission-b")
	pC := gorbac.NewStdPermission("permission-c")

	type TestCase struct {
		RoleName    string
		Permissions []gorbac.Permission
		Parents     []string
	}

	type TestCases []TestCase

	testCases := TestCases{
		{
			RoleName:    mediaManager,
			Permissions: []gorbac.Permission{pA},
			Parents:     nil,
		},
		{
			RoleName:    systemMaintenanceManager,
			Permissions: []gorbac.Permission{pB},
			Parents:     nil,
		},
		{
			RoleName:    systemAdmin,
			Permissions: []gorbac.Permission{pC},
			Parents:     []string{mediaManager, systemMaintenanceManager},
		},
	}

	rbac := gorbac.New()

	for _, testCase := range testCases {
		// Get some new roles
		role := gorbac.NewStdRole(testCase.RoleName)

		// Add the permission to roles
		for _, v := range testCase.Permissions {
			role.Assign(v)
		}

		// Add the roles to the RBAC instance
		rbac.Add(role)

		// Set the inheritance
		rbac.SetParents(testCase.RoleName, testCase.Parents)
	}

	if rbac.IsGranted(systemAdmin, pA, nil) && rbac.IsGranted(systemAdmin, pB, nil) && rbac.IsGranted(systemAdmin, pC, nil) {
		fmt.Printf("The %s has been granted permission-a, b and c.\n", systemAdmin)
	} else {
		fmt.Printf("The %s has not been granted permission-a, b or c.\n", systemAdmin)
	}
}
