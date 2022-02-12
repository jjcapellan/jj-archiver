![GitHub tag (latest by date)](https://img.shields.io/github/tag-date/jjcapellan/jj-archiver.svg)
![GitHub license](https://img.shields.io/github/license/jjcapellan/jj-archiver.svg)  
# JJ-ARCHIVER
A simple convenience utility library to **pack** (tar), **compress** (gzip) and **encrypt** (AES256 GCM) files.

## Usage
```golang
// ...
import . "github.com/jjcapellan/jj-archiver"
// ...

// Pack folder "user/projects" into file "./projects.tar"
err := PackFolder("user/projects", "projects")
if err!= nil{
 //   ...
}

// Compress file "projects.tar" into "user/compressed/projects.tar.gz"
err := Compress("projects.tar", "user/compressed")
if err!= nil{
 //   ...
}

// Encrypt file "user/compressed/projects.tar.gz" into "./projects.tar.gz.crp"
err := Encrypt("user/compressed/projects.tar.gz", "", "mypassword")
if err!= nil{
 //   ...
}

// Decrypt file "./projects.tar.gz.crp" into "./projects.tar.gz"
err := Decrypt("projects.tar.gz.crp","", "mypassword")
if err!= nil{
 //   ...
}

// Decompress file "./projects.tar.gz" into "./projects.tar"
err := Decompress("projects.tar.gz", "")
if err!= nil{
 //   ...
}

// Unpack file "./projects.tar" in folder "user/copyfolder"
err := Unpack("projects.tar", "user/copyfolder")
if err!= nil{
 //   ...
}

```

## Dependencies
This library is built over standard golang libraries, so it hasn't external dependencies.