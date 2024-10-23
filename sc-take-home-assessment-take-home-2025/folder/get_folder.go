package folder

import "github.com/gofrs/uuid"
import (
	"fmt"
	"strings"
)
func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {

	// checking if the given input is empty - if yes, handling it gracefully
	if orgID == uuid.Nil || name == "" {
		fmt.Printf("Empty orgId or name provided. Please try again and provide valid input")
		return nil
	}

	foldersWithOrgID := f.GetFoldersByOrgID(orgID)
	//if no folders were reuturned, providing an error message:
	if len(	foldersWithOrgID) == 0 {
		fmt.Printf("No folders were found having the organisation id: %s\n Please check if you are entering the right organisation Id", orgID)
		return nil
	}

	// checking if the folder exists at all:
	folderExists := false
	allFolders := f.folders
	
	fmt.Printf("%v",allFolders)

	for _,folder := range allFolders{
		if folder.Name == name{
			folderExists = true
		}
	}
	if folderExists == false{
		fmt.Printf("The folder %s does not exist ", name)
		return nil
	}

	// now checking if the folder exists in the given organisation
	folderExistsInOrg := false

	for _,folder := range foldersWithOrgID{
		if folder.Name == name{
			folderExistsInOrg = true
		}
	}
	if folderExistsInOrg == false{
		fmt.Printf("The folder %s not found in organisationId %s ", name, orgID)
		return nil
	}

	childFolders := []Folder{}
	parentPath := name + "."

	// positive case - getting all child folders
	for _, folder := range foldersWithOrgID{
		if strings.HasPrefix(folder.Paths, parentPath){
			childFolders = append(childFolders, folder)
		}
	}

	// if no child folders are found
	if len(childFolders) == 0{
		fmt.Printf("Folder %s found in OrganisationID %s but is empty.", name, orgID)
	}
	return childFolders
}
