PangolinDB  Design & Interface Considerations
=============================================

> This repo follows along with a blog series I'm writing that you can follow [here](https://dylanlott.com/writing-a-database-in-go-from-scratch/)

# BIG TODO: These interfaces need types. Stop using interface{} everywhere. 

# Key Value Store

The heart of this database is going to be the key value store that everything interacts with.

Key value stores are great for a couple reasons, but most importantly they're fast, they scale reasonably well, and they impose the bare minimum amount of structure on our data that's reasonably possible.

This makes them reasonably suited to being the underlying store for our database.

Our interface for our key value store should be fairly simple, and abstract away all of the complexity of our actual key value store, which we'll be handling elsewhere.

```
type KVStore interface {
    Get(key string) (interface{}, bool, error)
    Put(key string, value interface{}) (error)
    Delete(key string) (error)
}
```

This interface allows us to plug and play any driver we want behind our database, so we can plug in a different driver later on with minimal effort if our current solution has problems or if we want better performance.

For right now, we're going to try writing our own kv store just for the challenge of it.


# Indexes
Indexes are what will make our database really fast. We can accomplish O(log n) complexity inserts, updates, lookups, and deletes with indexes, which makes them perfectly suited for handling our indexes.

Our indexes will be built with AVL trees, which are auto balancing and very well suited for this type of data manipulation. We are going to use the same interface pattern we used above, but we're going to have two different types of indexes that can fulfill our Index interface.

First, let's outline our index interface.

```
type Index interface {
    Get(key string) (interface{}, error)
    Put(key string, value interface{}) error
    Remove(key string) error
}
```

# Collections

Our indexes will be tied quite closely to our Collections, so we'll cover that next.

```
type Collection interface {
    Open(name string) (*Driver, error)
    Close() error
    AddIndex(field string) error
    RemoveIndex(field string) error
}
```

Our *Driver pointer will be what we do our CRUD operations on, but the other functions are specifically for handling collection level tasks.

The Driver struct will bridge the gap between our Collections and our KVStore. The Driver will fulfill the interface for the KVStore package.

# Queries

```
type Field struct {
    Key     string
    Join    string
}

type Population struct {
    Paths []*Field  
}

type QueryOpts struct {
    Populate    *Population
    Limit       int64
    Skip        int64
}

type Query func(coll *Collection, query interface{}, opts *QueryOpts) (interface{}, error)
```

We only have one Query function, and it will take a collection to operate on, a `query` parameter which can be anything that is JSON compatible, and then a QueryOpts struct for controlling details about the query.

Example Field Struct:  // this probably needs improvement
```
&Field{
    Collection: "User",
    Key: Where{
        Field: "books",
        Value: "Id",
        Match: "books.Id",
    },
    Join: Where{
        Collection: "Books",
        Field: "Id",
    }
}
```
