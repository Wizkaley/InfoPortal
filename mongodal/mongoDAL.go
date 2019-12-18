package mongodal

import (
	"gopkg.in/mgo.v2"
)

// MgoSessionDAL contains mongoDB collection utility
type MgoSessionDAL interface {
	DB(db string) MgoDBDAL
}

// MgoDBDAL contains mongo DB collection utility
type MgoDBDAL interface {
	C(collection string) MgoCollectionDAL
}

// MgoCollectionDAL contains mgo Collection Utilities
type MgoCollectionDAL interface {
	Insert(docs ...interface{}) error
	Find(query interface{}) MgoQueryDAL
	Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	Update(selectot interface{}, update interface{}) error
	UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	Remove(selector interface{}) error
	RemoveAll(selctor interface{}) (err error)
	Pipe(pipeline interface{}) MgoPipeDAL
}

// MgoQueryDAL contains MGO Query utilities
type MgoQueryDAL interface {
	All(result interface{}) error
	One(result interface{}) error
	Count() (int, error)
	Skip(n int) MgoQueryDAL
	Limit(n int) MgoQueryDAL
	Sort(fields ...string) MgoQueryDAL
	Apply(change mgo.Change, result interface{}) (*mgo.ChangeInfo, error)
}

// MgoPipeDAL contains mgo Pipe utilities
type MgoPipeDAL interface {
	All(result interface{}) error
	Iter() MgoIterDAL
	One(result interface{}) error
}

// MgoIterDAL contains iterator utilities . its a wrapper interface
type MgoIterDAL interface {
	All(result interface{}) error
	Close() error
	Err() error
	Next(result interface{}) bool
	Timeout() bool
}

// MongoSessionDAL is the actual implementation of the MgoSessionDAL interface
type MongoSessionDAL struct {
	session *mgo.Session
}

// MongoDAL is the actual implementation of the MgoDBDAL interface
type MongoDAL struct {
	db     *mgo.Database
	dbName string
}

// MongoCollectionDAL is the actual implementation of the MgoCollectionDAL
type MongoCollectionDAL struct {
	col     *mgo.Collection
	colName string
}

// MongoQueryDAL is the actual implementation of the MgoQueryDAL
type MongoQueryDAL struct {
	query *mgo.Query
}

// MongoPipeDAL is the actual implementation of MgoPipeDAL
type MongoPipeDAL struct {
	pipe *mgo.Pipe
}

// MongoIterDAL is the actual implementation of MgoIterDAL
type MongoIterDAL struct {
	iter *mgo.Iter
}

// MongoSessDAL is the funtion type
type MongoSessDAL func(sess *mgo.Session) MgoSessionDAL

// NewMongoSessDAL returns an Initialized object of MongoSessionDAL
func NewMongoSessDAL(sess *mgo.Session) MgoSessionDAL {
	s := &MongoSessionDAL{
		session: sess,
	}
	return s
}

// DB get DB from Mongo
func (s *MongoSessionDAL) DB(name string) MgoDBDAL {
	return s.DB(name)
}

// MongoDBDAL is the function type
type MongoDBDAL func(database *mgo.Database) MgoDBDAL

// NewMongoDBDAL returns an Initialized object of MongoDAL
func NewMongoDBDAL(database *mgo.Database) MgoDBDAL {
	mongo := &MongoDAL{
		db:     database,
		dbName: database.Name,
	}
	return mongo
}

// C get collection from mongo
func (m *MongoDAL) C(c string) MgoCollectionDAL {
	coll := &MongoCollectionDAL{
		col:     m.db.C(c),
		colName: c,
	}
	return coll
}

// Insert is my mongo Insert
func (c *MongoCollectionDAL) Insert(docs ...interface{}) error {
	return c.col.Insert(docs...)
}

// Remove is My Mongo Remove
func (c *MongoCollectionDAL) Remove(selector interface{}) error {
	return c.col.Remove(selector)
}

// RemoveAll is My Mongo RemoveAll
func (c *MongoCollectionDAL) RemoveAll(selector interface{}) error {
	return c.col.Remove(selector)
}

//Find mongo Find
func (c *MongoCollectionDAL) Find(query interface{}) MgoQueryDAL {
	q := &MongoQueryDAL{
		query: c.col.Find(query),
	}
	return q
}

// Update mongo Update
func (c *MongoCollectionDAL) Update(selector interface{}, update interface{}) error {
	return c.col.Update(selector, update)
}

// UpdateAll mongo UpdateAll
func (c *MongoCollectionDAL) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	return c.col.UpdateAll(selector, update)
}

// Upsert mongo Upsert
func (c *MongoCollectionDAL) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	return c.col.Upsert(selector, update)
}

// Pipe mongo Pipe
func (c *MongoCollectionDAL) Pipe(pipeline interface{}) MgoPipeDAL {
	pp := &MongoPipeDAL{
		pipe: c.col.Pipe(pipeline),
	}
	return pp
}

// All mongo All
func (q *MongoQueryDAL) All(result interface{}) error {
	return q.query.All(result)
}

//Count mongo countr
func (q *MongoQueryDAL) Count() (int, error) {
	return q.query.Count()
}

//One mongo one
func (q *MongoQueryDAL) One(result interface{}) error {
	return q.query.One(result)
}

//Skip mongo Skip
func (q *MongoQueryDAL) Skip(n int) MgoQueryDAL {
	q.query.Skip(n)
	return q
}

//Sort mongo Sort
func (q *MongoQueryDAL) Sort(fields ...string) MgoQueryDAL {
	q.query.Sort(fields...)
	return q
}

// Limit mongo limit
func (q *MongoQueryDAL) Limit(n int) MgoQueryDAL {
	q.query.Limit(n)
	return q
}

//Apply mongo Apply
func (q *MongoQueryDAL) Apply(change mgo.Change, result interface{}) (*mgo.ChangeInfo, error) {
	return q.query.Apply(change, result)
}

//All invokes All on iterator object
func (it *MongoIterDAL) All(result interface{}) error {
	return it.iter.All(result)
}

//All mongo All
func (p *MongoPipeDAL) All(result interface{}) error {
	return p.pipe.All(result)
}

//One mongo One
func (p *MongoPipeDAL) One(result interface{}) error {
	return p.pipe.One(result)
}

//Iter mongo Iter
func (p *MongoPipeDAL) Iter() MgoIterDAL {
	it := &MongoIterDAL{
		iter: p.pipe.Iter(),
	}
	return it
}

//Close invokes Close on iterator object
func (it *MongoIterDAL) Close() error {
	return it.iter.Close()
}

//Err invokes Err on iterator object
func (it *MongoIterDAL) Err() error {
	return it.iter.Err()
}

//Next invokes Next on iterator object
func (it *MongoIterDAL) Next(result interface{}) bool {
	return it.iter.Next(result)
}

//Timeout invokes timeout on iterator object
func (it *MongoIterDAL) Timeout() bool {
	return it.iter.Timeout()
}
