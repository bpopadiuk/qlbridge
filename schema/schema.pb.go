// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schema.proto

/*
Package schema is a generated protocol buffer package.

It is generated from these files:
	schema.proto

It has these top-level messages:
	Command
	SchemaPb
	TablePartition
	Partition
	TablePb
	FieldPb
	Index
	ConfigSchema
	ConfigSource
	ConfigNode
*/
package schema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Command_Operation int32

const (
	Command_Unknown   Command_Operation = 0
	Command_AddUpdate Command_Operation = 1
	Command_Drop      Command_Operation = 2
)

var Command_Operation_name = map[int32]string{
	0: "Unknown",
	1: "AddUpdate",
	2: "Drop",
}
var Command_Operation_value = map[string]int32{
	"Unknown":   0,
	"AddUpdate": 1,
	"Drop":      2,
}

func (x Command_Operation) String() string {
	return proto.EnumName(Command_Operation_name, int32(x))
}
func (Command_Operation) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Command defines a Schema Replication command/message such as
// Drop, Alter, Create-Schema, etc.  Used to replicate schema changes
// across servers.
type Command struct {
	Op     Command_Operation `protobuf:"varint,1,opt,name=op,enum=schema.Command_Operation" json:"op,omitempty"`
	Origin string            `protobuf:"bytes,2,opt,name=origin" json:"origin,omitempty"`
	Schema string            `protobuf:"bytes,3,opt,name=schema" json:"schema,omitempty"`
	Type   string            `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	Index  uint64            `protobuf:"varint,5,opt,name=index" json:"index,omitempty"`
	Ts     int64             `protobuf:"varint,6,opt,name=ts" json:"ts,omitempty"`
	Msg    []byte            `protobuf:"bytes,7,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Command) GetOp() Command_Operation {
	if m != nil {
		return m.Op
	}
	return Command_Unknown
}

func (m *Command) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Command) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *Command) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Command) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Command) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *Command) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

// SchemaPb defines the fields that define schema attributes, and
// can be serialized.
type SchemaPb struct {
	// Name of schema lowercased
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Tables is map of tables
	Tables map[string]*TablePb `protobuf:"bytes,2,rep,name=tables" json:"tables,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// source configuration
	Conf *ConfigSource `protobuf:"bytes,3,opt,name=conf" json:"conf,omitempty"`
}

func (m *SchemaPb) Reset()                    { *m = SchemaPb{} }
func (m *SchemaPb) String() string            { return proto.CompactTextString(m) }
func (*SchemaPb) ProtoMessage()               {}
func (*SchemaPb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SchemaPb) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SchemaPb) GetTables() map[string]*TablePb {
	if m != nil {
		return m.Tables
	}
	return nil
}

func (m *SchemaPb) GetConf() *ConfigSource {
	if m != nil {
		return m.Conf
	}
	return nil
}

// Partition describes a range of data (in a Table).
// left-key is contained in this partition
// right key is not contained in this partition, in the next partition.
// So any value >= left-key, and < right-key is contained herein.
type TablePartition struct {
	Table      string       `protobuf:"bytes,1,opt,name=table" json:"table,omitempty"`
	Keys       []string     `protobuf:"bytes,2,rep,name=keys" json:"keys,omitempty"`
	Partitions []*Partition `protobuf:"bytes,3,rep,name=partitions" json:"partitions,omitempty"`
}

func (m *TablePartition) Reset()                    { *m = TablePartition{} }
func (m *TablePartition) String() string            { return proto.CompactTextString(m) }
func (*TablePartition) ProtoMessage()               {}
func (*TablePartition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TablePartition) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *TablePartition) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *TablePartition) GetPartitions() []*Partition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

// Partition describes a range of data
// the left-key is contained in this partition
// the right key is not contained in this partition, in the next one
type Partition struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Left  string `protobuf:"bytes,2,opt,name=left" json:"left,omitempty"`
	Right string `protobuf:"bytes,3,opt,name=right" json:"right,omitempty"`
}

func (m *Partition) Reset()                    { *m = Partition{} }
func (m *Partition) String() string            { return proto.CompactTextString(m) }
func (*Partition) ProtoMessage()               {}
func (*Partition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Partition) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Partition) GetLeft() string {
	if m != nil {
		return m.Left
	}
	return ""
}

func (m *Partition) GetRight() string {
	if m != nil {
		return m.Right
	}
	return ""
}

// TablePb defines the fields that define table attributes, and
// can be serialized.
type TablePb struct {
	// Name of table lowercased
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Name of table (not lowercased)
	NameOriginal string `protobuf:"bytes,2,opt,name=nameOriginal" json:"nameOriginal,omitempty"`
	// some dbs are more hiearchical (table-column-family)
	Parent string `protobuf:"bytes,3,opt,name=parent" json:"parent,omitempty"`
	// Character set, default = utf8
	Charset uint32 `protobuf:"varint,4,opt,name=Charset" json:"Charset,omitempty"`
	// Partitions in this table, optional may be empty
	Partition *TablePartition `protobuf:"bytes,5,opt,name=partition" json:"partition,omitempty"`
	// Partition Count
	PartitionCt uint32 `protobuf:"varint,6,opt,name=PartitionCt" json:"PartitionCt,omitempty"`
	// List of indexes for this table
	Indexes []*Index `protobuf:"bytes,7,rep,name=indexes" json:"indexes,omitempty"`
	// context is additional arbitrary map values
	Context map[string]string `protobuf:"bytes,8,rep,name=context" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// List of Fields, in order
	Fieldpbs []*FieldPb `protobuf:"bytes,9,rep,name=fieldpbs" json:"fieldpbs,omitempty"`
}

