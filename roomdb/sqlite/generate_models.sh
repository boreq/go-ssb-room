#!/bin/sh

set -e

# ensure tools are installed
go get github.com/volatiletech/sqlboiler/v4
go get github.com/volatiletech/sqlboiler-sqlite3

# run the migrations (creates testrun/TestSchema/roomdb)
go test -run Schema

# make sure the sqlite file was created
test -f testrun/TestSchema/roomdb || {
    echo 'roomdb file missing'
    exit 1
}

# generate the models package
sqlboiler sqlite3 --wipe --no-tests

echo "all done. models updated!"