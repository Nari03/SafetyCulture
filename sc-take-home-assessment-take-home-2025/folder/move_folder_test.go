package folder_test

import (
	"testing"
	"github.com/georgechieng-sc/interns-2022/folder"

	"github.com/gofrs/uuid"

)

func Test_folder_MoveFolder_general(t *testing.T) {
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

func Test_MoveFolder_NullInput(t *testing.T) {
	orgID := uuid.Must(uuid.NewV4())

	testFolders := []folder.Folder{
		{Name: "alpha", OrgId: orgID, Paths: "alpha"},
	}
	f := folder.NewDriver(testFolders)

	// Test: Empty folder name
	_, err := f.MoveFolder("", "alpha")
	if err == nil || err.Error() != "Source folder  does not exist" {
		t.Errorf("Expected error 'Source folder  does not exist', but got %v", err)
	}

	// Test: Empty destination folder name
	_, err = f.MoveFolder("alpha", "")
	if err == nil || err.Error() != "Destination folder  does not exist" {
		t.Errorf("Expected error 'Destination folder does not exist', but got %v", err)
	}
}

func Test_MoveFolder_CircularMove(t *testing.T) {
	orgID := uuid.Must(uuid.NewV4())

	testFolders := []folder.Folder{
		{Name: "alpha", OrgId: orgID, Paths: "alpha"},
		{Name: "bravo", OrgId: orgID, Paths: "alpha.bravo"},
	}

	f := folder.NewDriver(testFolders)

	_, err := f.MoveFolder("alpha", "bravo")
	if err == nil || err.Error() != "Cannot move folder to one of its sub folders" {
		t.Errorf("Expected circular move error, but got: %v", err)
	}
}

func Test_MoveFolder_DifferentOrgID(t *testing.T) {
	orgID1 := uuid.Must(uuid.NewV4())
	orgID2 := uuid.Must(uuid.NewV4())

	testFolders := []folder.Folder{
		{Name: "alpha", OrgId: orgID1, Paths: "alpha"},
		{Name: "bravo", OrgId: orgID2, Paths: "bravo"},
	}

	f := folder.NewDriver(testFolders)

	_, err := f.MoveFolder("alpha", "bravo")
	if err == nil || err.Error() != "Cannot move folder since source and destination have different orgIds" {
		t.Errorf("Expected different orgID error, but got: %v", err)
	}
}

func Test_MoveFolder_EmptyInput(t *testing.T) {
	orgID := uuid.Must(uuid.NewV4())

	testFolders := []folder.Folder{
		{Name: "alpha", OrgId: orgID, Paths: "alpha"},
	}

	f := folder.NewDriver(testFolders)

	// Empty source folder name
	_, err := f.MoveFolder("", "alpha")
	if err == nil || err.Error() != "Source folder  does not exist" {
		t.Errorf("Expected 'Source folder  does not exist', but got: %v", err)
	}

	// Empty destination folder name
	_, err = f.MoveFolder("alpha", "")
	if err == nil || err.Error() != "Destination folder  does not exist" {
		t.Errorf("Expected 'Destination folder  does not exist', but got: %v", err)
	}
}