// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package realestate

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Rate             protoreflect.MessageDescriptor
	fd_Rate_symbol      protoreflect.FieldDescriptor
	fd_Rate_rate        protoreflect.FieldDescriptor
	fd_Rate_name        protoreflect.FieldDescriptor
	fd_Rate_description protoreflect.FieldDescriptor
	fd_Rate_creator     protoreflect.FieldDescriptor
)

func init() {
	file_realfin_realestate_rate_proto_init()
	md_Rate = File_realfin_realestate_rate_proto.Messages().ByName("Rate")
	fd_Rate_symbol = md_Rate.Fields().ByName("symbol")
	fd_Rate_rate = md_Rate.Fields().ByName("rate")
	fd_Rate_name = md_Rate.Fields().ByName("name")
	fd_Rate_description = md_Rate.Fields().ByName("description")
	fd_Rate_creator = md_Rate.Fields().ByName("creator")
}

var _ protoreflect.Message = (*fastReflection_Rate)(nil)

type fastReflection_Rate Rate

func (x *Rate) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Rate)(x)
}

func (x *Rate) slowProtoReflect() protoreflect.Message {
	mi := &file_realfin_realestate_rate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Rate_messageType fastReflection_Rate_messageType
var _ protoreflect.MessageType = fastReflection_Rate_messageType{}

type fastReflection_Rate_messageType struct{}

func (x fastReflection_Rate_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Rate)(nil)
}
func (x fastReflection_Rate_messageType) New() protoreflect.Message {
	return new(fastReflection_Rate)
}
func (x fastReflection_Rate_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Rate
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Rate) Descriptor() protoreflect.MessageDescriptor {
	return md_Rate
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Rate) Type() protoreflect.MessageType {
	return _fastReflection_Rate_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Rate) New() protoreflect.Message {
	return new(fastReflection_Rate)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Rate) Interface() protoreflect.ProtoMessage {
	return (*Rate)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Rate) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Symbol != "" {
		value := protoreflect.ValueOfString(x.Symbol)
		if !f(fd_Rate_symbol, value) {
			return
		}
	}
	if x.Rate != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Rate)
		if !f(fd_Rate_rate, value) {
			return
		}
	}
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_Rate_name, value) {
			return
		}
	}
	if x.Description != "" {
		value := protoreflect.ValueOfString(x.Description)
		if !f(fd_Rate_description, value) {
			return
		}
	}
	if x.Creator != "" {
		value := protoreflect.ValueOfString(x.Creator)
		if !f(fd_Rate_creator, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Rate) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "realfin.realestate.Rate.symbol":
		return x.Symbol != ""
	case "realfin.realestate.Rate.rate":
		return x.Rate != uint64(0)
	case "realfin.realestate.Rate.name":
		return x.Name != ""
	case "realfin.realestate.Rate.description":
		return x.Description != ""
	case "realfin.realestate.Rate.creator":
		return x.Creator != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: realfin.realestate.Rate"))
		}
		panic(fmt.Errorf("message realfin.realestate.Rate does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Rate) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "realfin.realestate.Rate.symbol":
		x.Symbol = ""
	case "realfin.realestate.Rate.rate":
		x.Rate = uint64(0)
	case "realfin.realestate.Rate.name":
		x.Name = ""
	case "realfin.realestate.Rate.description":
		x.Description = ""
	case "realfin.realestate.Rate.creator":
		x.Creator = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: realfin.realestate.Rate"))
		}
		panic(fmt.Errorf("message realfin.realestate.Rate does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Rate) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "realfin.realestate.Rate.symbol":
		value := x.Symbol
		return protoreflect.ValueOfString(value)
	case "realfin.realestate.Rate.rate":
		value := x.Rate
		return protoreflect.ValueOfUint64(value)
	case "realfin.realestate.Rate.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	case "realfin.realestate.Rate.description":
		value := x.Description
		return protoreflect.ValueOfString(value)
	case "realfin.realestate.Rate.creator":
		value := x.Creator
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: realfin.realestate.Rate"))
		}
		panic(fmt.Errorf("message realfin.realestate.Rate does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Rate) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "realfin.realestate.Rate.symbol":
		x.Symbol = value.Interface().(string)
	case "realfin.realestate.Rate.rate":
		x.Rate = value.Uint()
	case "realfin.realestate.Rate.name":
		x.Name = value.Interface().(string)
	case "realfin.realestate.Rate.description":
		x.Description = value.Interface().(string)
	case "realfin.realestate.Rate.creator":
		x.Creator = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: realfin.realestate.Rate"))
		}
		panic(fmt.Errorf("message realfin.realestate.Rate does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Rate) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "realfin.realestate.Rate.symbol":
		panic(fmt.Errorf("field symbol of message realfin.realestate.Rate is not mutable"))
	case "realfin.realestate.Rate.rate":
		panic(fmt.Errorf("field rate of message realfin.realestate.Rate is not mutable"))
	case "realfin.realestate.Rate.name":
		panic(fmt.Errorf("field name of message realfin.realestate.Rate is not mutable"))
	case "realfin.realestate.Rate.description":
		panic(fmt.Errorf("field description of message realfin.realestate.Rate is not mutable"))
	case "realfin.realestate.Rate.creator":
		panic(fmt.Errorf("field creator of message realfin.realestate.Rate is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: realfin.realestate.Rate"))
		}
		panic(fmt.Errorf("message realfin.realestate.Rate does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Rate) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "realfin.realestate.Rate.symbol":
		return protoreflect.ValueOfString("")
	case "realfin.realestate.Rate.rate":
		return protoreflect.ValueOfUint64(uint64(0))
	case "realfin.realestate.Rate.name":
		return protoreflect.ValueOfString("")
	case "realfin.realestate.Rate.description":
		return protoreflect.ValueOfString("")
	case "realfin.realestate.Rate.creator":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: realfin.realestate.Rate"))
		}
		panic(fmt.Errorf("message realfin.realestate.Rate does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Rate) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in realfin.realestate.Rate", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Rate) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Rate) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Rate) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Rate) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Rate)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Symbol)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Rate != 0 {
			n += 1 + runtime.Sov(uint64(x.Rate))
		}
		l = len(x.Name)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Description)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Creator)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Rate)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Creator) > 0 {
			i -= len(x.Creator)
			copy(dAtA[i:], x.Creator)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Creator)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.Description) > 0 {
			i -= len(x.Description)
			copy(dAtA[i:], x.Description)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Description)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
			i--
			dAtA[i] = 0x1a
		}
		if x.Rate != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Rate))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Symbol) > 0 {
			i -= len(x.Symbol)
			copy(dAtA[i:], x.Symbol)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Symbol)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Rate)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Rate: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Rate: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Symbol = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
				}
				x.Rate = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Rate |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Name = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Description = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Creator = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: realfin/realestate/rate.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Rate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol      string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Rate        uint64 `protobuf:"varint,2,opt,name=rate,proto3" json:"rate,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Creator     string `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (x *Rate) Reset() {
	*x = Rate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_realfin_realestate_rate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rate) ProtoMessage() {}

