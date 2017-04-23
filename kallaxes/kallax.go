// IMPORTANT! This is auto generated code by https://github.com/src-d/go-kallax
// Please, do not touch the code below, and if you do, do it under your own
// risk. Take into account that all the code you write here will be completely
// erased from earth the next time you generate the kallax models.
package kallaxes

import (
	"database/sql"
	"fmt"

	"gopkg.in/src-d/go-kallax.v1"
	"gopkg.in/src-d/go-kallax.v1/types"
)

var _ types.SQLType
var _ fmt.Formatter

// NewAirport returns a new instance of Airport.
func NewAirport() (record *Airport) {
	return new(Airport)
}

// GetID returns the primary key of the model.
func (r *Airport) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Airport) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "size":
		return types.Nullable(&r.Size), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Airport: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Airport) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "size":
		if r.Size == (*int)(nil) {
			return nil, nil
		}
		return r.Size, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Airport: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Airport) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model Airport has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *Airport) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model Airport has no relationships")
}

// AirportStore is the entity to access the records of the type Airport
// in the database.
type AirportStore struct {
	*kallax.Store
}

// NewAirportStore creates a new instance of AirportStore
// using a SQL database.
func NewAirportStore(db *sql.DB) *AirportStore {
	return &AirportStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *AirportStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *AirportStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Insert inserts a Airport in the database. A non-persisted object is
// required for this operation.
func (s *AirportStore) Insert(record *Airport) error {

	return s.Store.Insert(Schema.Airport.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *AirportStore) Update(record *Airport, cols ...kallax.SchemaField) (updated int64, err error) {

	return s.Store.Update(Schema.Airport.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *AirportStore) Save(record *Airport) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *AirportStore) Delete(record *Airport) error {

	return s.Store.Delete(Schema.Airport.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *AirportStore) Find(q *AirportQuery) (*AirportResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewAirportResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *AirportStore) MustFind(q *AirportQuery) *AirportResultSet {
	return NewAirportResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *AirportStore) Count(q *AirportQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *AirportStore) MustCount(q *AirportQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *AirportStore) FindOne(q *AirportQuery) (*Airport, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *AirportStore) MustFindOne(q *AirportQuery) *Airport {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Airport with the data in the database and
// makes it writable.
func (s *AirportStore) Reload(record *Airport) error {
	return s.Store.Reload(Schema.Airport.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *AirportStore) Transaction(callback func(*AirportStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&AirportStore{store})
	})
}

// AirportQuery is the object used to create queries for the Airport
// entity.
type AirportQuery struct {
	*kallax.BaseQuery
}

// NewAirportQuery returns a new instance of AirportQuery.
func NewAirportQuery() *AirportQuery {
	return &AirportQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Airport.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *AirportQuery) Select(columns ...kallax.SchemaField) *AirportQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *AirportQuery) SelectNot(columns ...kallax.SchemaField) *AirportQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *AirportQuery) Copy() *AirportQuery {
	return &AirportQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *AirportQuery) Order(cols ...kallax.ColumnOrder) *AirportQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *AirportQuery) BatchSize(size uint64) *AirportQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *AirportQuery) Limit(n uint64) *AirportQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *AirportQuery) Offset(n uint64) *AirportQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *AirportQuery) Where(cond kallax.Condition) *AirportQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *AirportQuery) FindByID(v ...int64) *AirportQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Airport.ID, values...))
}

// AirportResultSet is the set of results returned by a query to the
// database.
type AirportResultSet struct {
	ResultSet kallax.ResultSet
	last      *Airport
	lastErr   error
}

// NewAirportResultSet creates a new result set for rows of the type
// Airport.
func NewAirportResultSet(rs kallax.ResultSet) *AirportResultSet {
	return &AirportResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *AirportResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Airport.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Airport)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Airport")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *AirportResultSet) Get() (*Airport, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *AirportResultSet) ForEach(fn func(*Airport) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *AirportResultSet) All() ([]*Airport, error) {
	var result []*Airport
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *AirportResultSet) One() (*Airport, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *AirportResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *AirportResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewHangar returns a new instance of Hangar.
func NewHangar() (record *Hangar) {
	return new(Hangar)
}

// GetID returns the primary key of the model.
func (r *Hangar) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Hangar) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "name":
		return types.Nullable(&r.Name), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Hangar: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Hangar) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "name":
		if r.Name == (*string)(nil) {
			return nil, nil
		}
		return r.Name, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Hangar: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Hangar) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model Hangar has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *Hangar) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model Hangar has no relationships")
}

