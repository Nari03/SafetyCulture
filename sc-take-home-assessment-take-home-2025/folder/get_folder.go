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
	// Your code here...
	folders_w_orgID := f.GetFoldersByOrgID(orgID)
	fmt.Printf("found folders w orgID")
	childFolders := []Folder{}

	parentPath := name + "."

	for _, folder := range folders_w_orgID{
		if strings.HasPrefix(folder.Paths, parentPath){
			childFolders = append(childFolders, folder)
		}
	}
	// need to add error cases for when the folder does not exist and when folder does not exist in specified organisation
	return childFolders
}
