# Simple copy scripts 
## Description

Simple copy scripts in Python and Golang which copy recursively files form source folder to existing destination folder
Both scripts works on Windows, Linux and Mac OS

## Table of Contents
  
* [Prerequisites](#prerequisites)

* [Usage](#usage)

## Prerequisites

To use those scripts you need installed [Python 3](https://www.python.org/downloads/) for python script or/and [Go Lang](https://go.dev/doc/install)

## Usage

When you have installed Python and/or Glang:

Clone this repo to local machine:

```md
git clone https://github.com/HubGab-Git/copy_script.git
```

If you would like run Python script Enter python folder:

```md
cd copy_script/python
```

run example:
```md
python copy.py -s . -d ../
```

If you would like run Golang script Enter python folder:

```md
cd copy_script/golang
```

run example:
```md
go run main.go -s . -d ../
```

Above examples copy all files from current directory to parent directory

help message:
```md
usage: copy_script [-h] [-s SOURCE_DIR] [-d DESTINATION_DIR]

Copy files form one folder to another

options:
  -h, --help            show this help message and exit
  -s SOURCE_DIR, --source-dir SOURCE_DIR
                        Source directory from where files should be copied
  -d DESTINATION_DIR, --destination-dir DESTINATION_DIR
                        Source directory from where files should be copied
```
