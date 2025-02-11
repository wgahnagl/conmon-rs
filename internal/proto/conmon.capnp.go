// Code generated by capnpc-go. DO NOT EDIT.

package proto

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
	server "capnproto.org/go/capnp/v3/server"
	context "context"
)

type Conmon struct{ Client *capnp.Client }

// Conmon_TypeID is the unique identifier for the type Conmon.
const Conmon_TypeID = 0xb737e899dd6633f1

func (c Conmon) Version(ctx context.Context, params func(Conmon_version_Params) error) (Conmon_version_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xb737e899dd6633f1,
			MethodID:      0,
			InterfaceName: "proto/conmon.capnp:Conmon",
			MethodName:    "version",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Conmon_version_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return Conmon_version_Results_Future{Future: ans.Future()}, release
}

func (c Conmon) AddRef() Conmon {
	return Conmon{
		Client: c.Client.AddRef(),
	}
}

func (c Conmon) Release() {
	c.Client.Release()
}

// A Conmon_Server is a Conmon with a local implementation.
type Conmon_Server interface {
	Version(context.Context, Conmon_version) error
}

// Conmon_NewServer creates a new Server from an implementation of Conmon_Server.
func Conmon_NewServer(s Conmon_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Conmon_Methods(nil, s), s, c, policy)
}

// Conmon_ServerToClient creates a new Client from an implementation of Conmon_Server.
// The caller is responsible for calling Release on the returned Client.
func Conmon_ServerToClient(s Conmon_Server, policy *server.Policy) Conmon {
	return Conmon{Client: capnp.NewClient(Conmon_NewServer(s, policy))}
}

// Conmon_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Conmon_Methods(methods []server.Method, s Conmon_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xb737e899dd6633f1,
			MethodID:      0,
			InterfaceName: "proto/conmon.capnp:Conmon",
			MethodName:    "version",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Version(ctx, Conmon_version{call})
		},
	})

	return methods
}

// Conmon_version holds the state for a server call to Conmon.version.
// See server.Call for documentation.
type Conmon_version struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Conmon_version) Args() Conmon_version_Params {
	return Conmon_version_Params{Struct: c.Call.Args()}
}

// AllocResults allocates the results struct.
func (c Conmon_version) AllocResults() (Conmon_version_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Conmon_version_Results{Struct: r}, err
}

type Conmon_VersionResponse struct{ capnp.Struct }

// Conmon_VersionResponse_TypeID is the unique identifier for the type Conmon_VersionResponse.
const Conmon_VersionResponse_TypeID = 0xf34be5cbac1feed1

func NewConmon_VersionResponse(s *capnp.Segment) (Conmon_VersionResponse, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Conmon_VersionResponse{st}, err
}

func NewRootConmon_VersionResponse(s *capnp.Segment) (Conmon_VersionResponse, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Conmon_VersionResponse{st}, err
}

func ReadRootConmon_VersionResponse(msg *capnp.Message) (Conmon_VersionResponse, error) {
	root, err := msg.Root()
	return Conmon_VersionResponse{root.Struct()}, err
}

func (s Conmon_VersionResponse) String() string {
	str, _ := text.Marshal(0xf34be5cbac1feed1, s.Struct)
	return str
}

func (s Conmon_VersionResponse) Version() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Conmon_VersionResponse) HasVersion() bool {
	return s.Struct.HasPtr(0)
}

func (s Conmon_VersionResponse) VersionBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Conmon_VersionResponse) SetVersion(v string) error {
	return s.Struct.SetText(0, v)
}

// Conmon_VersionResponse_List is a list of Conmon_VersionResponse.
type Conmon_VersionResponse_List struct{ capnp.List }

// NewConmon_VersionResponse creates a new list of Conmon_VersionResponse.
func NewConmon_VersionResponse_List(s *capnp.Segment, sz int32) (Conmon_VersionResponse_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Conmon_VersionResponse_List{l}, err
}

func (s Conmon_VersionResponse_List) At(i int) Conmon_VersionResponse {
	return Conmon_VersionResponse{s.List.Struct(i)}
}

