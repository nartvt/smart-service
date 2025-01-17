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
	"github.com/nartvt/smart-service/ent/predicate"
	"github.com/nartvt/smart-service/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUsername(s *string) *UserUpdate {
	if s != nil {
		uu.SetUsername(*s)
	}
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetAvatar sets the "avatar" field.
func (uu *UserUpdate) SetAvatar(s string) *UserUpdate {
	uu.mutation.SetAvatar(s)
	return uu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatar(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatar(*s)
	}
	return uu
}

// ClearAvatar clears the value of the "avatar" field.
func (uu *UserUpdate) ClearAvatar() *UserUpdate {
	uu.mutation.ClearAvatar()
	return uu
}

// SetBirthDay sets the "birth_day" field.
func (uu *UserUpdate) SetBirthDay(s string) *UserUpdate {
	uu.mutation.SetBirthDay(s)
	return uu
}

// SetNillableBirthDay sets the "birth_day" field if the given value is not nil.
func (uu *UserUpdate) SetNillableBirthDay(s *string) *UserUpdate {
	if s != nil {
		uu.SetBirthDay(*s)
	}
	return uu
}

// ClearBirthDay clears the value of the "birth_day" field.
func (uu *UserUpdate) ClearBirthDay() *UserUpdate {
	uu.mutation.ClearBirthDay()
	return uu
}

