# library-system

> library-system with cli(console) made with go

## Build Setup

``` bash

# setup db & ddl
just execute script on db/ddl.sql on root

# make binary
make console

# run binary
./bin/console_library <flag>

list of flag console:
1. add[#space]code_of_book(unique)[#space]name_of_book (add new book to inventory)
2. rent[#space]code_of_book (for update status of book rented)
3. rented (for display all rented books)
4. return[#space]code_of_book (for update status of book returned)
5. list (for display of list book and status of the book)
6. get[#space]code_of_book (show name of book by code)

```

For a detailed explanation on how things work, check out the [guide](https://golang.org/doc/) and [anymore](https://golang.org/).

