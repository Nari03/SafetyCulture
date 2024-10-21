package folder_test

import (
	"testing"
	"github.com/georgechieng-sc/interns-2022/folder"

	"github.com/gofrs/uuid"

)
// Tests to make
// 1. for null input - null orgId, null name, and both null
// 2. 
func Test_folder_MoveFolder(t *testing.T) {
	org1 := uuid.Must(uuid.NewV4())

	testFolders := []folder.Folder{
		{Name: "alpha", OrgId: org1, Paths: "alpha"},
		{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"},
		{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},
		{Name: "delta", OrgId: org1, Paths: "alpha.delta"},

	}
	f:=folder.NewDriver(testFolders)

	// trying to move bravo to delta
	updatedFolders, err := f.MoveFolder("bravo", "delta")
	if err!= nil {
		t.Errorf("Unexpected error %v", err)
	}
	expectedPath := "alpha.delta.bravo"
	works := false
	for _, folder := range updatedFolders {
		if folder.Name == "bravo" && folder.Paths == expectedPath {
			works = true
			break
		}
	}
	if !works {
		t.Errorf("Expected folder to have path %s, but failed", expectedPath)
	}
}