func (m *TablePb) Reset()                    { *m = TablePb{} }
func (m *TablePb) String() string            { return proto.CompactTextString(m) }
func (*TablePb) ProtoMessage()               {}
func (*TablePb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TablePb) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TablePb) GetNameOriginal() string {
	if m != nil {
		return m.NameOriginal
	}
	return ""
}

func (m *TablePb) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *TablePb) GetCharset() uint32 {
	if m != nil {
		return m.Charset
	}
	return 0
}

func (m *TablePb) GetPartition() *TablePartition {
	if m != nil {
		return m.Partition
	}
	return nil
}

func (m *TablePb) GetPartitionCt() uint32 {
	if m != nil {
		return m.PartitionCt
	}
	return 0
}

func (m *TablePb) GetIndexes() []*Index {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *TablePb) GetContext() map[string]string {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *TablePb) GetFieldpbs() []*FieldPb {
	if m != nil {
		return m.Fieldpbs
	}
	return nil
}

// FieldPb defines attributes of a field/column that can
// be serialized and transported.
type FieldPb struct {
	Name        string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Key         string   `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Extra       string   `protobuf:"bytes,4,opt,name=extra" json:"extra,omitempty"`
	Data        string   `protobuf:"bytes,5,opt,name=data" json:"data,omitempty"`
	Length      uint32   `protobuf:"varint,6,opt,name=length" json:"length,omitempty"`
	Type        uint32   `protobuf:"varint,7,opt,name=type" json:"type,omitempty"`
	NativeType  uint32   `protobuf:"varint,8,opt,name=nativeType" json:"nativeType,omitempty"`
	DefLength   uint64   `protobuf:"varint,9,opt,name=defLength" json:"defLength,omitempty"`
	DefVal      []byte   `protobuf:"bytes,11,opt,name=defVal,proto3" json:"defVal,omitempty"`
	Indexed     bool     `protobuf:"varint,13,opt,name=indexed" json:"indexed,omitempty"`
	NoNulls     bool     `protobuf:"varint,14,opt,name=noNulls" json:"noNulls,omitempty"`
	Collation   string   `protobuf:"bytes,15,opt,name=collation" json:"collation,omitempty"`
	Roles       []string `protobuf:"bytes,16,rep,name=roles" json:"roles,omitempty"`
	Indexes     []*Index `protobuf:"bytes,17,rep,name=indexes" json:"indexes,omitempty"`
	// context is additional arbitrary map values
	Context  map[string]string `protobuf:"bytes,18,rep,name=context" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Position uint64            `protobuf:"varint,19,opt,name=position" json:"position,omitempty"`
}

