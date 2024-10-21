package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)
// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	org1 := uuid.Must(uuid.NewV4())
	org2 := uuid.Must(uuid.NewV4())

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
		// normal case - folder exists in org and has child folders, and nested child folders
			name: "alpha",
			orgID: org1,
			folders: []folder.Folder{
				{Name: "alpha", OrgId: org1, Paths: "alpha"},
				{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},
			},
			want: []folder.Folder{
				{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},
			},
		},
		{
			// No child folders
			name: "echo",
			orgID: org1,
			folders: []folder.Folder{
				{Name: "echo", OrgId: org1, Paths: "echo"},

			},
			want: []folder.Folder{
			},
		},
		{
			// Folder does not exist in org
			name: "echo",
			orgID: org1,
			folders: []folder.Folder{
				{Name: "alpha", OrgId: org1, Paths: "alpha"},

			},
			want: []folder.Folder{
			},
		},
		{
			// Folder exists, but wrong orgId
			name: "echo",
			orgID: org2,
			folders: []folder.Folder{
				{Name: "alpha", OrgId: org1, Paths: "alpha"},
				{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},

			},
			want: []folder.Folder{
			},
		},
		{
			// No folders found with given organisationId
			name: "echo",
			orgID: org2,
			folders: []folder.Folder{
				{Name: "alpha", OrgId: org1, Paths: "alpha"},
				{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},

			},
			want: []folder.Folder{
			},
		},
		{
			// No name provided
			name: "",
			orgID: org2,
			folders: []folder.Folder{
				{Name: "alpha", OrgId: org1, Paths: "alpha"},
				{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},

			},
			want: []folder.Folder{
			},
		},
		{
			// No organisation provided
			name: "alpha",
			orgID: uuid.Nil,
			folders: []folder.Folder{
				{Name: "alpha", OrgId: org1, Paths: "alpha"},
				{Name: "bravo", OrgId: org1, Paths: "alpha.bravo"},
				{Name: "charlie", OrgId: org1, Paths: "alpha.bravo.charlie"},

			},
			want: []folder.Folder{
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// t.Errorf("Printing all folders in this test: %v", tt.folders)
			f := folder.NewDriver(tt.folders)
			get := f.GetAllChildFolders(tt.orgID, tt.name)
			if len(get) != len(tt.want){
				t.Errorf("Got %v folders instead of %v", len(get), len(tt.want))
			}

			for i:= range(get){
				if(get[i].Name != tt.want[i].Name || get[i].Paths != tt.want[i].Paths || get[i].OrgId != tt.want[i].OrgId){
					t.Errorf("Got folder %v, expecting folder %v", get[i], tt.want[i])
				}
			}

		})
	}
}