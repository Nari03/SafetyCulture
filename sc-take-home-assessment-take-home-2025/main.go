package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)
	res := folder.GetAllFolders()

	// example usage
	folderDriver := folder.NewDriver(res)
	// orgFolder := folderDriver.GetFoldersByOrgID(orgID)
	// allFolders := folderDriver.GetAllFolders()
	
	childFolder:= folderDriver.GetAllChildFolders(orgID, "noble-vixen")
	
	// folder.PrettyPrint(res)
	// fmt.Printf("\n Folders for orgID: %s", orgID)
	// folder.PrettyPrint(orgFolder)

	fmt.Printf("\n Child folders of noble-vixen in orgid: %s ", orgID)
	fmt.Printf("\n All folders from main  ")

	folder.PrettyPrint(childFolder)
	
}