// HangarStore is the entity to access the records of the type Hangar
// in the database.
type HangarStore struct {
	*kallax.Store
}

// NewHangarStore creates a new instance of HangarStore
// using a SQL database.
func NewHangarStore(db *sql.DB) *HangarStore {
	return &HangarStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *HangarStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *HangarStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Insert inserts a Hangar in the database. A non-persisted object is
// required for this operation.
func (s *HangarStore) Insert(record *Hangar) error {

	return s.Store.Insert(Schema.Hangar.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *HangarStore) Update(record *Hangar, cols ...kallax.SchemaField) (updated int64, err error) {

	return s.Store.Update(Schema.Hangar.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *HangarStore) Save(record *Hangar) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *HangarStore) Delete(record *Hangar) error {

	return s.Store.Delete(Schema.Hangar.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *HangarStore) Find(q *HangarQuery) (*HangarResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewHangarResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *HangarStore) MustFind(q *HangarQuery) *HangarResultSet {
	return NewHangarResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *HangarStore) Count(q *HangarQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *HangarStore) MustCount(q *HangarQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *HangarStore) FindOne(q *HangarQuery) (*Hangar, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *HangarStore) MustFindOne(q *HangarQuery) *Hangar {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Hangar with the data in the database and
// makes it writable.
func (s *HangarStore) Reload(record *Hangar) error {
	return s.Store.Reload(Schema.Hangar.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *HangarStore) Transaction(callback func(*HangarStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&HangarStore{store})
	})
}

// HangarQuery is the object used to create queries for the Hangar
// entity.
type HangarQuery struct {
	*kallax.BaseQuery
}

// NewHangarQuery returns a new instance of HangarQuery.
func NewHangarQuery() *HangarQuery {
	return &HangarQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Hangar.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *HangarQuery) Select(columns ...kallax.SchemaField) *HangarQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *HangarQuery) SelectNot(columns ...kallax.SchemaField) *HangarQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *HangarQuery) Copy() *HangarQuery {
	return &HangarQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *HangarQuery) Order(cols ...kallax.ColumnOrder) *HangarQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *HangarQuery) BatchSize(size uint64) *HangarQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *HangarQuery) Limit(n uint64) *HangarQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *HangarQuery) Offset(n uint64) *HangarQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *HangarQuery) Where(cond kallax.Condition) *HangarQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *HangarQuery) FindByID(v ...int64) *HangarQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Hangar.ID, values...))
}

// HangarResultSet is the set of results returned by a query to the
// database.
type HangarResultSet struct {
	ResultSet kallax.ResultSet
	last      *Hangar
	lastErr   error
}

// NewHangarResultSet creates a new result set for rows of the type
// Hangar.
func NewHangarResultSet(rs kallax.ResultSet) *HangarResultSet {
	return &HangarResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *HangarResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Hangar.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Hangar)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Hangar")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *HangarResultSet) Get() (*Hangar, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *HangarResultSet) ForEach(fn func(*Hangar) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *HangarResultSet) All() ([]*Hangar, error) {
	var result []*Hangar
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *HangarResultSet) One() (*Hangar, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *HangarResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *HangarResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewJet returns a new instance of Jet.
func NewJet() (record *Jet) {
	return new(Jet)
}

// GetID returns the primary key of the model.
func (r *Jet) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Jet) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "pilot_id":
		return &r.PilotID, nil
	case "airport_id":
		return &r.AirportID, nil
	case "name":
		return &r.Name, nil
	case "color":
		return types.Nullable(&r.Color), nil
	case "uuid":
		return &r.UUID, nil
	case "identifier":
		return &r.Identifier, nil
	case "cargo":
		return types.Slice(&r.Cargo), nil
	case "manifest":
		return types.Slice(&r.Manifest), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Jet: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Jet) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "pilot_id":
		return r.PilotID, nil
	case "airport_id":
		return r.AirportID, nil
	case "name":
		return r.Name, nil
	case "color":
		if r.Color == (*string)(nil) {
			return nil, nil
		}
		return r.Color, nil
	case "uuid":
		return r.UUID, nil
	case "identifier":
		return r.Identifier, nil
	case "cargo":
		return types.Slice(r.Cargo), nil
	case "manifest":
		return types.Slice(r.Manifest), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Jet: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Jet) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model Jet has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *Jet) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model Jet has no relationships")
}

