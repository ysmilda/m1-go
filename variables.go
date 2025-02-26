package m1

import (
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"

	"github.com/ysmilda/m1-go/internals/m1binary"
	"github.com/ysmilda/m1-go/internals/m1client"
	"github.com/ysmilda/m1-go/modules/svi"
)

type Variable struct {
	Name string

	svi.Variable

	initialized bool
	module      *Module
	client      *m1client.Client
}

// NewVariable creates a new variable on the given module with the given name.
func NewVariable(target *Target, module *Module, name string) (*Variable, error) {
	v := &Variable{
		Name:   name,
		module: module,
		client: target.client,
	}

	err := v.initialize()
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (v *Variable) initialize() error {
	reply, err := call(v.client, v.module.info, svi.Procedures.GetExtendedAddress(svi.GetExtendedAddressCall{
		Name: v.Name,
	}))
	if err != nil {
		return err
	}
	v.initialized = true
	v.Address = reply.Address
	v.Format = reply.Format
	v.Length = reply.Length
	return nil
}

func (v Variable) String() string {
	t := v.GetGoDataType()

	rw := ""
	if v.IsReadable() {
		rw += "r"
	}
	if v.IsWritable() {
		rw += "w"
	}

	if v.IsBlock() {
		return fmt.Sprintf(
			"(%s) %s: %T[%d] (%s)",
			v.module.name, v.Name, reflect.ValueOf(t).Elem().Interface(), v.GetArrayLength(), rw,
		)
	}
	return fmt.Sprintf("(%s) %s: %T (%s)", v.module.name, v.Name, reflect.ValueOf(t).Elem().Interface(), rw)
}

func (v *Variable) SetValue(value any) error {
	t := v.GetGoDataType()
	if reflect.TypeOf(value) != reflect.TypeOf(v.GetGoDataType()).Elem() {
		return fmt.Errorf("expected %T, got %T", reflect.ValueOf(t).Elem().Interface(), value)
	}

	if !v.initialized {
		err := v.initialize()
		if err != nil {
			return err
		}
	}

	if v.IsBlock() {
		// TODO: extended call
		return errors.New("setvalue not yet implemented for block values")
	}

	buf, err := m1binary.Encode(value)
	if err != nil {
		return err
	}
	buf = append(buf, make([]byte, 4-len(buf))...)

	_, err = call(v.client, v.module.info, svi.Procedures.SetValue(svi.SetValueCall{
		Address: v.Address,
		Value:   binary.LittleEndian.Uint32(buf),
	}))
	return err
}

func (v *Variable) GetValue() (any, error) {
	if !v.initialized {
		err := v.initialize()
		if err != nil {
			return nil, err
		}
	}

	if v.IsBlock() {
		// TODO: extended call
		return nil, errors.New("getvalue not yet implemented for block values")
	}

	reply, err := call(v.client, v.module.info, svi.Procedures.GetValue(svi.GetValueCall{
		Address: v.Address,
	}))
	if err != nil {
		return nil, err
	}

	t := v.GetGoDataType()
	_, err = m1binary.Decode(binary.LittleEndian.AppendUint32(nil, reply.Value), &t)
	if err != nil {
		return nil, err
	}

	return reflect.ValueOf(t).Elem().Interface(), nil
}
