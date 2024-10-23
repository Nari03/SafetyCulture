# SafetyCulture
Hi, this is my solution for the Safety Culture take home assessment.

## Component 1:
Implemented GetAllChildFolders function in get_folder.go that returns all child folders of a given folder. I have taken care of certain edge cases when implementing this, such as
-   No organisationId provided
-   No foler name provided
-   No folders with the given Organisation id found
-   If folder does not exist at all (this could be a potential security breach)
-   If the folder is not found in the given organisation 
-   The folder does not contain any child folders

## Component 2:
