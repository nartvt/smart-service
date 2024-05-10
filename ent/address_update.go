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
	"github.com/nartvt/smart-service/ent/address"
	"github.com/nartvt/smart-service/ent/predicate"
)

// AddressUpdate is the builder for updating Address entities.
type AddressUpdate struct {
	config
	hooks    []Hook
	mutation *AddressMutation
}

// Where appends a list predicates to the AddressUpdate builder.
func (au *AddressUpdate) Where(ps ...predicate.Address) *AddressUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetStreet sets the "street" field.
func (au *AddressUpdate) SetStreet(s string) *AddressUpdate {
	au.mutation.SetStreet(s)
	return au
}

// SetNillableStreet sets the "street" field if the given value is not nil.
func (au *AddressUpdate) SetNillableStreet(s *string) *AddressUpdate {
	if s != nil {
		au.SetStreet(*s)
	}
	return au
}

// SetDistrict sets the "district" field.
func (au *AddressUpdate) SetDistrict(s string) *AddressUpdate {
	au.mutation.SetDistrict(s)
	return au
}

// SetNillableDistrict sets the "district" field if the given value is not nil.
func (au *AddressUpdate) SetNillableDistrict(s *string) *AddressUpdate {
	if s != nil {
		au.SetDistrict(*s)
	}
	return au
}

// SetCity sets the "city" field.
func (au *AddressUpdate) SetCity(s string) *AddressUpdate {
	au.mutation.SetCity(s)
	return au
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (au *AddressUpdate) SetNillableCity(s *string) *AddressUpdate {
	if s != nil {
		au.SetCity(*s)
	}
	return au
}

// SetCountry sets the "country" field.
func (au *AddressUpdate) SetCountry(s string) *AddressUpdate {
	au.mutation.SetCountry(s)
	return au
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (au *AddressUpdate) SetNillableCountry(s *string) *AddressUpdate {
	if s != nil {
		au.SetCountry(*s)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AddressUpdate) SetUpdatedAt(t time.Time) *AddressUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// Mutation returns the AddressMutation object of the builder.
func (au *AddressUpdate) Mutation() *AddressMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AddressUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AddressUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AddressUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AddressUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AddressUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := address.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

func (au *AddressUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(address.Table, address.Columns, sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Street(); ok {
		_spec.SetField(address.FieldStreet, field.TypeString, value)
	}
	if value, ok := au.mutation.District(); ok {
		_spec.SetField(address.FieldDistrict, field.TypeString, value)
	}
	if value, ok := au.mutation.City(); ok {
		_spec.SetField(address.FieldCity, field.TypeString, value)
	}
	if value, ok := au.mutation.Country(); ok {
		_spec.SetField(address.FieldCountry, field.TypeString, value)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(address.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{address.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AddressUpdateOne is the builder for updating a single Address entity.
type AddressUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AddressMutation
}

// SetStreet sets the "street" field.
func (auo *AddressUpdateOne) SetStreet(s string) *AddressUpdateOne {
	auo.mutation.SetStreet(s)
	return auo
}

// SetNillableStreet sets the "street" field if the given value is not nil.
func (auo *AddressUpdateOne) SetNillableStreet(s *string) *AddressUpdateOne {
	if s != nil {
		auo.SetStreet(*s)
	}
	return auo
}

// SetDistrict sets the "district" field.
func (auo *AddressUpdateOne) SetDistrict(s string) *AddressUpdateOne {
	auo.mutation.SetDistrict(s)
	return auo
}

// SetNillableDistrict sets the "district" field if the given value is not nil.
func (auo *AddressUpdateOne) SetNillableDistrict(s *string) *AddressUpdateOne {
	if s != nil {
		auo.SetDistrict(*s)
	}
	return auo
}

// SetCity sets the "city" field.
func (auo *AddressUpdateOne) SetCity(s string) *AddressUpdateOne {
	auo.mutation.SetCity(s)
	return auo
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (auo *AddressUpdateOne) SetNillableCity(s *string) *AddressUpdateOne {
	if s != nil {
		auo.SetCity(*s)
	}
	return auo
}

// SetCountry sets the "country" field.
func (auo *AddressUpdateOne) SetCountry(s string) *AddressUpdateOne {
	auo.mutation.SetCountry(s)
	return auo
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (auo *AddressUpdateOne) SetNillableCountry(s *string) *AddressUpdateOne {
	if s != nil {
		auo.SetCountry(*s)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AddressUpdateOne) SetUpdatedAt(t time.Time) *AddressUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// Mutation returns the AddressMutation object of the builder.
func (auo *AddressUpdateOne) Mutation() *AddressMutation {
	return auo.mutation
}

// Where appends a list predicates to the AddressUpdate builder.
func (auo *AddressUpdateOne) Where(ps ...predicate.Address) *AddressUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AddressUpdateOne) Select(field string, fields ...string) *AddressUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Address entity.
func (auo *AddressUpdateOne) Save(ctx context.Context) (*Address, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AddressUpdateOne) SaveX(ctx context.Context) *Address {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AddressUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AddressUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AddressUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := address.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

func (auo *AddressUpdateOne) sqlSave(ctx context.Context) (_node *Address, err error) {
	_spec := sqlgraph.NewUpdateSpec(address.Table, address.Columns, sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Address.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, address.FieldID)
		for _, f := range fields {
			if !address.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != address.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Street(); ok {
		_spec.SetField(address.FieldStreet, field.TypeString, value)
	}
	if value, ok := auo.mutation.District(); ok {
		_spec.SetField(address.FieldDistrict, field.TypeString, value)
	}
	if value, ok := auo.mutation.City(); ok {
		_spec.SetField(address.FieldCity, field.TypeString, value)
	}
	if value, ok := auo.mutation.Country(); ok {
		_spec.SetField(address.FieldCountry, field.TypeString, value)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(address.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Address{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{address.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}