// JetStore is the entity to access the records of the type Jet
// in the database.
type JetStore struct {
	*kallax.Store
}

// NewJetStore creates a new instance of JetStore
// using a SQL database.
func NewJetStore(db *sql.DB) *JetStore {
	return &JetStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *JetStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *JetStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Insert inserts a Jet in the database. A non-persisted object is
// required for this operation.
func (s *JetStore) Insert(record *Jet) error {

	return s.Store.Insert(Schema.Jet.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *JetStore) Update(record *Jet, cols ...kallax.SchemaField) (updated int64, err error) {

	return s.Store.Update(Schema.Jet.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *JetStore) Save(record *Jet) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *JetStore) Delete(record *Jet) error {

	return s.Store.Delete(Schema.Jet.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *JetStore) Find(q *JetQuery) (*JetResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewJetResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *JetStore) MustFind(q *JetQuery) *JetResultSet {
	return NewJetResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *JetStore) Count(q *JetQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *JetStore) MustCount(q *JetQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *JetStore) FindOne(q *JetQuery) (*Jet, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *JetStore) MustFindOne(q *JetQuery) *Jet {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Jet with the data in the database and
// makes it writable.
func (s *JetStore) Reload(record *Jet) error {
	return s.Store.Reload(Schema.Jet.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *JetStore) Transaction(callback func(*JetStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&JetStore{store})
	})
}

// JetQuery is the object used to create queries for the Jet
// entity.
type JetQuery struct {
	*kallax.BaseQuery
}

// NewJetQuery returns a new instance of JetQuery.
func NewJetQuery() *JetQuery {
	return &JetQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Jet.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *JetQuery) Select(columns ...kallax.SchemaField) *JetQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *JetQuery) SelectNot(columns ...kallax.SchemaField) *JetQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *JetQuery) Copy() *JetQuery {
	return &JetQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *JetQuery) Order(cols ...kallax.ColumnOrder) *JetQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *JetQuery) BatchSize(size uint64) *JetQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *JetQuery) Limit(n uint64) *JetQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *JetQuery) Offset(n uint64) *JetQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *JetQuery) Where(cond kallax.Condition) *JetQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *JetQuery) FindByID(v ...int64) *JetQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Jet.ID, values...))
}

// FindByPilotID adds a new filter to the query that will require that
// the PilotID property is equal to the passed value.
func (q *JetQuery) FindByPilotID(cond kallax.ScalarCond, v int) *JetQuery {
	return q.Where(cond(Schema.Jet.PilotID, v))
}

// FindByAirportID adds a new filter to the query that will require that
// the AirportID property is equal to the passed value.
func (q *JetQuery) FindByAirportID(cond kallax.ScalarCond, v int) *JetQuery {
	return q.Where(cond(Schema.Jet.AirportID, v))
}

// FindByName adds a new filter to the query that will require that
// the Name property is equal to the passed value.
func (q *JetQuery) FindByName(v string) *JetQuery {
	return q.Where(kallax.Eq(Schema.Jet.Name, v))
}

// FindByUUID adds a new filter to the query that will require that
// the UUID property is equal to the passed value.
func (q *JetQuery) FindByUUID(v string) *JetQuery {
	return q.Where(kallax.Eq(Schema.Jet.UUID, v))
}