func (s Conmon_VersionResponse_List) Set(i int, v Conmon_VersionResponse) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Conmon_VersionResponse_List) String() string {
	str, _ := text.MarshalList(0xf34be5cbac1feed1, s.List)
	return str
}

// Conmon_VersionResponse_Future is a wrapper for a Conmon_VersionResponse promised by a client call.
type Conmon_VersionResponse_Future struct{ *capnp.Future }

func (p Conmon_VersionResponse_Future) Struct() (Conmon_VersionResponse, error) {
	s, err := p.Future.Struct()
	return Conmon_VersionResponse{s}, err
}

type Conmon_version_Params struct{ capnp.Struct }

// Conmon_version_Params_TypeID is the unique identifier for the type Conmon_version_Params.
const Conmon_version_Params_TypeID = 0xcc2f70676afee4e7

func NewConmon_version_Params(s *capnp.Segment) (Conmon_version_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Conmon_version_Params{st}, err
}

func NewRootConmon_version_Params(s *capnp.Segment) (Conmon_version_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Conmon_version_Params{st}, err
}

func ReadRootConmon_version_Params(msg *capnp.Message) (Conmon_version_Params, error) {
	root, err := msg.Root()
	return Conmon_version_Params{root.Struct()}, err
}

func (s Conmon_version_Params) String() string {
	str, _ := text.Marshal(0xcc2f70676afee4e7, s.Struct)
	return str
}

// Conmon_version_Params_List is a list of Conmon_version_Params.
type Conmon_version_Params_List struct{ capnp.List }

// NewConmon_version_Params creates a new list of Conmon_version_Params.
func NewConmon_version_Params_List(s *capnp.Segment, sz int32) (Conmon_version_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Conmon_version_Params_List{l}, err
}

func (s Conmon_version_Params_List) At(i int) Conmon_version_Params {
	return Conmon_version_Params{s.List.Struct(i)}
}

func (s Conmon_version_Params_List) Set(i int, v Conmon_version_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Conmon_version_Params_List) String() string {
	str, _ := text.MarshalList(0xcc2f70676afee4e7, s.List)
	return str
}

// Conmon_version_Params_Future is a wrapper for a Conmon_version_Params promised by a client call.
type Conmon_version_Params_Future struct{ *capnp.Future }

func (p Conmon_version_Params_Future) Struct() (Conmon_version_Params, error) {
	s, err := p.Future.Struct()
	return Conmon_version_Params{s}, err
}

type Conmon_version_Results struct{ capnp.Struct }

// Conmon_version_Results_TypeID is the unique identifier for the type Conmon_version_Results.
const Conmon_version_Results_TypeID = 0xe313695ea9477b30

func NewConmon_version_Results(s *capnp.Segment) (Conmon_version_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Conmon_version_Results{st}, err
}

func NewRootConmon_version_Results(s *capnp.Segment) (Conmon_version_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Conmon_version_Results{st}, err
}

func ReadRootConmon_version_Results(msg *capnp.Message) (Conmon_version_Results, error) {
	root, err := msg.Root()
	return Conmon_version_Results{root.Struct()}, err
}

func (s Conmon_version_Results) String() string {
	str, _ := text.Marshal(0xe313695ea9477b30, s.Struct)
	return str
}

func (s Conmon_version_Results) Response() (Conmon_VersionResponse, error) {
	p, err := s.Struct.Ptr(0)
	return Conmon_VersionResponse{Struct: p.Struct()}, err
}

func (s Conmon_version_Results) HasResponse() bool {
	return s.Struct.HasPtr(0)
}