func (m *FieldPb) Reset()                    { *m = FieldPb{} }
func (m *FieldPb) String() string            { return proto.CompactTextString(m) }
func (*FieldPb) ProtoMessage()               {}
func (*FieldPb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FieldPb) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FieldPb) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *FieldPb) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *FieldPb) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

func (m *FieldPb) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *FieldPb) GetLength() uint32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *FieldPb) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *FieldPb) GetNativeType() uint32 {
	if m != nil {
		return m.NativeType
	}
	return 0
}

func (m *FieldPb) GetDefLength() uint64 {
	if m != nil {
		return m.DefLength
	}
	return 0
}

func (m *FieldPb) GetDefVal() []byte {
	if m != nil {
		return m.DefVal
	}
	return nil
}

func (m *FieldPb) GetIndexed() bool {
	if m != nil {
		return m.Indexed
	}
	return false
}

func (m *FieldPb) GetNoNulls() bool {
	if m != nil {
		return m.NoNulls
	}
	return false
}

func (m *FieldPb) GetCollation() string {
	if m != nil {
		return m.Collation
	}
	return ""
}

func (m *FieldPb) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *FieldPb) GetIndexes() []*Index {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *FieldPb) GetContext() map[string]string {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *FieldPb) GetPosition() uint64 {
	if m != nil {
		return m.Position
	}
	return 0
}

// Index a description of how field(s) should be indexed for a table.
type Index struct {
	Name          string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Fields        []string `protobuf:"bytes,2,rep,name=fields" json:"fields,omitempty"`
	PrimaryKey    bool     `protobuf:"varint,3,opt,name=primaryKey" json:"primaryKey,omitempty"`
	HashPartition []string `protobuf:"bytes,4,rep,name=hashPartition" json:"hashPartition,omitempty"`
	PartitionSize int32    `protobuf:"varint,5,opt,name=partitionSize" json:"partitionSize,omitempty"`
}

func (m *Index) Reset()                    { *m = Index{} }
func (m *Index) String() string            { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()               {}
func (*Index) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Index) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Index) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Index) GetPrimaryKey() bool {
	if m != nil {
		return m.PrimaryKey
	}
	return false
}

func (m *Index) GetHashPartition() []string {
	if m != nil {
		return m.HashPartition
	}
	return nil
}

func (m *Index) GetPartitionSize() int32 {
	if m != nil {
		return m.PartitionSize
	}
	return 0
}

// ConfigSchema is the config block for Schema, the data-sources
// that make up this Virtual Schema.  Must have a name and list
// of sources to include.
type ConfigSchema struct {
	Name    string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Sources []string `protobuf:"bytes,2,rep,name=sources" json:"sources,omitempty"`
}

func (m *ConfigSchema) Reset()                    { *m = ConfigSchema{} }
func (m *ConfigSchema) String() string            { return proto.CompactTextString(m) }
func (*ConfigSchema) ProtoMessage()               {}
func (*ConfigSchema) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ConfigSchema) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ConfigSchema) GetSources() []string {
	if m != nil {
		return m.Sources
	}
	return nil
}

