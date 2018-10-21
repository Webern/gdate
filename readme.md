gdate
=====

The purpose of gdate is to provide a simple yyyy-mm-dd date type in Golang which we can use to bypass all the complexities of DateTime datatypes (e.g. Golang time.Time).

* It parses from a proper yyyy-mm-dd string and writes to a yyyy-mm-dd string.
* It has no timezone, we just want to represent a calendar date.