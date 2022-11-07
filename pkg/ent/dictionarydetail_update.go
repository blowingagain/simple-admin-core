// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionary"
	"github.com/suyuan32/simple-admin-core/pkg/ent/dictionarydetail"
	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
)

// DictionaryDetailUpdate is the builder for updating DictionaryDetail entities.
type DictionaryDetailUpdate struct {
	config
	hooks    []Hook
	mutation *DictionaryDetailMutation
}

// Where appends a list predicates to the DictionaryDetailUpdate builder.
func (ddu *DictionaryDetailUpdate) Where(ps ...predicate.DictionaryDetail) *DictionaryDetailUpdate {
	ddu.mutation.Where(ps...)
	return ddu
}

// SetUpdatedAt sets the "updated_at" field.
func (ddu *DictionaryDetailUpdate) SetUpdatedAt(t time.Time) *DictionaryDetailUpdate {
	ddu.mutation.SetUpdatedAt(t)
	return ddu
}

// SetStatus sets the "status" field.
func (ddu *DictionaryDetailUpdate) SetStatus(u uint8) *DictionaryDetailUpdate {
	ddu.mutation.ResetStatus()
	ddu.mutation.SetStatus(u)
	return ddu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ddu *DictionaryDetailUpdate) SetNillableStatus(u *uint8) *DictionaryDetailUpdate {
	if u != nil {
		ddu.SetStatus(*u)
	}
	return ddu
}

// AddStatus adds u to the "status" field.
func (ddu *DictionaryDetailUpdate) AddStatus(u int8) *DictionaryDetailUpdate {
	ddu.mutation.AddStatus(u)
	return ddu
}

// ClearStatus clears the value of the "status" field.
func (ddu *DictionaryDetailUpdate) ClearStatus() *DictionaryDetailUpdate {
	ddu.mutation.ClearStatus()
	return ddu
}

// SetTitle sets the "title" field.
func (ddu *DictionaryDetailUpdate) SetTitle(s string) *DictionaryDetailUpdate {
	ddu.mutation.SetTitle(s)
	return ddu
}

// SetKey sets the "key" field.
func (ddu *DictionaryDetailUpdate) SetKey(s string) *DictionaryDetailUpdate {
	ddu.mutation.SetKey(s)
	return ddu
}

// SetValue sets the "value" field.
func (ddu *DictionaryDetailUpdate) SetValue(s string) *DictionaryDetailUpdate {
	ddu.mutation.SetValue(s)
	return ddu
}

// SetDictionaryID sets the "dictionary" edge to the Dictionary entity by ID.
func (ddu *DictionaryDetailUpdate) SetDictionaryID(id uint64) *DictionaryDetailUpdate {
	ddu.mutation.SetDictionaryID(id)
	return ddu
}

// SetNillableDictionaryID sets the "dictionary" edge to the Dictionary entity by ID if the given value is not nil.
func (ddu *DictionaryDetailUpdate) SetNillableDictionaryID(id *uint64) *DictionaryDetailUpdate {
	if id != nil {
		ddu = ddu.SetDictionaryID(*id)
	}
	return ddu
}

// SetDictionary sets the "dictionary" edge to the Dictionary entity.
func (ddu *DictionaryDetailUpdate) SetDictionary(d *Dictionary) *DictionaryDetailUpdate {
	return ddu.SetDictionaryID(d.ID)
}

// Mutation returns the DictionaryDetailMutation object of the builder.
func (ddu *DictionaryDetailUpdate) Mutation() *DictionaryDetailMutation {
	return ddu.mutation
}