// ConfigSource are backend datasources ie : storage/database/csvfiles
// Each represents a single source type/config.  May belong to more
// than one schema.
type ConfigSource struct {
	Name           string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Schema         string            `protobuf:"bytes,2,opt,name=schema" json:"schema,omitempty"`
	Type           string            `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Tables_To_Load []string          `protobuf:"bytes,4,rep,name=tables_To_Load,json=tablesToLoad" json:"tables_To_Load,omitempty"`
	TableAliases   map[string]string `protobuf:"bytes,5,rep,name=tableAliases" json:"tableAliases,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Hosts          []string          `protobuf:"bytes,7,rep,name=hosts" json:"hosts,omitempty"`
	Settings       map[string]string `protobuf:"bytes,8,rep,name=settings" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Partitions     []*TablePartition `protobuf:"bytes,9,rep,name=partitions" json:"partitions,omitempty"`
	PartitionCt    uint32            `protobuf:"varint,10,opt,name=partitionCt" json:"partitionCt,omitempty"`
}

func (m *ConfigSource) Reset()                    { *m = ConfigSource{} }
func (m *ConfigSource) String() string            { return proto.CompactTextString(m) }
func (*ConfigSource) ProtoMessage()               {}
func (*ConfigSource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ConfigSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ConfigSource) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *ConfigSource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ConfigSource) GetTables_To_Load() []string {
	if m != nil {
		return m.Tables_To_Load
	}
	return nil
}

func (m *ConfigSource) GetTableAliases() map[string]string {
	if m != nil {
		return m.TableAliases
	}
	return nil
}

func (m *ConfigSource) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *ConfigSource) GetSettings() map[string]string {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *ConfigSource) GetPartitions() []*TablePartition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

func (m *ConfigSource) GetPartitionCt() uint32 {
	if m != nil {
		return m.PartitionCt
	}
	return 0
}

// ConfigNode are Servers/Services, ie a running instance of said Source
// - each must represent a single source type
// - normal use is a server, describing partitions of servers
// - may have arbitrary config info in Settings.
type ConfigNode struct {
	Name    string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Source  string            `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	Address string            `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	Context map[string]string `protobuf:"bytes,4,rep,name=context" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ConfigNode) Reset()                    { *m = ConfigNode{} }
func (m *ConfigNode) String() string            { return proto.CompactTextString(m) }
func (*ConfigNode) ProtoMessage()               {}
func (*ConfigNode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ConfigNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ConfigNode) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *ConfigNode) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ConfigNode) GetContext() map[string]string {
	if m != nil {
		return m.Context
	}
	return nil
}

func init() {
	proto.RegisterType((*Command)(nil), "schema.Command")
	proto.RegisterType((*SchemaPb)(nil), "schema.SchemaPb")
	proto.RegisterType((*TablePartition)(nil), "schema.TablePartition")
	proto.RegisterType((*Partition)(nil), "schema.Partition")
	proto.RegisterType((*TablePb)(nil), "schema.TablePb")
	proto.RegisterType((*FieldPb)(nil), "schema.FieldPb")
	proto.RegisterType((*Index)(nil), "schema.Index")
	proto.RegisterType((*ConfigSchema)(nil), "schema.ConfigSchema")
	proto.RegisterType((*ConfigSource)(nil), "schema.ConfigSource")
	proto.RegisterType((*ConfigNode)(nil), "schema.ConfigNode")
	proto.RegisterEnum("schema.Command_Operation", Command_Operation_name, Command_Operation_value)
}

