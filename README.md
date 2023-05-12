# Ranger is a commandline tool to organize your files.


## How to use

most importnat thing first:
use ```--help``` after any command to get all options and flags.

The following will organize your directory up to the level of months.
You can use ```years``` and ```days``` respectively
```bash
$ ranger years
```
if you had a directory that looks like this:

```bash
├── myDirectory/
    └── 20141106_092204.jpg
    └── 20141115_124714.jpg
    └── 20141128_152219.jpg
    └── 20141128_152227.jpg
    └── 20150503_104114.jpg
    └── 20150503_111314.jpg
    └── 20150512_131524.jpg
    └── 20150622_204250.jpg
    └── 20150630_170123.jpg
    └── 20150714_194224-1.jpg
```
it will turn into this: *(more languages will potentially be added)*
```bash
├── myDirectory/
    ├───2014
    │   └── 11-November
    │       └── 20141106_092204.jpg
    │       └── 20141115_124714.jpg
    │       └── 20141128_152219.jpg
    │       └── 20141128_152227.jpg
    └───2015
        ├── 05-Mai
        │   └── 20150503_104114.jpg
        │   └── 20150503_111314.jpg
        │   └── 20150512_131524.jpg
        ├── 06-Juni
        │   └── 20150622_204250.jpg
        │   └── 20150630_170123.jpg
        └── 07-Juli
            └── 20150714_194224-1.jpg
```
the flags ```--modtime``` and ```--exif``` will use the last modification time or exif data to retrieve the time. By default ranger searches the filename for time information (```--filename```)

Tipp: you can use the different sources in succession for the best results if you have different types of files
```bash
$ ranger days
$ ranger days --modtime
```


if you made a mistake or have a lot of nested directories you can use
```bash
$ ranger compile
```
this will scan all subdirectories and move all files to the directory from where you ran the command. It will also delete all the folders (they are empty anyway now, no worrys)

## Install

```bash
$ git clone https://github.com/theryann/ranger.git
```

in the root level of the repo run the following (you should have set your gopath)

```bash
$ go install .
```


