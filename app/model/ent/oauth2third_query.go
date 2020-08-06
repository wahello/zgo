// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/suisrc/zgo/app/model/ent/oauth2third"
	"github.com/suisrc/zgo/app/model/ent/predicate"
)

// Oauth2ThirdQuery is the builder for querying Oauth2Third entities.
type Oauth2ThirdQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Oauth2Third
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (oq *Oauth2ThirdQuery) Where(ps ...predicate.Oauth2Third) *Oauth2ThirdQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit adds a limit step to the query.
func (oq *Oauth2ThirdQuery) Limit(limit int) *Oauth2ThirdQuery {
	oq.limit = &limit
	return oq
}

// Offset adds an offset step to the query.
func (oq *Oauth2ThirdQuery) Offset(offset int) *Oauth2ThirdQuery {
	oq.offset = &offset
	return oq
}

// Order adds an order step to the query.
func (oq *Oauth2ThirdQuery) Order(o ...OrderFunc) *Oauth2ThirdQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// First returns the first Oauth2Third entity in the query. Returns *NotFoundError when no oauth2third was found.
func (oq *Oauth2ThirdQuery) First(ctx context.Context) (*Oauth2Third, error) {
	os, err := oq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(os) == 0 {
		return nil, &NotFoundError{oauth2third.Label}
	}
	return os[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *Oauth2ThirdQuery) FirstX(ctx context.Context) *Oauth2Third {
	o, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return o
}

// FirstID returns the first Oauth2Third id in the query. Returns *NotFoundError when no id was found.
func (oq *Oauth2ThirdQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{oauth2third.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (oq *Oauth2ThirdQuery) FirstXID(ctx context.Context) int {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Oauth2Third entity in the query, returns an error if not exactly one entity was returned.
func (oq *Oauth2ThirdQuery) Only(ctx context.Context) (*Oauth2Third, error) {
	os, err := oq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(os) {
	case 1:
		return os[0], nil
	case 0:
		return nil, &NotFoundError{oauth2third.Label}
	default:
		return nil, &NotSingularError{oauth2third.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *Oauth2ThirdQuery) OnlyX(ctx context.Context) *Oauth2Third {
	o, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return o
}

// OnlyID returns the only Oauth2Third id in the query, returns an error if not exactly one id was returned.
func (oq *Oauth2ThirdQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = oq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{oauth2third.Label}
	default:
		err = &NotSingularError{oauth2third.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *Oauth2ThirdQuery) OnlyIDX(ctx context.Context) int {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Oauth2Thirds.
func (oq *Oauth2ThirdQuery) All(ctx context.Context) ([]*Oauth2Third, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oq *Oauth2ThirdQuery) AllX(ctx context.Context) []*Oauth2Third {
	os, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return os
}

// IDs executes the query and returns a list of Oauth2Third ids.
func (oq *Oauth2ThirdQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := oq.Select(oauth2third.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *Oauth2ThirdQuery) IDsX(ctx context.Context) []int {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *Oauth2ThirdQuery) Count(ctx context.Context) (int, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oq *Oauth2ThirdQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *Oauth2ThirdQuery) Exist(ctx context.Context) (bool, error) {
	if err := oq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *Oauth2ThirdQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *Oauth2ThirdQuery) Clone() *Oauth2ThirdQuery {
	return &Oauth2ThirdQuery{
		config:     oq.config,
		limit:      oq.limit,
		offset:     oq.offset,
		order:      append([]OrderFunc{}, oq.order...),
		unique:     append([]string{}, oq.unique...),
		predicates: append([]predicate.Oauth2Third{}, oq.predicates...),
		// clone intermediate query.
		sql:  oq.sql.Clone(),
		path: oq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Platform string `json:"platform,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Oauth2Third.Query().
//		GroupBy(oauth2third.FieldPlatform).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (oq *Oauth2ThirdQuery) GroupBy(field string, fields ...string) *Oauth2ThirdGroupBy {
	group := &Oauth2ThirdGroupBy{config: oq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Platform string `json:"platform,omitempty"`
//	}
//
//	client.Oauth2Third.Query().
//		Select(oauth2third.FieldPlatform).
//		Scan(ctx, &v)
//
func (oq *Oauth2ThirdQuery) Select(field string, fields ...string) *Oauth2ThirdSelect {
	selector := &Oauth2ThirdSelect{config: oq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oq.sqlQuery(), nil
	}
	return selector
}

func (oq *Oauth2ThirdQuery) prepareQuery(ctx context.Context) error {
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *Oauth2ThirdQuery) sqlAll(ctx context.Context) ([]*Oauth2Third, error) {
	var (
		nodes = []*Oauth2Third{}
		_spec = oq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &Oauth2Third{config: oq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (oq *Oauth2ThirdQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *Oauth2ThirdQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (oq *Oauth2ThirdQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   oauth2third.Table,
			Columns: oauth2third.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: oauth2third.FieldID,
			},
		},
		From:   oq.sql,
		Unique: true,
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *Oauth2ThirdQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(oauth2third.Table)
	selector := builder.Select(t1.Columns(oauth2third.Columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(oauth2third.Columns...)...)
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Oauth2ThirdGroupBy is the builder for group-by Oauth2Third entities.
type Oauth2ThirdGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *Oauth2ThirdGroupBy) Aggregate(fns ...AggregateFunc) *Oauth2ThirdGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the group-by query and scan the result into the given value.
func (ogb *Oauth2ThirdGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ogb.path(ctx)
	if err != nil {
		return err
	}
	ogb.sql = query
	return ogb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ogb *Oauth2ThirdGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ogb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ogb *Oauth2ThirdGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: Oauth2ThirdGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ogb *Oauth2ThirdGroupBy) StringsX(ctx context.Context) []string {
	v, err := ogb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (ogb *Oauth2ThirdGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ogb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{oauth2third.Label}
	default:
		err = fmt.Errorf("ent: Oauth2ThirdGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ogb *Oauth2ThirdGroupBy) StringX(ctx context.Context) string {
	v, err := ogb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ogb *Oauth2ThirdGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: Oauth2ThirdGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ogb *Oauth2ThirdGroupBy) IntsX(ctx context.Context) []int {
	v, err := ogb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (ogb *Oauth2ThirdGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ogb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{oauth2third.Label}
	default:
		err = fmt.Errorf("ent: Oauth2ThirdGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ogb *Oauth2ThirdGroupBy) IntX(ctx context.Context) int {
	v, err := ogb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ogb *Oauth2ThirdGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: Oauth2ThirdGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ogb *Oauth2ThirdGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ogb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (ogb *Oauth2ThirdGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ogb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{oauth2third.Label}
	default:
		err = fmt.Errorf("ent: Oauth2ThirdGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ogb *Oauth2ThirdGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ogb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ogb *Oauth2ThirdGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ogb.fields) > 1 {
		return nil, errors.New("ent: Oauth2ThirdGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ogb *Oauth2ThirdGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ogb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (ogb *Oauth2ThirdGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ogb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{oauth2third.Label}
	default:
		err = fmt.Errorf("ent: Oauth2ThirdGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ogb *Oauth2ThirdGroupBy) BoolX(ctx context.Context) bool {
	v, err := ogb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ogb *Oauth2ThirdGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ogb.sqlQuery().Query()
	if err := ogb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ogb *Oauth2ThirdGroupBy) sqlQuery() *sql.Selector {
	selector := ogb.sql
	columns := make([]string, 0, len(ogb.fields)+len(ogb.fns))
	columns = append(columns, ogb.fields...)
	for _, fn := range ogb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(ogb.fields...)
}

// Oauth2ThirdSelect is the builder for select fields of Oauth2Third entities.
type Oauth2ThirdSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (os *Oauth2ThirdSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := os.path(ctx)
	if err != nil {
		return err
	}
	os.sql = query
	return os.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (os *Oauth2ThirdSelect) ScanX(ctx context.Context, v interface{}) {
	if err := os.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (os *Oauth2ThirdSelect) Strings(ctx context.Context) ([]string, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: Oauth2ThirdSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (os *Oauth2ThirdSelect) StringsX(ctx context.Context) []string {
	v, err := os.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (os *Oauth2ThirdSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = os.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{oauth2third.Label}
	default:
		err = fmt.Errorf("ent: Oauth2ThirdSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (os *Oauth2ThirdSelect) StringX(ctx context.Context) string {
	v, err := os.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (os *Oauth2ThirdSelect) Ints(ctx context.Context) ([]int, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: Oauth2ThirdSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (os *Oauth2ThirdSelect) IntsX(ctx context.Context) []int {
	v, err := os.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (os *Oauth2ThirdSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = os.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{oauth2third.Label}
	default:
		err = fmt.Errorf("ent: Oauth2ThirdSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (os *Oauth2ThirdSelect) IntX(ctx context.Context) int {
	v, err := os.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (os *Oauth2ThirdSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: Oauth2ThirdSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (os *Oauth2ThirdSelect) Float64sX(ctx context.Context) []float64 {
	v, err := os.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (os *Oauth2ThirdSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = os.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{oauth2third.Label}
	default:
		err = fmt.Errorf("ent: Oauth2ThirdSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (os *Oauth2ThirdSelect) Float64X(ctx context.Context) float64 {
	v, err := os.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (os *Oauth2ThirdSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(os.fields) > 1 {
		return nil, errors.New("ent: Oauth2ThirdSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := os.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (os *Oauth2ThirdSelect) BoolsX(ctx context.Context) []bool {
	v, err := os.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (os *Oauth2ThirdSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = os.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{oauth2third.Label}
	default:
		err = fmt.Errorf("ent: Oauth2ThirdSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (os *Oauth2ThirdSelect) BoolX(ctx context.Context) bool {
	v, err := os.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (os *Oauth2ThirdSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := os.sqlQuery().Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (os *Oauth2ThirdSelect) sqlQuery() sql.Querier {
	selector := os.sql
	selector.Select(selector.Columns(os.fields...)...)
	return selector
}