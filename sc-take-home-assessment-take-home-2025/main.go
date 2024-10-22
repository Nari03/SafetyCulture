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
	
	childFolder:= folderDriver.GetAllChildFolders(orgID, "noble-vixen")
	

	fmt.Printf("\n Child folders of noble-vixen in orgid: %s ", orgID)
	folder.PrettyPrint(childFolder)

	fmt.Printf("\n Moving folder quick-cyber to national-screwball %s ", orgID)
	movingFolders, err := folderDriver.MoveFolder("quick-cyber", "national-screwball")
	if err != nil {
		fmt.Printf("Error occured in moving folder: %v\n", err)
		return
	}
	folder.PrettyPrint(movingFolders)
	
}