// FindByIdentifier adds a new filter to the query that will require that
// the Identifier property is equal to the passed value.
func (q *JetQuery) FindByIdentifier(v string) *JetQuery {
	return q.Where(kallax.Eq(Schema.Jet.Identifier, v))
}

// FindByCargo adds a new filter to the query that will require that
// the Cargo property contains all the passed values; if no passed values,
// it will do nothing.
func (q *JetQuery) FindByCargo(v ...byte) *JetQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.ArrayContains(Schema.Jet.Cargo, values...))
}

// FindByManifest adds a new filter to the query that will require that
// the Manifest property contains all the passed values; if no passed values,
// it will do nothing.
func (q *JetQuery) FindByManifest(v ...byte) *JetQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.ArrayContains(Schema.Jet.Manifest, values...))
}

// JetResultSet is the set of results returned by a query to the
// database.
type JetResultSet struct {
	ResultSet kallax.ResultSet
	last      *Jet
	lastErr   error
}

// NewJetResultSet creates a new result set for rows of the type
// Jet.
func NewJetResultSet(rs kallax.ResultSet) *JetResultSet {
	return &JetResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *JetResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Jet.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Jet)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Jet")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *JetResultSet) Get() (*Jet, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *JetResultSet) ForEach(fn func(*Jet) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *JetResultSet) All() ([]*Jet, error) {
	var result []*Jet
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *JetResultSet) One() (*Jet, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *JetResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *JetResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewLanguage returns a new instance of Language.
func NewLanguage() (record *Language) {
	return new(Language)
}

// GetID returns the primary key of the model.
func (r *Language) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Language) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "language":
		return &r.Language, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Language: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Language) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "language":
		return r.Language, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Language: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Language) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model Language has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *Language) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model Language has no relationships")
}

// LanguageStore is the entity to access the records of the type Language
// in the database.
type LanguageStore struct {
	*kallax.Store
}

// NewLanguageStore creates a new instance of LanguageStore
// using a SQL database.
func NewLanguageStore(db *sql.DB) *LanguageStore {
	return &LanguageStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *LanguageStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *LanguageStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Insert inserts a Language in the database. A non-persisted object is
// required for this operation.
func (s *LanguageStore) Insert(record *Language) error {

	return s.Store.Insert(Schema.Language.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *LanguageStore) Update(record *Language, cols ...kallax.SchemaField) (updated int64, err error) {

	return s.Store.Update(Schema.Language.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *LanguageStore) Save(record *Language) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *LanguageStore) Delete(record *Language) error {

	return s.Store.Delete(Schema.Language.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *LanguageStore) Find(q *LanguageQuery) (*LanguageResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewLanguageResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *LanguageStore) MustFind(q *LanguageQuery) *LanguageResultSet {
	return NewLanguageResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *LanguageStore) Count(q *LanguageQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *LanguageStore) MustCount(q *LanguageQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *LanguageStore) FindOne(q *LanguageQuery) (*Language, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *LanguageStore) MustFindOne(q *LanguageQuery) *Language {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Language with the data in the database and
// makes it writable.
func (s *LanguageStore) Reload(record *Language) error {
	return s.Store.Reload(Schema.Language.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *LanguageStore) Transaction(callback func(*LanguageStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&LanguageStore{store})
	})
}

// LanguageQuery is the object used to create queries for the Language
// entity.
type LanguageQuery struct {
	*kallax.BaseQuery
}

// NewLanguageQuery returns a new instance of LanguageQuery.
func NewLanguageQuery() *LanguageQuery {
	return &LanguageQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Language.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *LanguageQuery) Select(columns ...kallax.SchemaField) *LanguageQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *LanguageQuery) SelectNot(columns ...kallax.SchemaField) *LanguageQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *LanguageQuery) Copy() *LanguageQuery {
	return &LanguageQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *LanguageQuery) Order(cols ...kallax.ColumnOrder) *LanguageQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *LanguageQuery) BatchSize(size uint64) *LanguageQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *LanguageQuery) Limit(n uint64) *LanguageQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *LanguageQuery) Offset(n uint64) *LanguageQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *LanguageQuery) Where(cond kallax.Condition) *LanguageQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *LanguageQuery) FindByID(v ...int64) *LanguageQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Language.ID, values...))
}