// SetFirstName sets the "first_name" field.
func (uu *UserUpdate) SetFirstName(s string) *UserUpdate {
	uu.mutation.SetFirstName(s)
	return uu
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableFirstName(s *string) *UserUpdate {
	if s != nil {
		uu.SetFirstName(*s)
	}
	return uu
}

// ClearFirstName clears the value of the "first_name" field.
func (uu *UserUpdate) ClearFirstName() *UserUpdate {
	uu.mutation.ClearFirstName()
	return uu
}

// SetLastName sets the "last_name" field.
func (uu *UserUpdate) SetLastName(s string) *UserUpdate {
	uu.mutation.SetLastName(s)
	return uu
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastName(s *string) *UserUpdate {
	if s != nil {
		uu.SetLastName(*s)
	}
	return uu
}

// ClearLastName clears the value of the "last_name" field.
func (uu *UserUpdate) ClearLastName() *UserUpdate {
	uu.mutation.ClearLastName()
	return uu
}

// SetPhoneNumber sets the "phone_number" field.
func (uu *UserUpdate) SetPhoneNumber(s string) *UserUpdate {
	uu.mutation.SetPhoneNumber(s)
	return uu
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhoneNumber(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhoneNumber(*s)
	}
	return uu
}

// ClearPhoneNumber clears the value of the "phone_number" field.
func (uu *UserUpdate) ClearPhoneNumber() *UserUpdate {
	uu.mutation.ClearPhoneNumber()
	return uu
}

// SetGender sets the "gender" field.
func (uu *UserUpdate) SetGender(i int) *UserUpdate {
	uu.mutation.ResetGender()
	uu.mutation.SetGender(i)
	return uu
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (uu *UserUpdate) SetNillableGender(i *int) *UserUpdate {
	if i != nil {
		uu.SetGender(*i)
	}
	return uu
}

// AddGender adds i to the "gender" field.
func (uu *UserUpdate) AddGender(i int) *UserUpdate {
	uu.mutation.AddGender(i)
	return uu
}

// ClearGender clears the value of the "gender" field.
func (uu *UserUpdate) ClearGender() *UserUpdate {
	uu.mutation.ClearGender()
	return uu
}

// SetStatus sets the "status" field.
func (uu *UserUpdate) SetStatus(s string) *UserUpdate {
	uu.mutation.SetStatus(s)
	return uu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uu *UserUpdate) SetNillableStatus(s *string) *UserUpdate {
	if s != nil {
		uu.SetStatus(*s)
	}
	return uu
}

// SetAddress sets the "address" field.
func (uu *UserUpdate) SetAddress(s string) *UserUpdate {
	uu.mutation.SetAddress(s)
	return uu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAddress(s *string) *UserUpdate {
	if s != nil {
		uu.SetAddress(*s)
	}
	return uu
}

// ClearAddress clears the value of the "address" field.
func (uu *UserUpdate) ClearAddress() *UserUpdate {
	uu.mutation.ClearAddress()
	return uu
}

// SetStreet sets the "street" field.
func (uu *UserUpdate) SetStreet(s string) *UserUpdate {
	uu.mutation.SetStreet(s)
	return uu
}

// SetNillableStreet sets the "street" field if the given value is not nil.
func (uu *UserUpdate) SetNillableStreet(s *string) *UserUpdate {
	if s != nil {
		uu.SetStreet(*s)
	}
	return uu
}

// ClearStreet clears the value of the "street" field.
func (uu *UserUpdate) ClearStreet() *UserUpdate {
	uu.mutation.ClearStreet()
	return uu
}

// SetDistrictID sets the "district_id" field.
func (uu *UserUpdate) SetDistrictID(s string) *UserUpdate {
	uu.mutation.SetDistrictID(s)
	return uu
}

// SetNillableDistrictID sets the "district_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDistrictID(s *string) *UserUpdate {
	if s != nil {
		uu.SetDistrictID(*s)
	}
	return uu
}

// ClearDistrictID clears the value of the "district_id" field.
func (uu *UserUpdate) ClearDistrictID() *UserUpdate {
	uu.mutation.ClearDistrictID()
	return uu
}

// SetCityID sets the "city_id" field.
func (uu *UserUpdate) SetCityID(s string) *UserUpdate {
	uu.mutation.SetCityID(s)
	return uu
}

// SetNillableCityID sets the "city_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCityID(s *string) *UserUpdate {
	if s != nil {
		uu.SetCityID(*s)
	}
	return uu
}

// ClearCityID clears the value of the "city_id" field.
func (uu *UserUpdate) ClearCityID() *UserUpdate {
	uu.mutation.ClearCityID()
	return uu
}

// SetCountryCode sets the "country_code" field.
func (uu *UserUpdate) SetCountryCode(s string) *UserUpdate {
	uu.mutation.SetCountryCode(s)
	return uu
}

// SetNillableCountryCode sets the "country_code" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCountryCode(s *string) *UserUpdate {
	if s != nil {
		uu.SetCountryCode(*s)
	}
	return uu
}

// ClearCountryCode clears the value of the "country_code" field.
func (uu *UserUpdate) ClearCountryCode() *UserUpdate {
	uu.mutation.ClearCountryCode()
	return uu
}

// SetHashPassword sets the "hash_password" field.
func (uu *UserUpdate) SetHashPassword(s string) *UserUpdate {
	uu.mutation.SetHashPassword(s)
	return uu
}

// SetNillableHashPassword sets the "hash_password" field if the given value is not nil.
func (uu *UserUpdate) SetNillableHashPassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetHashPassword(*s)
	}
	return uu
}