func init() { proto.RegisterFile("schema.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 997 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0xc6, 0xce, 0x8f, 0xe3, 0x93, 0x9f, 0x66, 0x87, 0xd5, 0x6a, 0x88, 0x2a, 0x88, 0xac, 0x02,
	0x41, 0x48, 0x2b, 0x35, 0x54, 0x15, 0x14, 0x04, 0xaa, 0x96, 0x22, 0x51, 0xaa, 0x76, 0x35, 0xbb,
	0xe5, 0xb6, 0x9a, 0xc4, 0x93, 0xc4, 0x5a, 0xc7, 0x63, 0x79, 0x66, 0xcb, 0x86, 0xf7, 0xe0, 0x16,
	0x9e, 0x07, 0x5e, 0x83, 0x7b, 0x9e, 0x01, 0xcd, 0x99, 0xb1, 0x63, 0xab, 0x59, 0xa9, 0x55, 0xaf,
	0xe2, 0xef, 0x9b, 0x33, 0x67, 0xce, 0xcf, 0x37, 0x67, 0x02, 0x03, 0xb5, 0xdc, 0x88, 0x2d, 0x3f,
	0xcd, 0x0b, 0xa9, 0x25, 0xe9, 0x5a, 0x14, 0xfd, 0xeb, 0x41, 0x70, 0x26, 0xb7, 0x5b, 0x9e, 0xc5,
	0xe4, 0x0b, 0xf0, 0x65, 0x4e, 0xbd, 0xa9, 0x37, 0x1b, 0xcd, 0x3f, 0x3a, 0x75, 0xe6, 0x6e, 0xf1,
	0xf4, 0x45, 0x2e, 0x0a, 0xae, 0x13, 0x99, 0x31, 0x5f, 0xe6, 0xe4, 0x04, 0xba, 0xb2, 0x48, 0xd6,
	0x49, 0x46, 0xfd, 0xa9, 0x37, 0x0b, 0x99, 0x43, 0x86, 0xb7, 0xfb, 0x68, 0xcb, 0xf2, 0x16, 0x11,
	0x02, 0x6d, 0xbd, 0xcb, 0x05, 0x6d, 0x23, 0x8b, 0xdf, 0xe4, 0x18, 0x3a, 0x49, 0x16, 0x8b, 0x1b,
	0xda, 0x99, 0x7a, 0xb3, 0x36, 0xb3, 0x80, 0x8c, 0xc0, 0xd7, 0x8a, 0x76, 0xa7, 0xde, 0xac, 0xc5,
	0x7c, 0xad, 0xc8, 0x18, 0x5a, 0x5b, 0xb5, 0xa6, 0xc1, 0xd4, 0x9b, 0x0d, 0x98, 0xf9, 0x8c, 0xee,
	0x43, 0x58, 0x05, 0x43, 0xfa, 0x10, 0xbc, 0xcc, 0xae, 0x32, 0xf9, 0x5b, 0x36, 0xfe, 0x80, 0x0c,
	0x21, 0x7c, 0x1c, 0xc7, 0x2f, 0xf3, 0x98, 0x6b, 0x31, 0xf6, 0x48, 0x0f, 0xda, 0x3f, 0x16, 0x32,
	0x1f, 0xfb, 0xd1, 0x3f, 0x1e, 0xf4, 0x2e, 0x30, 0x92, 0xf3, 0x85, 0x89, 0x25, 0xe3, 0x5b, 0x81,
	0x89, 0x86, 0x0c, 0xbf, 0xc9, 0x03, 0xe8, 0x6a, 0xbe, 0x48, 0x85, 0xa2, 0xfe, 0xb4, 0x35, 0xeb,
	0xcf, 0xef, 0x96, 0xe9, 0x97, 0xbb, 0x4e, 0x2f, 0x71, 0xf9, 0x49, 0xa6, 0x8b, 0x1d, 0x73, 0xb6,
	0x64, 0x06, 0xed, 0xa5, 0xcc, 0x56, 0x98, 0x6b, 0x7f, 0x7e, 0xbc, 0x2f, 0x59, 0xb6, 0x4a, 0xd6,
	0x17, 0xf2, 0xba, 0x58, 0x0a, 0x86, 0x16, 0x93, 0xa7, 0xd0, 0xaf, 0x39, 0x30, 0x49, 0x5d, 0x89,
	0x9d, 0x8b, 0xc0, 0x7c, 0x92, 0x4f, 0xa1, 0xf3, 0x9a, 0xa7, 0xd7, 0x02, 0xeb, 0xd9, 0x9f, 0xdf,
	0x29, 0x7d, 0xe1, 0xae, 0xf3, 0x05, 0xb3, 0xab, 0x8f, 0xfc, 0xaf, 0xbd, 0x68, 0x0b, 0x23, 0xcb,
	0xf2, 0x42, 0x27, 0x58, 0x84, 0x63, 0xe8, 0x60, 0x44, 0xce, 0xa1, 0x05, 0x26, 0xcf, 0x2b, 0xb1,
	0xb3, 0x19, 0x85, 0x0c, 0xbf, 0xc9, 0x7d, 0x80, 0xbc, 0xdc, 0xa6, 0x68, 0x0b, 0x73, 0x3d, 0x2a,
	0xcf, 0xaa, 0x1c, 0xb2, 0x9a, 0x51, 0xf4, 0x04, 0xc2, 0xfd, 0x49, 0x23, 0xf0, 0x93, 0xd8, 0x1d,
	0xe3, 0x27, 0xb1, 0x39, 0x23, 0x15, 0x2b, 0xed, 0x54, 0x80, 0xdf, 0x26, 0x9a, 0x22, 0x59, 0x6f,
	0xb4, 0x93, 0x80, 0x05, 0xd1, 0x1f, 0x2d, 0x08, 0x5c, 0x32, 0x07, 0x3b, 0x10, 0xc1, 0xc0, 0xfc,
	0xbe, 0x40, 0x1d, 0xf1, 0xd4, 0x79, 0x6c, 0x70, 0x46, 0x5d, 0x39, 0x2f, 0x44, 0x56, 0xba, 0x76,
	0x88, 0x50, 0x08, 0xce, 0x36, 0xbc, 0x50, 0x42, 0xa3, 0xc0, 0x86, 0xac, 0x84, 0xe4, 0x01, 0x84,
	0x55, 0x2a, 0xa8, 0xb3, 0xfe, 0xfc, 0xa4, 0x59, 0xda, 0x2a, 0xe7, 0xbd, 0x21, 0x99, 0x42, 0xbf,
	0xe2, 0xcf, 0x34, 0x8a, 0x71, 0xc8, 0xea, 0x14, 0xf9, 0x1c, 0x02, 0x94, 0xab, 0x50, 0x34, 0xc0,
	0x22, 0x0e, 0x4b, 0xaf, 0x3f, 0x1b, 0x9a, 0x95, 0xab, 0xe4, 0x21, 0x04, 0x4b, 0x99, 0x69, 0x71,
	0xa3, 0x69, 0xaf, 0xa9, 0x2c, 0x57, 0x0c, 0xa3, 0x16, 0xb3, 0x6c, 0x95, 0x55, 0x1a, 0x93, 0x2f,
	0xa1, 0xb7, 0x4a, 0x44, 0x1a, 0xe7, 0x0b, 0x45, 0x43, 0xdc, 0x58, 0x49, 0xe2, 0x27, 0xc3, 0x9f,
	0x2f, 0x58, 0x65, 0x30, 0x79, 0x04, 0x83, 0xba, 0x97, 0x03, 0xf2, 0x3a, 0xae, 0xcb, 0x2b, 0xac,
	0xab, 0xe9, 0xcf, 0x36, 0x04, 0xce, 0xe3, 0xc1, 0xbe, 0x4c, 0xa1, 0x1f, 0x0b, 0xb5, 0x2c, 0x92,
	0x1c, 0x6b, 0x68, 0xf7, 0xd7, 0xa9, 0xf2, 0xb4, 0x56, 0xe3, 0x34, 0x71, 0xa3, 0x0b, 0xee, 0xae,
	0xbb, 0x05, 0xc6, 0x7b, 0xcc, 0x35, 0xc7, 0x36, 0x84, 0x0c, 0xbf, 0x4d, 0x47, 0x53, 0x91, 0xad,
	0xf5, 0xc6, 0x15, 0xd9, 0xa1, 0x6a, 0x5e, 0x04, 0xc8, 0xda, 0x79, 0xf1, 0x31, 0x40, 0xc6, 0x75,
	0xf2, 0x5a, 0x5c, 0x9a, 0x95, 0x1e, 0xae, 0xd4, 0x18, 0x72, 0x17, 0xc2, 0x58, 0xac, 0x9e, 0x59,
	0x77, 0x21, 0xce, 0x94, 0x3d, 0x61, 0x4e, 0x8a, 0xc5, 0xea, 0x57, 0x9e, 0xd2, 0x3e, 0x8e, 0x12,
	0x87, 0x8c, 0x76, 0x6c, 0xaf, 0x62, 0x3a, 0x9c, 0x7a, 0xb3, 0x5e, 0xd9, 0xba, 0xd8, 0xac, 0x64,
	0xf2, 0xf9, 0x75, 0x9a, 0x2a, 0x3a, 0xb2, 0x2b, 0x0e, 0x9a, 0x93, 0x96, 0x32, 0x4d, 0x71, 0x02,
	0xd1, 0x3b, 0x98, 0xce, 0x9e, 0x40, 0xfd, 0x4b, 0x33, 0x4a, 0xc6, 0x78, 0xf1, 0x2c, 0xa8, 0x2b,
	0xe6, 0xe8, 0x6d, 0x15, 0x43, 0x9a, 0x8a, 0x71, 0x6d, 0xba, 0x45, 0x31, 0x13, 0xe8, 0xe5, 0x52,
	0x59, 0xa5, 0x7f, 0x88, 0xd9, 0x57, 0xf8, 0xbd, 0x04, 0xf2, 0x97, 0x07, 0x1d, 0x0c, 0xf1, 0xa0,
	0x3c, 0x4e, 0xa0, 0x8b, 0x32, 0x2c, 0xc7, 0x8c, 0x43, 0xa6, 0x59, 0x79, 0x91, 0x6c, 0x79, 0xb1,
	0xfb, 0xc5, 0x69, 0xa3, 0xc7, 0x6a, 0x0c, 0xb9, 0x07, 0xc3, 0x0d, 0x57, 0x9b, 0xea, 0x4e, 0xd1,
	0x36, 0x6e, 0x6f, 0x92, 0xc6, 0xaa, 0xba, 0x95, 0x17, 0xc9, 0xef, 0x02, 0xb5, 0xd3, 0x61, 0x4d,
	0x32, 0xfa, 0x0e, 0xb3, 0x33, 0x23, 0xb7, 0x7a, 0x6c, 0xde, 0x88, 0x93, 0x42, 0xa0, 0x70, 0x20,
	0x97, 0x81, 0x96, 0x30, 0xfa, 0xaf, 0x55, 0x6d, 0x47, 0xe6, 0xb6, 0x34, 0xdd, 0xbb, 0xe6, 0x1f,
	0x7c, 0xd7, 0x5a, 0xb5, 0x77, 0xed, 0x1e, 0x8c, 0xec, 0xfb, 0xf0, 0xea, 0x52, 0xbe, 0x7a, 0x26,
	0x79, 0xec, 0x72, 0x1b, 0x58, 0xf6, 0x52, 0x1a, 0x8e, 0x3c, 0x05, 0x8b, 0x1f, 0xa7, 0x09, 0x57,
	0x42, 0xd1, 0x0e, 0xf6, 0xfa, 0xb3, 0x43, 0x6f, 0x88, 0x1d, 0x15, 0xce, 0xd0, 0x76, 0xbd, 0xb1,
	0xd7, 0x34, 0x6f, 0x23, 0x95, 0xb6, 0xb3, 0x28, 0x64, 0x16, 0x90, 0xef, 0xa1, 0xa7, 0x84, 0xd6,
	0x49, 0xb6, 0x56, 0x6e, 0xf6, 0x44, 0x07, 0xbd, 0x5f, 0x38, 0x23, 0xeb, 0xb9, 0xda, 0x43, 0x1e,
	0x36, 0xde, 0x0a, 0x3b, 0x84, 0x6e, 0x1b, 0x9e, 0x35, 0x4b, 0x33, 0x31, 0xf2, 0xda, 0xf4, 0x04,
	0x3b, 0x3d, 0x6b, 0xd4, 0xe4, 0x07, 0x38, 0x7a, 0x23, 0xa5, 0x77, 0xd1, 0xe4, 0xe4, 0x5b, 0x18,
	0x36, 0xa2, 0x7e, 0x27, 0x41, 0xff, 0xed, 0x01, 0xd8, 0x02, 0x3c, 0x97, 0xf1, 0xed, 0xed, 0xc6,
	0xe2, 0x54, 0xed, 0xb6, 0xd2, 0xa0, 0x10, 0xf0, 0x38, 0x2e, 0x84, 0x52, 0xae, 0xe3, 0x25, 0x24,
	0xdf, 0xec, 0x6f, 0x6d, 0x1b, 0x2b, 0xf5, 0x49, 0xb3, 0xd6, 0xe6, 0xa8, 0xc3, 0x17, 0xf7, 0x7d,
	0x2e, 0xe7, 0xa2, 0x8b, 0xff, 0xe6, 0xbe, 0xfa, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x49, 0xcb, 0xbb,
	0xca, 0xdd, 0x09, 0x00, 0x00,
}