// FindByLanguage adds a new filter to the query that will require that
// the Language property is equal to the passed value.
func (q *LanguageQuery) FindByLanguage(v string) *LanguageQuery {
	return q.Where(kallax.Eq(Schema.Language.Language, v))
}

// LanguageResultSet is the set of results returned by a query to the
// database.
type LanguageResultSet struct {
	ResultSet kallax.ResultSet
	last      *Language
	lastErr   error
}

// NewLanguageResultSet creates a new result set for rows of the type
// Language.
func NewLanguageResultSet(rs kallax.ResultSet) *LanguageResultSet {
	return &LanguageResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *LanguageResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Language.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Language)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Language")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *LanguageResultSet) Get() (*Language, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *LanguageResultSet) ForEach(fn func(*Language) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *LanguageResultSet) All() ([]*Language, error) {
	var result []*Language
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *LanguageResultSet) One() (*Language, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *LanguageResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *LanguageResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewLicense returns a new instance of License.
func NewLicense() (record *License) {
	return new(License)
}

// GetID returns the primary key of the model.
func (r *License) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *License) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "pilot_id":
		return types.Nullable(&r.PilotID), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in License: %s", col)
	}
}

// Value returns the value of the given column.
func (r *License) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "pilot_id":
		if r.PilotID == (*int64)(nil) {
			return nil, nil
		}
		return r.PilotID, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in License: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *License) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model License has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *License) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model License has no relationships")
}

// LicenseStore is the entity to access the records of the type License
// in the database.
type LicenseStore struct {
	*kallax.Store
}

// NewLicenseStore creates a new instance of LicenseStore
// using a SQL database.
func NewLicenseStore(db *sql.DB) *LicenseStore {
	return &LicenseStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *LicenseStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *LicenseStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Insert inserts a License in the database. A non-persisted object is
// required for this operation.
func (s *LicenseStore) Insert(record *License) error {

	return s.Store.Insert(Schema.License.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *LicenseStore) Update(record *License, cols ...kallax.SchemaField) (updated int64, err error) {

	return s.Store.Update(Schema.License.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *LicenseStore) Save(record *License) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *LicenseStore) Delete(record *License) error {

	return s.Store.Delete(Schema.License.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *LicenseStore) Find(q *LicenseQuery) (*LicenseResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewLicenseResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *LicenseStore) MustFind(q *LicenseQuery) *LicenseResultSet {
	return NewLicenseResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *LicenseStore) Count(q *LicenseQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *LicenseStore) MustCount(q *LicenseQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *LicenseStore) FindOne(q *LicenseQuery) (*License, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *LicenseStore) MustFindOne(q *LicenseQuery) *License {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the License with the data in the database and
// makes it writable.
func (s *LicenseStore) Reload(record *License) error {
	return s.Store.Reload(Schema.License.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *LicenseStore) Transaction(callback func(*LicenseStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&LicenseStore{store})
	})
}

// LicenseQuery is the object used to create queries for the License
// entity.
type LicenseQuery struct {
	*kallax.BaseQuery
}

// NewLicenseQuery returns a new instance of LicenseQuery.
func NewLicenseQuery() *LicenseQuery {
	return &LicenseQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.License.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *LicenseQuery) Select(columns ...kallax.SchemaField) *LicenseQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *LicenseQuery) SelectNot(columns ...kallax.SchemaField) *LicenseQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *LicenseQuery) Copy() *LicenseQuery {
	return &LicenseQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *LicenseQuery) Order(cols ...kallax.ColumnOrder) *LicenseQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *LicenseQuery) BatchSize(size uint64) *LicenseQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *LicenseQuery) Limit(n uint64) *LicenseQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *LicenseQuery) Offset(n uint64) *LicenseQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *LicenseQuery) Where(cond kallax.Condition) *LicenseQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *LicenseQuery) FindByID(v ...int64) *LicenseQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.License.ID, values...))
}