// SetVerified sets the "verified" field.
func (uu *UserUpdate) SetVerified(b bool) *UserUpdate {
	uu.mutation.SetVerified(b)
	return uu
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (uu *UserUpdate) SetNillableVerified(b *bool) *UserUpdate {
	if b != nil {
		uu.SetVerified(*b)
	}
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.HashPassword(); ok {
		if err := user.HashPasswordValidator(v); err != nil {
			return &ValidationError{Name: "hash_password", err: fmt.Errorf(`ent: validator failed for field "User.hash_password": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uu.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if value, ok := uu.mutation.BirthDay(); ok {
		_spec.SetField(user.FieldBirthDay, field.TypeString, value)
	}
	if uu.mutation.BirthDayCleared() {
		_spec.ClearField(user.FieldBirthDay, field.TypeString)
	}
	if value, ok := uu.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if uu.mutation.FirstNameCleared() {
		_spec.ClearField(user.FieldFirstName, field.TypeString)
	}
	if value, ok := uu.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if uu.mutation.LastNameCleared() {
		_spec.ClearField(user.FieldLastName, field.TypeString)
	}
	if value, ok := uu.mutation.PhoneNumber(); ok {
		_spec.SetField(user.FieldPhoneNumber, field.TypeString, value)
	}
	if uu.mutation.PhoneNumberCleared() {
		_spec.ClearField(user.FieldPhoneNumber, field.TypeString)
	}
	if value, ok := uu.mutation.Gender(); ok {
		_spec.SetField(user.FieldGender, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedGender(); ok {
		_spec.AddField(user.FieldGender, field.TypeInt, value)
	}
	if uu.mutation.GenderCleared() {
		_spec.ClearField(user.FieldGender, field.TypeInt)
	}
	if value, ok := uu.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeString, value)
	}
	if value, ok := uu.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
	}
	if uu.mutation.AddressCleared() {
		_spec.ClearField(user.FieldAddress, field.TypeString)
	}
	if value, ok := uu.mutation.Street(); ok {
		_spec.SetField(user.FieldStreet, field.TypeString, value)
	}
	if uu.mutation.StreetCleared() {
		_spec.ClearField(user.FieldStreet, field.TypeString)
	}
	if value, ok := uu.mutation.DistrictID(); ok {
		_spec.SetField(user.FieldDistrictID, field.TypeString, value)
	}
	if uu.mutation.DistrictIDCleared() {
		_spec.ClearField(user.FieldDistrictID, field.TypeString)
	}
	if value, ok := uu.mutation.CityID(); ok {
		_spec.SetField(user.FieldCityID, field.TypeString, value)
	}
	if uu.mutation.CityIDCleared() {
		_spec.ClearField(user.FieldCityID, field.TypeString)
	}
	if value, ok := uu.mutation.CountryCode(); ok {
		_spec.SetField(user.FieldCountryCode, field.TypeString, value)
	}
	if uu.mutation.CountryCodeCleared() {
		_spec.ClearField(user.FieldCountryCode, field.TypeString)
	}
	if value, ok := uu.mutation.HashPassword(); ok {
		_spec.SetField(user.FieldHashPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Verified(); ok {
		_spec.SetField(user.FieldVerified, field.TypeBool, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUsername(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUsername(*s)
	}
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetAvatar sets the "avatar" field.
func (uuo *UserUpdateOne) SetAvatar(s string) *UserUpdateOne {
	uuo.mutation.SetAvatar(s)
	return uuo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatar(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatar(*s)
	}
	return uuo
}

// ClearAvatar clears the value of the "avatar" field.
func (uuo *UserUpdateOne) ClearAvatar() *UserUpdateOne {
	uuo.mutation.ClearAvatar()
	return uuo
}

// SetBirthDay sets the "birth_day" field.
func (uuo *UserUpdateOne) SetBirthDay(s string) *UserUpdateOne {
	uuo.mutation.SetBirthDay(s)
	return uuo
}

// SetNillableBirthDay sets the "birth_day" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableBirthDay(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetBirthDay(*s)
	}
	return uuo
}

// ClearBirthDay clears the value of the "birth_day" field.
func (uuo *UserUpdateOne) ClearBirthDay() *UserUpdateOne {
	uuo.mutation.ClearBirthDay()
	return uuo
}

// SetFirstName sets the "first_name" field.
func (uuo *UserUpdateOne) SetFirstName(s string) *UserUpdateOne {
	uuo.mutation.SetFirstName(s)
	return uuo
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFirstName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetFirstName(*s)
	}
	return uuo
}

// ClearFirstName clears the value of the "first_name" field.
func (uuo *UserUpdateOne) ClearFirstName() *UserUpdateOne {
	uuo.mutation.ClearFirstName()
	return uuo
}

// SetLastName sets the "last_name" field.
func (uuo *UserUpdateOne) SetLastName(s string) *UserUpdateOne {
	uuo.mutation.SetLastName(s)
	return uuo
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLastName(*s)
	}
	return uuo
}

// ClearLastName clears the value of the "last_name" field.
func (uuo *UserUpdateOne) ClearLastName() *UserUpdateOne {
	uuo.mutation.ClearLastName()
	return uuo
}

// SetPhoneNumber sets the "phone_number" field.
func (uuo *UserUpdateOne) SetPhoneNumber(s string) *UserUpdateOne {
	uuo.mutation.SetPhoneNumber(s)
	return uuo
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhoneNumber(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhoneNumber(*s)
	}
	return uuo
}

// ClearPhoneNumber clears the value of the "phone_number" field.
func (uuo *UserUpdateOne) ClearPhoneNumber() *UserUpdateOne {
	uuo.mutation.ClearPhoneNumber()
	return uuo
}

// SetGender sets the "gender" field.
func (uuo *UserUpdateOne) SetGender(i int) *UserUpdateOne {
	uuo.mutation.ResetGender()
	uuo.mutation.SetGender(i)
	return uuo
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableGender(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetGender(*i)
	}
	return uuo
}

// AddGender adds i to the "gender" field.
func (uuo *UserUpdateOne) AddGender(i int) *UserUpdateOne {
	uuo.mutation.AddGender(i)
	return uuo
}

// ClearGender clears the value of the "gender" field.
func (uuo *UserUpdateOne) ClearGender() *UserUpdateOne {
	uuo.mutation.ClearGender()
	return uuo
}

// SetStatus sets the "status" field.
func (uuo *UserUpdateOne) SetStatus(s string) *UserUpdateOne {
	uuo.mutation.SetStatus(s)
	return uuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableStatus(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetStatus(*s)
	}
	return uuo
}

// SetAddress sets the "address" field.
func (uuo *UserUpdateOne) SetAddress(s string) *UserUpdateOne {
	uuo.mutation.SetAddress(s)
	return uuo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAddress(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAddress(*s)
	}
	return uuo
}

// ClearAddress clears the value of the "address" field.
func (uuo *UserUpdateOne) ClearAddress() *UserUpdateOne {
	uuo.mutation.ClearAddress()
	return uuo
}

// SetStreet sets the "street" field.
func (uuo *UserUpdateOne) SetStreet(s string) *UserUpdateOne {
	uuo.mutation.SetStreet(s)
	return uuo
}

// SetNillableStreet sets the "street" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableStreet(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetStreet(*s)
	}
	return uuo
}

// ClearStreet clears the value of the "street" field.
func (uuo *UserUpdateOne) ClearStreet() *UserUpdateOne {
	uuo.mutation.ClearStreet()
	return uuo
}

// SetDistrictID sets the "district_id" field.
func (uuo *UserUpdateOne) SetDistrictID(s string) *UserUpdateOne {
	uuo.mutation.SetDistrictID(s)
	return uuo
}

// SetNillableDistrictID sets the "district_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDistrictID(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetDistrictID(*s)
	}
	return uuo
}

// ClearDistrictID clears the value of the "district_id" field.
func (uuo *UserUpdateOne) ClearDistrictID() *UserUpdateOne {
	uuo.mutation.ClearDistrictID()
	return uuo
}

// SetCityID sets the "city_id" field.
func (uuo *UserUpdateOne) SetCityID(s string) *UserUpdateOne {
	uuo.mutation.SetCityID(s)
	return uuo
}

// SetNillableCityID sets the "city_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCityID(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetCityID(*s)
	}
	return uuo
}

// ClearCityID clears the value of the "city_id" field.
func (uuo *UserUpdateOne) ClearCityID() *UserUpdateOne {
	uuo.mutation.ClearCityID()
	return uuo
}

// SetCountryCode sets the "country_code" field.
func (uuo *UserUpdateOne) SetCountryCode(s string) *UserUpdateOne {
	uuo.mutation.SetCountryCode(s)
	return uuo
}

// SetNillableCountryCode sets the "country_code" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCountryCode(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetCountryCode(*s)
	}
	return uuo
}

// ClearCountryCode clears the value of the "country_code" field.
func (uuo *UserUpdateOne) ClearCountryCode() *UserUpdateOne {
	uuo.mutation.ClearCountryCode()
	return uuo
}

// SetHashPassword sets the "hash_password" field.
func (uuo *UserUpdateOne) SetHashPassword(s string) *UserUpdateOne {
	uuo.mutation.SetHashPassword(s)
	return uuo
}

// SetNillableHashPassword sets the "hash_password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableHashPassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetHashPassword(*s)
	}
	return uuo
}

// SetVerified sets the "verified" field.
func (uuo *UserUpdateOne) SetVerified(b bool) *UserUpdateOne {
	uuo.mutation.SetVerified(b)
	return uuo
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableVerified(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetVerified(*b)
	}
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Username(); ok {
		if err := user.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "User.username": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.HashPassword(); ok {
		if err := user.HashPasswordValidator(v); err != nil {
			return &ValidationError{Name: "hash_password", err: fmt.Errorf(`ent: validator failed for field "User.hash_password": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if uuo.mutation.AvatarCleared() {
		_spec.ClearField(user.FieldAvatar, field.TypeString)
	}
	if value, ok := uuo.mutation.BirthDay(); ok {
		_spec.SetField(user.FieldBirthDay, field.TypeString, value)
	}
	if uuo.mutation.BirthDayCleared() {
		_spec.ClearField(user.FieldBirthDay, field.TypeString)
	}
	if value, ok := uuo.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if uuo.mutation.FirstNameCleared() {
		_spec.ClearField(user.FieldFirstName, field.TypeString)
	}
	if value, ok := uuo.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if uuo.mutation.LastNameCleared() {
		_spec.ClearField(user.FieldLastName, field.TypeString)
	}
	if value, ok := uuo.mutation.PhoneNumber(); ok {
		_spec.SetField(user.FieldPhoneNumber, field.TypeString, value)
	}
	if uuo.mutation.PhoneNumberCleared() {
		_spec.ClearField(user.FieldPhoneNumber, field.TypeString)
	}
	if value, ok := uuo.mutation.Gender(); ok {
		_spec.SetField(user.FieldGender, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedGender(); ok {
		_spec.AddField(user.FieldGender, field.TypeInt, value)
	}
	if uuo.mutation.GenderCleared() {
		_spec.ClearField(user.FieldGender, field.TypeInt)
	}
	if value, ok := uuo.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
	}
	if uuo.mutation.AddressCleared() {
		_spec.ClearField(user.FieldAddress, field.TypeString)
	}
	if value, ok := uuo.mutation.Street(); ok {
		_spec.SetField(user.FieldStreet, field.TypeString, value)
	}
	if uuo.mutation.StreetCleared() {
		_spec.ClearField(user.FieldStreet, field.TypeString)
	}
	if value, ok := uuo.mutation.DistrictID(); ok {
		_spec.SetField(user.FieldDistrictID, field.TypeString, value)
	}
	if uuo.mutation.DistrictIDCleared() {
		_spec.ClearField(user.FieldDistrictID, field.TypeString)
	}
	if value, ok := uuo.mutation.CityID(); ok {
		_spec.SetField(user.FieldCityID, field.TypeString, value)
	}
	if uuo.mutation.CityIDCleared() {
		_spec.ClearField(user.FieldCityID, field.TypeString)
	}
	if value, ok := uuo.mutation.CountryCode(); ok {
		_spec.SetField(user.FieldCountryCode, field.TypeString, value)
	}
	if uuo.mutation.CountryCodeCleared() {
		_spec.ClearField(user.FieldCountryCode, field.TypeString)
	}
	if value, ok := uuo.mutation.HashPassword(); ok {
		_spec.SetField(user.FieldHashPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Verified(); ok {
		_spec.SetField(user.FieldVerified, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
