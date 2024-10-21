package folder

import (
	"fmt"
	"strings"
	// "errors"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
// 	// Your code here...

	var srcFolder *Folder
	var dstFolder *Folder

	// extracting the source folder and destination folder
	for i:= range f.folders {
		if f.folders[i].Name == name {
			// need to use a pointer here otherwise I will be working with a copy of the folder
			srcFolder = &f.folders[i]
		}
		if f.folders[i].Name == dst{
			dstFolder = &f.folders[i]
		}
	}

	// Handling the case when a source folder does not exist
	if srcFolder == nil{
		fmt.Printf("Source folder %s does not exist", name)
		return nil, fmt.Errorf("Source folder %s does not exist", name)
	}

	// Handling the case when a destination folder does not exist
	if dstFolder == nil{
		fmt.Printf("Destination folder %s does not exist", dst)
		return nil, fmt.Errorf("Destination folder %s does not exist", dst)
	}

	// Handling the case when the folders do not belong to the same orgId
	if srcFolder.OrgId != dstFolder.OrgId {
		fmt.Printf("Cannot move folder since source and destination have different orgIds")
		return nil, fmt.Errorf("Cannot move folder since source and destination have different orgIds")
	}

	// Handling the case of circular move (when the user tries to move folder to a child)
	if strings.HasPrefix(dstFolder.Paths, srcFolder.Paths+"."){
		fmt.Printf("Cannot move folder to one of its sub folders")
		return nil, fmt.Errorf("Cannot move folder to one of its sub folders")
	}

	// changing the paths of srcFolder and its subfolders
	oldPath := srcFolder.Paths
	newPath := dstFolder.Paths + "." + srcFolder.Name 

	for i:= range f.folders {
		if strings.HasPrefix(f.folders[i].Paths, oldPath){
			f.folders[i].Paths = strings.Replace(f.folders[i].Paths, oldPath, newPath, 1)
		}
	}

	return f.folders, nil
}
