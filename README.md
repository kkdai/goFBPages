goFBPages
======================
A facebook page photo album tool that supports concurrency download.

Install
--------------

    go get -u -x github.com/kkdai/goFBPages


Usage
---------------------

    goFBPages [options] 

When [path...] is not given, "./" will be used as the default path for scanning files.


Options
---------------

- `-n` Facebook page name such as: [scottiepippen](https://www.facebook.com/scottiepippen).
- `-id` Facebook page id, if you know such as: 112743018776863 
- `-c=2` number of workers. (concurrency)


Examples
---------------

Download all photos from Scottie Pippen facebook pages

  goFBPages -n scottiepippen
