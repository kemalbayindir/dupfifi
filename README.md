<div style="text-align:center">

![alt text](https://github.com/kemalbayindir/dupfifi/blob/main/profile.jpeg?raw=true)

</div>

# DUPLICATE FILE FINDER
Duplicate File Finder (DUPFIFI) searches duplicate files and exposes found similar files (by hash) as JSON.



## .env

Searchable types
> VALID_EXT=.png,.jpg,.jpeg,.bmp

Dont search these folders
> EXCLUDE_DIR=.git

Dont search these files
> EXCLUDE_FILE=.DS_Store

Base search folder, root folder for scan operations
> SCAN_PATH="/Users/kemalbayindir/WORKSPACE/"

## TODO
- Generate human readable html report
- Serve managament UI to arrange files or custom searchs
- Search specific folder and files in a target folder
- Commandline arguments for above functions
- Executable build, cross platform builds
