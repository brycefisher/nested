Demonstrates Soft Flattening Bug
================================

## Steps to reproduce:

 1. go get github.com/brycefisher/nested
 2. cd $GOPATH/src/github.com/brycefisher/nested
 3. glide up

## Expected Result

In this case, `glide up` _should_ notice that the github/brycefisher/nested and github.com/brycefisher/glide-nested-dependency require difference versions of the package github.com/aws/aws-sdk-go. The vendor ought to look something like this:

```
$GOPATH/src/github.com/brycefisher/nested/

 * glide.lock
 * glide.yaml
 * main.go
 * vendor/
   * github.com/
     * aws/
       * aws-sdk-go/
         * ...
     * brycefisher/
       * glide-nested-dependency
         * get_newer_s3_client.go
         * glide.lock
         * glide.yaml
         * vendor/
           * aws/
             * aws-sdk-go/
               * ...
```

## Actual Result

Output using glide 0.9.3 on OSX El Capitan, go version 1.6.0:

```
nesting $  glide up
[INFO] Downloading dependencies. Please wait...
[INFO] Fetching updates for github.com/aws/aws-sdk-go.
[INFO] Fetching updates for github.com/brycefisher/glide-nested-dependency.
[INFO] Setting version for github.com/aws/aws-sdk-go to v0.9.3rc4.
[INFO] Setting version for github.com/brycefisher/glide-nested-dependency to master.
[INFO] Resolving imports
[INFO] Fetching github.com/vaughan0/go-ini into /Users/bfisher-fleig/work/src/github.com/brycefisher/nesting/vendor
[ERROR] Error scanning github.com/aws/aws-sdk-go/aws/session: open /Users/bfisher-fleig/work/src/github.com/brycefisher/nesting/vendor/github.com/aws/aws-sdk-go/aws/session: no such file or directory
[INFO] Downloading dependencies. Please wait...
[INFO] Setting references for remaining imports
[INFO] Project relies on 3 dependencies.
An Error has occurred
nesting $ tree -L 5
nesting master 🍻  tree -L 5
.
├── README.md
├── glide.lock
├── glide.yaml
├── main.go
└── vendor
    └── github.com
        ├── aws
        │   └── aws-sdk-go
        │       ├── Gemfile
        │       ├── LICENSE.txt
        │       ├── Makefile
        │       ├── NOTICE.txt
        │       ├── README.md
        │       ├── apis
        │       ├── aws
        │       ├── awsmigrate
        │       ├── doc-src
        │       ├── internal
        │       ├── sdk.go
        │       └── service
        ├── brycefisher
        │   └── glide-nested-dependency
        │       ├── get_newer_s3_client.go
        │       └── glide.yaml
        └── vaughan0
            └── go-ini
                ├── LICENSE
                ├── README.md
                ├── ini.go
                ├── ini_linux_test.go
                ├── ini_test.go
                └── test.ini

14 directories, 18 files
```
