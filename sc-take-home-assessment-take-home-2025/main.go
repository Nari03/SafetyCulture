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
	childFolder:= folderDriver.GetAllChildFolders(orgID, "noble-vixen")
	// folder.PrettyPrint(res)
	// fmt.Printf("\n Folders for orgID: %s", orgID)
	// folder.PrettyPrint(orgFolder)

	fmt.Printf("\n CHild folders of noble-vixen in orgid: %s ", orgID)
	folder.PrettyPrint(childFolder)
	
}
