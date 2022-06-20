![GitHub tag (latest by date)](https://img.shields.io/github/tag-date/jjcapellan/jj-archiver.svg)
![GitHub license](https://img.shields.io/github/license/jjcapellan/jj-archiver.svg)  
# JJ-ARCHIVER
A simple convenience utility library to **pack** (tar), **compress** (gzip) and **encrypt** (AES256 GCM) files.

## Usage
```golang
// ...
import . "github.com/jjcapellan/jj-archiver"
// ...

// This code packages, compresses and encrypts a directory ("user/projects").

// 1. Directory Packaging into a tar []byte (packedData) using PackFolder()
packedData, err := PackFolder("user/projects")
if err != nil{
    // ...
}

// 2. Compression process using Compress(). gzipData is a compressed []byte.
// The second param is the file name to store in the header name of the gzip file
gzipData, err := Compress(packedData, "projects.gz")
if err != nil{
    //...
}

// 3. Encryption process using Encrypt(). encryptedData is an encrypted []byte
encryptedData, err := Encrypt(gzipData, "mypassword")
if err != nil{
    //...
}

// 4. Write the result into a file named "backups/projects.crp" using WriteFile
err = WriteFile("backups/projects.crp", encryptedData)
if err != nil{
    //...
}

```
Api docs available here: https://pkg.go.dev/github.com/jjcapellan/jj-archiver

## Dependencies
This library is built over standard golang libraries, so it hasn't external dependencies.