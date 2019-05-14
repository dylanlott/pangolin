package indexes

//generate a `IntHashtable` hash table of `int` values
//go:generate genny -in hashtable.go -out hashtable-int.go gen "Value=int"

//generate a `StringHashtable` hash table of `string` values
//go:generate genny -in hashtable.go -out hashtable-string.go gen "Value=string"