// ClearDictionary clears the "dictionary" edge to the Dictionary entity.
func (ddu *DictionaryDetailUpdate) ClearDictionary() *DictionaryDetailUpdate {
	ddu.mutation.ClearDictionary()
	return ddu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ddu *DictionaryDetailUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ddu.defaults()
	if len(ddu.hooks) == 0 {
		affected, err = ddu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DictionaryDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ddu.mutation = mutation
			affected, err = ddu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ddu.hooks) - 1; i >= 0; i-- {
			if ddu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ddu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ddu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ddu *DictionaryDetailUpdate) SaveX(ctx context.Context) int {
	affected, err := ddu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ddu *DictionaryDetailUpdate) Exec(ctx context.Context) error {
	_, err := ddu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddu *DictionaryDetailUpdate) ExecX(ctx context.Context) {
	if err := ddu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ddu *DictionaryDetailUpdate) defaults() {
	if _, ok := ddu.mutation.UpdatedAt(); !ok {
		v := dictionarydetail.UpdateDefaultUpdatedAt()
		ddu.mutation.SetUpdatedAt(v)
	}
}

func (ddu *DictionaryDetailUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dictionarydetail.Table,
			Columns: dictionarydetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: dictionarydetail.FieldID,
			},
		},
	}
	if ps := ddu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ddu.mutation.UpdatedAt(); ok {
		_spec.SetField(dictionarydetail.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ddu.mutation.Status(); ok {
		_spec.SetField(dictionarydetail.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := ddu.mutation.AddedStatus(); ok {
		_spec.AddField(dictionarydetail.FieldStatus, field.TypeUint8, value)
	}
	if ddu.mutation.StatusCleared() {
		_spec.ClearField(dictionarydetail.FieldStatus, field.TypeUint8)
	}
	if value, ok := ddu.mutation.Title(); ok {
		_spec.SetField(dictionarydetail.FieldTitle, field.TypeString, value)
	}
	if value, ok := ddu.mutation.Key(); ok {
		_spec.SetField(dictionarydetail.FieldKey, field.TypeString, value)
	}
	if value, ok := ddu.mutation.Value(); ok {
		_spec.SetField(dictionarydetail.FieldValue, field.TypeString, value)
	}
	if ddu.mutation.DictionaryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dictionarydetail.DictionaryTable,
			Columns: []string{dictionarydetail.DictionaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: dictionary.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ddu.mutation.DictionaryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dictionarydetail.DictionaryTable,
			Columns: []string{dictionarydetail.DictionaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: dictionary.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ddu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dictionarydetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DictionaryDetailUpdateOne is the builder for updating a single DictionaryDetail entity.
type DictionaryDetailUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DictionaryDetailMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (dduo *DictionaryDetailUpdateOne) SetUpdatedAt(t time.Time) *DictionaryDetailUpdateOne {
	dduo.mutation.SetUpdatedAt(t)
	return dduo
}

// SetStatus sets the "status" field.
func (dduo *DictionaryDetailUpdateOne) SetStatus(u uint8) *DictionaryDetailUpdateOne {
	dduo.mutation.ResetStatus()
	dduo.mutation.SetStatus(u)
	return dduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dduo *DictionaryDetailUpdateOne) SetNillableStatus(u *uint8) *DictionaryDetailUpdateOne {
	if u != nil {
		dduo.SetStatus(*u)
	}
	return dduo
}

// AddStatus adds u to the "status" field.
func (dduo *DictionaryDetailUpdateOne) AddStatus(u int8) *DictionaryDetailUpdateOne {
	dduo.mutation.AddStatus(u)
	return dduo
}

// ClearStatus clears the value of the "status" field.
func (dduo *DictionaryDetailUpdateOne) ClearStatus() *DictionaryDetailUpdateOne {
	dduo.mutation.ClearStatus()
	return dduo
}

// SetTitle sets the "title" field.
func (dduo *DictionaryDetailUpdateOne) SetTitle(s string) *DictionaryDetailUpdateOne {
	dduo.mutation.SetTitle(s)
	return dduo
}

// SetKey sets the "key" field.
func (dduo *DictionaryDetailUpdateOne) SetKey(s string) *DictionaryDetailUpdateOne {
	dduo.mutation.SetKey(s)
	return dduo
}

// SetValue sets the "value" field.
func (dduo *DictionaryDetailUpdateOne) SetValue(s string) *DictionaryDetailUpdateOne {
	dduo.mutation.SetValue(s)
	return dduo
}

// SetDictionaryID sets the "dictionary" edge to the Dictionary entity by ID.
func (dduo *DictionaryDetailUpdateOne) SetDictionaryID(id uint64) *DictionaryDetailUpdateOne {
	dduo.mutation.SetDictionaryID(id)
	return dduo
}

// SetNillableDictionaryID sets the "dictionary" edge to the Dictionary entity by ID if the given value is not nil.
func (dduo *DictionaryDetailUpdateOne) SetNillableDictionaryID(id *uint64) *DictionaryDetailUpdateOne {
	if id != nil {
		dduo = dduo.SetDictionaryID(*id)
	}
	return dduo
}

// SetDictionary sets the "dictionary" edge to the Dictionary entity.
func (dduo *DictionaryDetailUpdateOne) SetDictionary(d *Dictionary) *DictionaryDetailUpdateOne {
	return dduo.SetDictionaryID(d.ID)
}

// Mutation returns the DictionaryDetailMutation object of the builder.
func (dduo *DictionaryDetailUpdateOne) Mutation() *DictionaryDetailMutation {
	return dduo.mutation
}

// ClearDictionary clears the "dictionary" edge to the Dictionary entity.
func (dduo *DictionaryDetailUpdateOne) ClearDictionary() *DictionaryDetailUpdateOne {
	dduo.mutation.ClearDictionary()
	return dduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dduo *DictionaryDetailUpdateOne) Select(field string, fields ...string) *DictionaryDetailUpdateOne {
	dduo.fields = append([]string{field}, fields...)
	return dduo
}

// Save executes the query and returns the updated DictionaryDetail entity.
func (dduo *DictionaryDetailUpdateOne) Save(ctx context.Context) (*DictionaryDetail, error) {
	var (
		err  error
		node *DictionaryDetail
	)
	dduo.defaults()
	if len(dduo.hooks) == 0 {
		node, err = dduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DictionaryDetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dduo.mutation = mutation
			node, err = dduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dduo.hooks) - 1; i >= 0; i-- {
			if dduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dduo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dduo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DictionaryDetail)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DictionaryDetailMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dduo *DictionaryDetailUpdateOne) SaveX(ctx context.Context) *DictionaryDetail {
	node, err := dduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dduo *DictionaryDetailUpdateOne) Exec(ctx context.Context) error {
	_, err := dduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dduo *DictionaryDetailUpdateOne) ExecX(ctx context.Context) {
	if err := dduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dduo *DictionaryDetailUpdateOne) defaults() {
	if _, ok := dduo.mutation.UpdatedAt(); !ok {
		v := dictionarydetail.UpdateDefaultUpdatedAt()
		dduo.mutation.SetUpdatedAt(v)
	}
}

func (dduo *DictionaryDetailUpdateOne) sqlSave(ctx context.Context) (_node *DictionaryDetail, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dictionarydetail.Table,
			Columns: dictionarydetail.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: dictionarydetail.FieldID,
			},
		},
	}
	id, ok := dduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DictionaryDetail.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dictionarydetail.FieldID)
		for _, f := range fields {
			if !dictionarydetail.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dictionarydetail.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dduo.mutation.UpdatedAt(); ok {
		_spec.SetField(dictionarydetail.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := dduo.mutation.Status(); ok {
		_spec.SetField(dictionarydetail.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := dduo.mutation.AddedStatus(); ok {
		_spec.AddField(dictionarydetail.FieldStatus, field.TypeUint8, value)
	}
	if dduo.mutation.StatusCleared() {
		_spec.ClearField(dictionarydetail.FieldStatus, field.TypeUint8)
	}
	if value, ok := dduo.mutation.Title(); ok {
		_spec.SetField(dictionarydetail.FieldTitle, field.TypeString, value)
	}
	if value, ok := dduo.mutation.Key(); ok {
		_spec.SetField(dictionarydetail.FieldKey, field.TypeString, value)
	}
	if value, ok := dduo.mutation.Value(); ok {
		_spec.SetField(dictionarydetail.FieldValue, field.TypeString, value)
	}
	if dduo.mutation.DictionaryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dictionarydetail.DictionaryTable,
			Columns: []string{dictionarydetail.DictionaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: dictionary.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dduo.mutation.DictionaryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dictionarydetail.DictionaryTable,
			Columns: []string{dictionarydetail.DictionaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: dictionary.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DictionaryDetail{config: dduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dictionarydetail.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