func (s Conmon_version_Results) SetResponse(v Conmon_VersionResponse) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewResponse sets the response field to a newly
// allocated Conmon_VersionResponse struct, preferring placement in s's segment.
func (s Conmon_version_Results) NewResponse() (Conmon_VersionResponse, error) {
	ss, err := NewConmon_VersionResponse(s.Struct.Segment())
	if err != nil {
		return Conmon_VersionResponse{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Conmon_version_Results_List is a list of Conmon_version_Results.
type Conmon_version_Results_List struct{ capnp.List }

// NewConmon_version_Results creates a new list of Conmon_version_Results.
func NewConmon_version_Results_List(s *capnp.Segment, sz int32) (Conmon_version_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Conmon_version_Results_List{l}, err
}

func (s Conmon_version_Results_List) At(i int) Conmon_version_Results {
	return Conmon_version_Results{s.List.Struct(i)}
}

func (s Conmon_version_Results_List) Set(i int, v Conmon_version_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Conmon_version_Results_List) String() string {
	str, _ := text.MarshalList(0xe313695ea9477b30, s.List)
	return str
}

// Conmon_version_Results_Future is a wrapper for a Conmon_version_Results promised by a client call.
type Conmon_version_Results_Future struct{ *capnp.Future }

func (p Conmon_version_Results_Future) Struct() (Conmon_version_Results, error) {
	s, err := p.Future.Struct()
	return Conmon_version_Results{s}, err
}

func (p Conmon_version_Results_Future) Response() Conmon_VersionResponse_Future {
	return Conmon_VersionResponse_Future{Future: p.Future.Field(0, nil)}
}

const schema_ffaaf7385bc4adad = "x\xda\x8c\x8f\xb1K\x1bQ\x1c\xc7\xbf\xdf\xbbwM\x87" +
	"\x14r\xbcBK\x97v\xe8\xda\xa4%CK\x97\x84t" +
	"\x08M\x97\xbc\x83v)\x0aG8%b\xde;\xef\xa2" +
	"\x8b\x83\xe0\xe6\xe0\x1f\xe0\xe0\xe0\xe0\xa0\x98\xd9\x7f@p" +
	"\xd0Ew'\x15\x9d\x04qp\xf3I8\xcf\x04\x07q" +
	"{\xf0>\xdf\xdf\xf7\xf3-m\xd6\x9do\xde\xbe\x03\xa8" +
	"O\xde+{]\x9d:Y\xbb\xfc\xbe\x0b_\xbav0" +
	"\xd8\xfb\xff\xe3v\xdb\x02\x94\xbfy,\xff\xf2\x1d C" +
	"6\xe5\x0a\x0b\x80\xbd8\xbb\x9b\x99\x8e+\x87\xf0?\x10" +
	"\x10\x05\xa0:\xc7\x16!\xec\xd7\xc5\xe6\xd6dW\x9ef" +
	"?\xde\x90\xaeN0 (\xbb\xac\x81\xf6\xe8\xea\xe3\xce" +
	"\xc1\xf9\x9f\x9b!0*\xcd\xc0U\x06\x94\x1b\xc3\xa7\\" +
	"g\x0d_l\x9c\x98\xbe\xa9t\x8c\xa3{F\x97;a" +
	"\xac\xe3\x9f\xbf\x8c\xee\x19j%8~\x8c\xcb\xf6_\x94" +
	"\xa4]\xa3\x03Filt\x1aA\x09\xd7\x1b\x93e\xee" +
	"\xe6\xfb\x0d8\xbeWXZ\xc8\x12u\xb6\xc9\xc7.\xf1" +
	"\xb4K\x97\x1f\xb8\xcf\xed0\x09{)\xf0\x126\x88\xd2" +
	"\xf9\xd9>S%\\\x01\x08\x02\xfe\x9b\x16\xa0\x8a.\xd5" +
	"{\x876\xc9-\x01\x96FC@\x96\xf0\xacL>3" +
	"\xcb3j\x93\xe3\x1d\x0d@\xbdv\xa9\xde:\xcc\xf7\xb1" +
	"\x08\x87E\xf0>\x00\x00\xff\xff}\xde\x9ba"

func init() {
	schemas.Register(schema_ffaaf7385bc4adad,
		0xb737e899dd6633f1,
		0xcc2f70676afee4e7,
		0xe313695ea9477b30,
		0xf34be5cbac1feed1)
}