// LicenseResultSet is the set of results returned by a query to the
// database.
type LicenseResultSet struct {
	ResultSet kallax.ResultSet
	last      *License
	lastErr   error
}

// NewLicenseResultSet creates a new result set for rows of the type
// License.
func NewLicenseResultSet(rs kallax.ResultSet) *LicenseResultSet {
	return &LicenseResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *LicenseResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.License.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*License)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *License")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *LicenseResultSet) Get() (*License, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *LicenseResultSet) ForEach(fn func(*License) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *LicenseResultSet) All() ([]*License, error) {
	var result []*License
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *LicenseResultSet) One() (*License, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *LicenseResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *LicenseResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewPilot returns a new instance of Pilot.
func NewPilot() (record *Pilot) {
	return new(Pilot)
}

// GetID returns the primary key of the model.
func (r *Pilot) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Pilot) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "name":
		return &r.Name, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Pilot: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Pilot) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "name":
		return r.Name, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Pilot: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Pilot) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model Pilot has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *Pilot) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model Pilot has no relationships")
}

// PilotStore is the entity to access the records of the type Pilot
// in the database.
type PilotStore struct {
	*kallax.Store
}

// NewPilotStore creates a new instance of PilotStore
// using a SQL database.
func NewPilotStore(db *sql.DB) *PilotStore {
	return &PilotStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *PilotStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *PilotStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Insert inserts a Pilot in the database. A non-persisted object is
// required for this operation.
func (s *PilotStore) Insert(record *Pilot) error {

	return s.Store.Insert(Schema.Pilot.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *PilotStore) Update(record *Pilot, cols ...kallax.SchemaField) (updated int64, err error) {

	return s.Store.Update(Schema.Pilot.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *PilotStore) Save(record *Pilot) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *PilotStore) Delete(record *Pilot) error {

	return s.Store.Delete(Schema.Pilot.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *PilotStore) Find(q *PilotQuery) (*PilotResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewPilotResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *PilotStore) MustFind(q *PilotQuery) *PilotResultSet {
	return NewPilotResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *PilotStore) Count(q *PilotQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *PilotStore) MustCount(q *PilotQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *PilotStore) FindOne(q *PilotQuery) (*Pilot, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *PilotStore) MustFindOne(q *PilotQuery) *Pilot {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Pilot with the data in the database and
// makes it writable.
func (s *PilotStore) Reload(record *Pilot) error {
	return s.Store.Reload(Schema.Pilot.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *PilotStore) Transaction(callback func(*PilotStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&PilotStore{store})
	})
}

// PilotQuery is the object used to create queries for the Pilot
// entity.
type PilotQuery struct {
	*kallax.BaseQuery
}

// NewPilotQuery returns a new instance of PilotQuery.
func NewPilotQuery() *PilotQuery {
	return &PilotQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Pilot.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *PilotQuery) Select(columns ...kallax.SchemaField) *PilotQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *PilotQuery) SelectNot(columns ...kallax.SchemaField) *PilotQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *PilotQuery) Copy() *PilotQuery {
	return &PilotQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *PilotQuery) Order(cols ...kallax.ColumnOrder) *PilotQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *PilotQuery) BatchSize(size uint64) *PilotQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *PilotQuery) Limit(n uint64) *PilotQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *PilotQuery) Offset(n uint64) *PilotQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *PilotQuery) Where(cond kallax.Condition) *PilotQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *PilotQuery) FindByID(v ...int64) *PilotQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Pilot.ID, values...))
}

// FindByName adds a new filter to the query that will require that
// the Name property is equal to the passed value.
func (q *PilotQuery) FindByName(v string) *PilotQuery {
	return q.Where(kallax.Eq(Schema.Pilot.Name, v))
}

// PilotResultSet is the set of results returned by a query to the
// database.
type PilotResultSet struct {
	ResultSet kallax.ResultSet
	last      *Pilot
	lastErr   error
}

