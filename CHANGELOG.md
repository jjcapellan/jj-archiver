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
    * **Unpack(input string, output string)** : umpacks *input* file into *output* path. If *output* == "" then uses current directory.
