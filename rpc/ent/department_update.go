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
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-core/rpc/ent/department"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/ent/user"
)

// DepartmentUpdate is the builder for updating Department entities.
type DepartmentUpdate struct {
	config
	hooks     []Hook
	mutation  *DepartmentMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the DepartmentUpdate builder.
func (du *DepartmentUpdate) Where(ps ...predicate.Department) *DepartmentUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DepartmentUpdate) SetUpdatedAt(t time.Time) *DepartmentUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// SetStatus sets the "status" field.
func (du *DepartmentUpdate) SetStatus(u uint8) *DepartmentUpdate {
	du.mutation.ResetStatus()
	du.mutation.SetStatus(u)
	return du
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (du *DepartmentUpdate) SetNillableStatus(u *uint8) *DepartmentUpdate {
	if u != nil {
		du.SetStatus(*u)
	}
	return du
}

// AddStatus adds u to the "status" field.
func (du *DepartmentUpdate) AddStatus(u int8) *DepartmentUpdate {
	du.mutation.AddStatus(u)
	return du
}

// ClearStatus clears the value of the "status" field.
func (du *DepartmentUpdate) ClearStatus() *DepartmentUpdate {
	du.mutation.ClearStatus()
	return du
}

// SetSort sets the "sort" field.
func (du *DepartmentUpdate) SetSort(u uint32) *DepartmentUpdate {
	du.mutation.ResetSort()
	du.mutation.SetSort(u)
	return du
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (du *DepartmentUpdate) SetNillableSort(u *uint32) *DepartmentUpdate {
	if u != nil {
		du.SetSort(*u)
	}
	return du
}

// AddSort adds u to the "sort" field.
func (du *DepartmentUpdate) AddSort(u int32) *DepartmentUpdate {
	du.mutation.AddSort(u)
	return du
}

// SetName sets the "name" field.
func (du *DepartmentUpdate) SetName(s string) *DepartmentUpdate {
	du.mutation.SetName(s)
	return du
}

// SetNillableName sets the "name" field if the given value is not nil.
func (du *DepartmentUpdate) SetNillableName(s *string) *DepartmentUpdate {
	if s != nil {
		du.SetName(*s)
	}
	return du
}

// SetAncestors sets the "ancestors" field.
func (du *DepartmentUpdate) SetAncestors(s string) *DepartmentUpdate {
	du.mutation.SetAncestors(s)
	return du
}

// SetNillableAncestors sets the "ancestors" field if the given value is not nil.
func (du *DepartmentUpdate) SetNillableAncestors(s *string) *DepartmentUpdate {
	if s != nil {
		du.SetAncestors(*s)
	}
	return du
}

// ClearAncestors clears the value of the "ancestors" field.
func (du *DepartmentUpdate) ClearAncestors() *DepartmentUpdate {
	du.mutation.ClearAncestors()
	return du
}

// SetLeader sets the "leader" field.
func (du *DepartmentUpdate) SetLeader(s string) *DepartmentUpdate {
	du.mutation.SetLeader(s)
	return du
}

// SetNillableLeader sets the "leader" field if the given value is not nil.
func (du *DepartmentUpdate) SetNillableLeader(s *string) *DepartmentUpdate {
	if s != nil {
		du.SetLeader(*s)
	}
	return du
}

// SetPhone sets the "phone" field.
func (du *DepartmentUpdate) SetPhone(s string) *DepartmentUpdate {
	du.mutation.SetPhone(s)
	return du
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (du *DepartmentUpdate) SetNillablePhone(s *string) *DepartmentUpdate {
	if s != nil {
		du.SetPhone(*s)
	}
	return du
}

// SetEmail sets the "email" field.
func (du *DepartmentUpdate) SetEmail(s string) *DepartmentUpdate {
	du.mutation.SetEmail(s)
	return du
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (du *DepartmentUpdate) SetNillableEmail(s *string) *DepartmentUpdate {
	if s != nil {
		du.SetEmail(*s)
	}
	return du
}

// SetRemark sets the "remark" field.
func (du *DepartmentUpdate) SetRemark(s string) *DepartmentUpdate {
	du.mutation.SetRemark(s)
	return du
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (du *DepartmentUpdate) SetNillableRemark(s *string) *DepartmentUpdate {
	if s != nil {
		du.SetRemark(*s)
	}
	return du
}

// ClearRemark clears the value of the "remark" field.
func (du *DepartmentUpdate) ClearRemark() *DepartmentUpdate {
	du.mutation.ClearRemark()
	return du
}

// SetParentID sets the "parent_id" field.
func (du *DepartmentUpdate) SetParentID(u uint64) *DepartmentUpdate {
	du.mutation.SetParentID(u)
	return du
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (du *DepartmentUpdate) SetNillableParentID(u *uint64) *DepartmentUpdate {
	if u != nil {
		du.SetParentID(*u)
	}
	return du
}

// ClearParentID clears the value of the "parent_id" field.
func (du *DepartmentUpdate) ClearParentID() *DepartmentUpdate {
	du.mutation.ClearParentID()
	return du
}

// SetParent sets the "parent" edge to the Department entity.
func (du *DepartmentUpdate) SetParent(d *Department) *DepartmentUpdate {
	return du.SetParentID(d.ID)
}

// AddChildIDs adds the "children" edge to the Department entity by IDs.
func (du *DepartmentUpdate) AddChildIDs(ids ...uint64) *DepartmentUpdate {
	du.mutation.AddChildIDs(ids...)
	return du
}

// AddChildren adds the "children" edges to the Department entity.
func (du *DepartmentUpdate) AddChildren(d ...*Department) *DepartmentUpdate {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddChildIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (du *DepartmentUpdate) AddUserIDs(ids ...uuid.UUID) *DepartmentUpdate {
	du.mutation.AddUserIDs(ids...)
	return du
}

// AddUsers adds the "users" edges to the User entity.
func (du *DepartmentUpdate) AddUsers(u ...*User) *DepartmentUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.AddUserIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (du *DepartmentUpdate) Mutation() *DepartmentMutation {
	return du.mutation
}

// ClearParent clears the "parent" edge to the Department entity.
func (du *DepartmentUpdate) ClearParent() *DepartmentUpdate {
	du.mutation.ClearParent()
	return du
}

// ClearChildren clears all "children" edges to the Department entity.
func (du *DepartmentUpdate) ClearChildren() *DepartmentUpdate {
	du.mutation.ClearChildren()
	return du
}

// RemoveChildIDs removes the "children" edge to Department entities by IDs.
func (du *DepartmentUpdate) RemoveChildIDs(ids ...uint64) *DepartmentUpdate {
	du.mutation.RemoveChildIDs(ids...)
	return du
}

// RemoveChildren removes "children" edges to Department entities.
func (du *DepartmentUpdate) RemoveChildren(d ...*Department) *DepartmentUpdate {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveChildIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (du *DepartmentUpdate) ClearUsers() *DepartmentUpdate {
	du.mutation.ClearUsers()
	return du
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (du *DepartmentUpdate) RemoveUserIDs(ids ...uuid.UUID) *DepartmentUpdate {
	du.mutation.RemoveUserIDs(ids...)
	return du
}

// RemoveUsers removes "users" edges to User entities.
func (du *DepartmentUpdate) RemoveUsers(u ...*User) *DepartmentUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DepartmentUpdate) Save(ctx context.Context) (int, error) {
	du.defaults()
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DepartmentUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DepartmentUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DepartmentUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DepartmentUpdate) defaults() {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		v := department.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (du *DepartmentUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DepartmentUpdate {
	du.modifiers = append(du.modifiers, modifiers...)
	return du
}

func (du *DepartmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(department.Table, department.Columns, sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.SetField(department.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := du.mutation.Status(); ok {
		_spec.SetField(department.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := du.mutation.AddedStatus(); ok {
		_spec.AddField(department.FieldStatus, field.TypeUint8, value)
	}
	if du.mutation.StatusCleared() {
		_spec.ClearField(department.FieldStatus, field.TypeUint8)
	}
	if value, ok := du.mutation.Sort(); ok {
		_spec.SetField(department.FieldSort, field.TypeUint32, value)
	}
	if value, ok := du.mutation.AddedSort(); ok {
		_spec.AddField(department.FieldSort, field.TypeUint32, value)
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.SetField(department.FieldName, field.TypeString, value)
	}
	if value, ok := du.mutation.Ancestors(); ok {
		_spec.SetField(department.FieldAncestors, field.TypeString, value)
	}
	if du.mutation.AncestorsCleared() {
		_spec.ClearField(department.FieldAncestors, field.TypeString)
	}
	if value, ok := du.mutation.Leader(); ok {
		_spec.SetField(department.FieldLeader, field.TypeString, value)
	}
	if value, ok := du.mutation.Phone(); ok {
		_spec.SetField(department.FieldPhone, field.TypeString, value)
	}
	if value, ok := du.mutation.Email(); ok {
		_spec.SetField(department.FieldEmail, field.TypeString, value)
	}
	if value, ok := du.mutation.Remark(); ok {
		_spec.SetField(department.FieldRemark, field.TypeString, value)
	}
	if du.mutation.RemarkCleared() {
		_spec.ClearField(department.FieldRemark, field.TypeString)
	}
	if du.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   department.ParentTable,
			Columns: []string{department.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   department.ParentTable,
			Columns: []string{department.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.ChildrenTable,
			Columns: []string{department.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !du.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.ChildrenTable,
			Columns: []string{department.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.ChildrenTable,
			Columns: []string{department.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   department.UsersTable,
			Columns: []string{department.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedUsersIDs(); len(nodes) > 0 && !du.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   department.UsersTable,
			Columns: []string{department.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   department.UsersTable,
			Columns: []string{department.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(du.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{department.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DepartmentUpdateOne is the builder for updating a single Department entity.
type DepartmentUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *DepartmentMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DepartmentUpdateOne) SetUpdatedAt(t time.Time) *DepartmentUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// SetStatus sets the "status" field.
func (duo *DepartmentUpdateOne) SetStatus(u uint8) *DepartmentUpdateOne {
	duo.mutation.ResetStatus()
	duo.mutation.SetStatus(u)
	return duo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (duo *DepartmentUpdateOne) SetNillableStatus(u *uint8) *DepartmentUpdateOne {
	if u != nil {
		duo.SetStatus(*u)
	}
	return duo
}

// AddStatus adds u to the "status" field.
func (duo *DepartmentUpdateOne) AddStatus(u int8) *DepartmentUpdateOne {
	duo.mutation.AddStatus(u)
	return duo
}

// ClearStatus clears the value of the "status" field.
func (duo *DepartmentUpdateOne) ClearStatus() *DepartmentUpdateOne {
	duo.mutation.ClearStatus()
	return duo
}

// SetSort sets the "sort" field.
func (duo *DepartmentUpdateOne) SetSort(u uint32) *DepartmentUpdateOne {
	duo.mutation.ResetSort()
	duo.mutation.SetSort(u)
	return duo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (duo *DepartmentUpdateOne) SetNillableSort(u *uint32) *DepartmentUpdateOne {
	if u != nil {
		duo.SetSort(*u)
	}
	return duo
}

// AddSort adds u to the "sort" field.
func (duo *DepartmentUpdateOne) AddSort(u int32) *DepartmentUpdateOne {
	duo.mutation.AddSort(u)
	return duo
}

// SetName sets the "name" field.
func (duo *DepartmentUpdateOne) SetName(s string) *DepartmentUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (duo *DepartmentUpdateOne) SetNillableName(s *string) *DepartmentUpdateOne {
	if s != nil {
		duo.SetName(*s)
	}
	return duo
}

// SetAncestors sets the "ancestors" field.
func (duo *DepartmentUpdateOne) SetAncestors(s string) *DepartmentUpdateOne {
	duo.mutation.SetAncestors(s)
	return duo
}

// SetNillableAncestors sets the "ancestors" field if the given value is not nil.
func (duo *DepartmentUpdateOne) SetNillableAncestors(s *string) *DepartmentUpdateOne {
	if s != nil {
		duo.SetAncestors(*s)
	}
	return duo
}

// ClearAncestors clears the value of the "ancestors" field.
func (duo *DepartmentUpdateOne) ClearAncestors() *DepartmentUpdateOne {
	duo.mutation.ClearAncestors()
	return duo
}

// SetLeader sets the "leader" field.
func (duo *DepartmentUpdateOne) SetLeader(s string) *DepartmentUpdateOne {
	duo.mutation.SetLeader(s)
	return duo
}

// SetNillableLeader sets the "leader" field if the given value is not nil.
func (duo *DepartmentUpdateOne) SetNillableLeader(s *string) *DepartmentUpdateOne {
	if s != nil {
		duo.SetLeader(*s)
	}
	return duo
}

// SetPhone sets the "phone" field.
func (duo *DepartmentUpdateOne) SetPhone(s string) *DepartmentUpdateOne {
	duo.mutation.SetPhone(s)
	return duo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (duo *DepartmentUpdateOne) SetNillablePhone(s *string) *DepartmentUpdateOne {
	if s != nil {
		duo.SetPhone(*s)
	}
	return duo
}

// SetEmail sets the "email" field.
func (duo *DepartmentUpdateOne) SetEmail(s string) *DepartmentUpdateOne {
	duo.mutation.SetEmail(s)
	return duo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (duo *DepartmentUpdateOne) SetNillableEmail(s *string) *DepartmentUpdateOne {
	if s != nil {
		duo.SetEmail(*s)
	}
	return duo
}

// SetRemark sets the "remark" field.
func (duo *DepartmentUpdateOne) SetRemark(s string) *DepartmentUpdateOne {
	duo.mutation.SetRemark(s)
	return duo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (duo *DepartmentUpdateOne) SetNillableRemark(s *string) *DepartmentUpdateOne {
	if s != nil {
		duo.SetRemark(*s)
	}
	return duo
}

// ClearRemark clears the value of the "remark" field.
func (duo *DepartmentUpdateOne) ClearRemark() *DepartmentUpdateOne {
	duo.mutation.ClearRemark()
	return duo
}

// SetParentID sets the "parent_id" field.
func (duo *DepartmentUpdateOne) SetParentID(u uint64) *DepartmentUpdateOne {
	duo.mutation.SetParentID(u)
	return duo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (duo *DepartmentUpdateOne) SetNillableParentID(u *uint64) *DepartmentUpdateOne {
	if u != nil {
		duo.SetParentID(*u)
	}
	return duo
}

// ClearParentID clears the value of the "parent_id" field.
func (duo *DepartmentUpdateOne) ClearParentID() *DepartmentUpdateOne {
	duo.mutation.ClearParentID()
	return duo
}

// SetParent sets the "parent" edge to the Department entity.
func (duo *DepartmentUpdateOne) SetParent(d *Department) *DepartmentUpdateOne {
	return duo.SetParentID(d.ID)
}

// AddChildIDs adds the "children" edge to the Department entity by IDs.
func (duo *DepartmentUpdateOne) AddChildIDs(ids ...uint64) *DepartmentUpdateOne {
	duo.mutation.AddChildIDs(ids...)
	return duo
}

// AddChildren adds the "children" edges to the Department entity.
func (duo *DepartmentUpdateOne) AddChildren(d ...*Department) *DepartmentUpdateOne {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddChildIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (duo *DepartmentUpdateOne) AddUserIDs(ids ...uuid.UUID) *DepartmentUpdateOne {
	duo.mutation.AddUserIDs(ids...)
	return duo
}

// AddUsers adds the "users" edges to the User entity.
func (duo *DepartmentUpdateOne) AddUsers(u ...*User) *DepartmentUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.AddUserIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (duo *DepartmentUpdateOne) Mutation() *DepartmentMutation {
	return duo.mutation
}

// ClearParent clears the "parent" edge to the Department entity.
func (duo *DepartmentUpdateOne) ClearParent() *DepartmentUpdateOne {
	duo.mutation.ClearParent()
	return duo
}

// ClearChildren clears all "children" edges to the Department entity.
func (duo *DepartmentUpdateOne) ClearChildren() *DepartmentUpdateOne {
	duo.mutation.ClearChildren()
	return duo
}

// RemoveChildIDs removes the "children" edge to Department entities by IDs.
func (duo *DepartmentUpdateOne) RemoveChildIDs(ids ...uint64) *DepartmentUpdateOne {
	duo.mutation.RemoveChildIDs(ids...)
	return duo
}

// RemoveChildren removes "children" edges to Department entities.
func (duo *DepartmentUpdateOne) RemoveChildren(d ...*Department) *DepartmentUpdateOne {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveChildIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (duo *DepartmentUpdateOne) ClearUsers() *DepartmentUpdateOne {
	duo.mutation.ClearUsers()
	return duo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (duo *DepartmentUpdateOne) RemoveUserIDs(ids ...uuid.UUID) *DepartmentUpdateOne {
	duo.mutation.RemoveUserIDs(ids...)
	return duo
}

// RemoveUsers removes "users" edges to User entities.
func (duo *DepartmentUpdateOne) RemoveUsers(u ...*User) *DepartmentUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the DepartmentUpdate builder.
func (duo *DepartmentUpdateOne) Where(ps ...predicate.Department) *DepartmentUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DepartmentUpdateOne) Select(field string, fields ...string) *DepartmentUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Department entity.
func (duo *DepartmentUpdateOne) Save(ctx context.Context) (*Department, error) {
	duo.defaults()
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DepartmentUpdateOne) SaveX(ctx context.Context) *Department {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DepartmentUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DepartmentUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DepartmentUpdateOne) defaults() {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		v := department.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (duo *DepartmentUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DepartmentUpdateOne {
	duo.modifiers = append(duo.modifiers, modifiers...)
	return duo
}

func (duo *DepartmentUpdateOne) sqlSave(ctx context.Context) (_node *Department, err error) {
	_spec := sqlgraph.NewUpdateSpec(department.Table, department.Columns, sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Department.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, department.FieldID)
		for _, f := range fields {
			if !department.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != department.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.SetField(department.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := duo.mutation.Status(); ok {
		_spec.SetField(department.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := duo.mutation.AddedStatus(); ok {
		_spec.AddField(department.FieldStatus, field.TypeUint8, value)
	}
	if duo.mutation.StatusCleared() {
		_spec.ClearField(department.FieldStatus, field.TypeUint8)
	}
	if value, ok := duo.mutation.Sort(); ok {
		_spec.SetField(department.FieldSort, field.TypeUint32, value)
	}
	if value, ok := duo.mutation.AddedSort(); ok {
		_spec.AddField(department.FieldSort, field.TypeUint32, value)
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.SetField(department.FieldName, field.TypeString, value)
	}
	if value, ok := duo.mutation.Ancestors(); ok {
		_spec.SetField(department.FieldAncestors, field.TypeString, value)
	}
	if duo.mutation.AncestorsCleared() {
		_spec.ClearField(department.FieldAncestors, field.TypeString)
	}
	if value, ok := duo.mutation.Leader(); ok {
		_spec.SetField(department.FieldLeader, field.TypeString, value)
	}
	if value, ok := duo.mutation.Phone(); ok {
		_spec.SetField(department.FieldPhone, field.TypeString, value)
	}
	if value, ok := duo.mutation.Email(); ok {
		_spec.SetField(department.FieldEmail, field.TypeString, value)
	}
	if value, ok := duo.mutation.Remark(); ok {
		_spec.SetField(department.FieldRemark, field.TypeString, value)
	}
	if duo.mutation.RemarkCleared() {
		_spec.ClearField(department.FieldRemark, field.TypeString)
	}
	if duo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   department.ParentTable,
			Columns: []string{department.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   department.ParentTable,
			Columns: []string{department.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.ChildrenTable,
			Columns: []string{department.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !duo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.ChildrenTable,
			Columns: []string{department.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.ChildrenTable,
			Columns: []string{department.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   department.UsersTable,
			Columns: []string{department.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !duo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   department.UsersTable,
			Columns: []string{department.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   department.UsersTable,
			Columns: []string{department.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(duo.modifiers...)
	_node = &Department{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{department.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