// Deprecated: Use Rate.ProtoReflect.Descriptor instead.
func (*Rate) Descriptor() ([]byte, []int) {
	return file_realfin_realestate_rate_proto_rawDescGZIP(), []int{0}
}

func (x *Rate) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Rate) GetRate() uint64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *Rate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Rate) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

var File_realfin_realestate_rate_proto protoreflect.FileDescriptor

var file_realfin_realestate_rate_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x61, 0x6c, 0x66, 0x69, 0x6e, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x72, 0x65, 0x61, 0x6c, 0x66, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x04, 0x52, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x42, 0xac, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d,
	0x2e, 0x72, 0x65, 0x61, 0x6c, 0x66, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x61, 0x6c, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x42, 0x09, 0x52, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x1e, 0x72, 0x65, 0x61, 0x6c, 0x66, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x61, 0x6c, 0x66, 0x69, 0x6e, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65,
	0xa2, 0x02, 0x03, 0x52, 0x52, 0x58, 0xaa, 0x02, 0x12, 0x52, 0x65, 0x61, 0x6c, 0x66, 0x69, 0x6e,
	0x2e, 0x52, 0x65, 0x61, 0x6c, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0xca, 0x02, 0x12, 0x52, 0x65,
	0x61, 0x6c, 0x66, 0x69, 0x6e, 0x5c, 0x52, 0x65, 0x61, 0x6c, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65,
	0xe2, 0x02, 0x1e, 0x52, 0x65, 0x61, 0x6c, 0x66, 0x69, 0x6e, 0x5c, 0x52, 0x65, 0x61, 0x6c, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x13, 0x52, 0x65, 0x61, 0x6c, 0x66, 0x69, 0x6e, 0x3a, 0x3a, 0x52, 0x65, 0x61,
	0x6c, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_realfin_realestate_rate_proto_rawDescOnce sync.Once
	file_realfin_realestate_rate_proto_rawDescData = file_realfin_realestate_rate_proto_rawDesc
)

func file_realfin_realestate_rate_proto_rawDescGZIP() []byte {
	file_realfin_realestate_rate_proto_rawDescOnce.Do(func() {
		file_realfin_realestate_rate_proto_rawDescData = protoimpl.X.CompressGZIP(file_realfin_realestate_rate_proto_rawDescData)
	})
	return file_realfin_realestate_rate_proto_rawDescData
}

var file_realfin_realestate_rate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_realfin_realestate_rate_proto_goTypes = []interface{}{
	(*Rate)(nil), // 0: realfin.realestate.Rate
}
var file_realfin_realestate_rate_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_realfin_realestate_rate_proto_init() }
func file_realfin_realestate_rate_proto_init() {
	if File_realfin_realestate_rate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_realfin_realestate_rate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_realfin_realestate_rate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_realfin_realestate_rate_proto_goTypes,
		DependencyIndexes: file_realfin_realestate_rate_proto_depIdxs,
		MessageInfos:      file_realfin_realestate_rate_proto_msgTypes,
	}.Build()
	File_realfin_realestate_rate_proto = out.File
	file_realfin_realestate_rate_proto_rawDesc = nil
	file_realfin_realestate_rate_proto_goTypes = nil
	file_realfin_realestate_rate_proto_depIdxs = nil
}
