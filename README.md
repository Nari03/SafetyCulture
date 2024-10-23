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

The testing for this function has been done using following the testing framework that was already provided, however, I recognise that this structure does not allow for checking for specific errors, and tests in a more general manner.

## Component 2:
Implemented MoveFolder function in move_folder.go, this function accepts 2 arguments - name of the source folder and name of the destination folder and returns the new structure with all folders (after moving), and any errors that occured
I am also checking for some edge cases and returning appropriate errors. The edge cases covered are as follows:
-   Source folder does not exist
-   Destination folder does not exist
-   Source folder and destination folder don't have the same organisation id
-   Destination folder is a child of the source folder - cannot move a folder to one of its subfolders
-   Source folder and destination folder are the same - cannot move the folder to itself

Testing for this function has been implemented differently. I have created test functions for testing different functionalities and checking if the function is returning the right errors where applicable.