// NewPilotResultSet creates a new result set for rows of the type
// Pilot.
func NewPilotResultSet(rs kallax.ResultSet) *PilotResultSet {
	return &PilotResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *PilotResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Pilot.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Pilot)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Pilot")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *PilotResultSet) Get() (*Pilot, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *PilotResultSet) ForEach(fn func(*Pilot) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *PilotResultSet) All() ([]*Pilot, error) {
	var result []*Pilot
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *PilotResultSet) One() (*Pilot, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *PilotResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *PilotResultSet) Close() error {
	return rs.ResultSet.Close()
}

type schema struct {
	Airport  *schemaAirport
	Hangar   *schemaHangar
	Jet      *schemaJet
	Language *schemaLanguage
	License  *schemaLicense
	Pilot    *schemaPilot
}

type schemaAirport struct {
	*kallax.BaseSchema
	ID   kallax.SchemaField
	Size kallax.SchemaField
}

type schemaHangar struct {
	*kallax.BaseSchema
	ID   kallax.SchemaField
	Name kallax.SchemaField
}

type schemaJet struct {
	*kallax.BaseSchema
	ID         kallax.SchemaField
	PilotID    kallax.SchemaField
	AirportID  kallax.SchemaField
	Name       kallax.SchemaField
	Color      kallax.SchemaField
	UUID       kallax.SchemaField
	Identifier kallax.SchemaField
	Cargo      kallax.SchemaField
	Manifest   kallax.SchemaField
}

type schemaLanguage struct {
	*kallax.BaseSchema
	ID       kallax.SchemaField
	Language kallax.SchemaField
}

type schemaLicense struct {
	*kallax.BaseSchema
	ID      kallax.SchemaField
	PilotID kallax.SchemaField
}

type schemaPilot struct {
	*kallax.BaseSchema
	ID   kallax.SchemaField
	Name kallax.SchemaField
}

var Schema = &schema{
	Airport: &schemaAirport{
		BaseSchema: kallax.NewBaseSchema(
			"airports",
			"__airport",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(Airport)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("size"),
		),
		ID:   kallax.NewSchemaField("id"),
		Size: kallax.NewSchemaField("size"),
	},
	Hangar: &schemaHangar{
		BaseSchema: kallax.NewBaseSchema(
			"hangars",
			"__hangar",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(Hangar)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("name"),
		),
		ID:   kallax.NewSchemaField("id"),
		Name: kallax.NewSchemaField("name"),
	},
	Jet: &schemaJet{
		BaseSchema: kallax.NewBaseSchema(
			"jets",
			"__jet",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(Jet)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("pilot_id"),
			kallax.NewSchemaField("airport_id"),
			kallax.NewSchemaField("name"),
			kallax.NewSchemaField("color"),
			kallax.NewSchemaField("uuid"),
			kallax.NewSchemaField("identifier"),
			kallax.NewSchemaField("cargo"),
			kallax.NewSchemaField("manifest"),
		),
		ID:         kallax.NewSchemaField("id"),
		PilotID:    kallax.NewSchemaField("pilot_id"),
		AirportID:  kallax.NewSchemaField("airport_id"),
		Name:       kallax.NewSchemaField("name"),
		Color:      kallax.NewSchemaField("color"),
		UUID:       kallax.NewSchemaField("uuid"),
		Identifier: kallax.NewSchemaField("identifier"),
		Cargo:      kallax.NewSchemaField("cargo"),
		Manifest:   kallax.NewSchemaField("manifest"),
	},
	Language: &schemaLanguage{
		BaseSchema: kallax.NewBaseSchema(
			"languages",
			"__language",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(Language)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("language"),
		),
		ID:       kallax.NewSchemaField("id"),
		Language: kallax.NewSchemaField("language"),
	},
	License: &schemaLicense{
		BaseSchema: kallax.NewBaseSchema(
			"licenses",
			"__license",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(License)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("pilot_id"),
		),
		ID:      kallax.NewSchemaField("id"),
		PilotID: kallax.NewSchemaField("pilot_id"),
	},
	Pilot: &schemaPilot{
		BaseSchema: kallax.NewBaseSchema(
			"pilots",
			"__pilot",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(Pilot)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("name"),
		),
		ID:   kallax.NewSchemaField("id"),
		Name: kallax.NewSchemaField("name"),
	},
}
