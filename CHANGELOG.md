## v0.5.2
* **Fix** : The previous version changes broke the functionality of *PackFolder* with absolute paths. Now it works as expected.
---

## v0.5.1
* **Fix** : *PackFolder()* stored absolute paths. Now only the target folder is stored.
---

## v0.5.0
Important: this version has api breaking changes.  
Almost every functions were modified to allow apply several processes over a file without write on disc until the last process.
Two helper functions are now public to allow us to write and read files when necessary.  
### New features
* **ReadFile(fileName string) ([]byte, error)** : reads a file from disc and returns an array of bytes. In this version a lot of functions use a []byte as input param.
* **WriteFile(filePath string, data []byte, perm os.FileMode) error** : writes a file. It will create a directory if necessary.  
### Changes (new functions signatures)
* **Encrypt(input []byte, password string) (output []byte, e error)** : encrypts *input* into *output*.
* **Decrypt(input []byte, password string) (output []byte, e error)** : decrypts *input* into *output*.
* **Compress(input []byte, fileName string) ([]byte, error)** : compress *input* into *output*. *fileName* will be stored in gzip file header.
* **Decompress(input []byte) (output []byte, fileName string, e error)** : decompress *input* file into *output*.
* **PackFolder(input string) ([]byte, error)** : packs *input* file into []byte.
* **Unpack(input []byte, output string) error** : unpacks *input* into *output* path.  
---  

## v0.4.0
Added some features to verify files.
### New features
* **GetCRC32(fileName string) (uint32, error)**  : returns the CRC32 checksum of any file using IEEE polynomial.
* **GetGzipCRC32(fileName string) (uint32, error)** :  returns the CRC32 of decompressed file stored in the gzip file *fileName*.
* **GetDecompressedSize(fileName string) (uint32, error)** : returns the size of decompressed file stored in the gzip file *fileName*.  
---  

## v0.3.0
Initial funcionalities was implemented.
### Features
* **Encryption**: uses AES256 CGM algorithm. Password is converted into a 256-bit key using SHA256 hash algorithm so it does not requires 32 bytes length.
    * **Encrypt(input string, output string, password string) error** : encrypts *input* file into *output* path.
    * **Decrypt(input string, output string, password string) error** : decrypts *input* file into *output* path.
* **Compression**: uses gzip package (good performance instead of great compression ratio).
    * **Compress(input string, output string) error** : compress *input* file into *output* file. Adds ".gz" extension to output file.
    * **Decompress(input string, output string) error** : decompress *input* file into *output* path. If *output* == "" then uses current directory.
* **Packaging**: uses tar package.
    * **PackFolder(input string, output string)** : packs *input* file into *output* file. Adds ".tar" extension to output file.
    * **Unpack(input string, output string)** : unpacks *input* file into *output* path. If *output* == "" then uses current directory.
