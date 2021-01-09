<div style="text-align:center">

![alt text](https://github.com/kemalbayindir/dupfifi/blob/main/profile.jpeg?raw=true)

</div>

# **Dup**licate **Fi**le **Fi**nder
Duplicate File Finder (DUPFIFI) searches duplicate files and exposes found similar files (by hash) as JSON.



## Parameters

-excludedDirs string
>    Not allowed directories to scan process. Please use comma to seperate extensions. (default ".git,node_modules")

-excludedFiles string
>    Not allowed files to scan process. Please use comma to seperate extensions. (default ".DS_Store")

-includedExtensions string
>    Allowed dile extension(s) during scan process. Please use comma to seperate extensions. (default ".png,.jpg,.jpeg,.bmp")

-scanPath string
>    Target path to scan process (default "./")

# SEARCH RESULT
Sample JSON result
```json
{
    "7c0e589fe406cddb3d094cdc35691cb1": [
        {
            "Path": "/Users/kemalbayindir/WORKSPACE/GOCOMPARE/test-area/f1/f1.1/profile.png"
        }
    ],
    "a6701c65a3dbc77edb70f27b52952c7b": [
        {
            "Path": "/Users/kemalbayindir/WORKSPACE/GOCOMPARE/test-area/f2/f2.2/heap.png"
        },
        {
            "Path": "/Users/kemalbayindir/WORKSPACE/GOCOMPARE/test-area/f2/f2.3/heap.png"
        },
        {
            "Path": "/Users/kemalbayindir/WORKSPACE/GOCOMPARE/test-area/f3/f3.2/heap copy.png"
        }
    ]
}
```

## TODO
- [ ] Generate human readable html report
- [ ] Serve managament UI to arrange files or custom searchs
- [x] Customize search process by custom filters (exclude/include dir/extension)
- [x] Commandline argument support
- [ ] Executable build, cross platform builds
- [ ] Use go routines to speedup
