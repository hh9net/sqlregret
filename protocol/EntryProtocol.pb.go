// Code generated by protoc-gen-go.
// source: EntryProtocol.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	EntryProtocol.proto

It has these top-level messages:
	Entry
	Header
	Column
	RowData
	RowChange
	TransactionBegin
	TransactionEnd
	Pair
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// *打散后的事件类型，主要用于标识事务的开始，变更数据，结束*
type EntryType int32

const (
	EntryType_TRANSACTIONBEGIN EntryType = 1
	EntryType_ROWDATA          EntryType = 2
	EntryType_TRANSACTIONEND   EntryType = 3
	// * 心跳类型，内部使用，外部暂不可见，可忽略 *
	EntryType_HEARTBEAT EntryType = 4
)

var EntryType_name = map[int32]string{
	1: "TRANSACTIONBEGIN",
	2: "ROWDATA",
	3: "TRANSACTIONEND",
	4: "HEARTBEAT",
}
var EntryType_value = map[string]int32{
	"TRANSACTIONBEGIN": 1,
	"ROWDATA":          2,
	"TRANSACTIONEND":   3,
	"HEARTBEAT":        4,
}

func (x EntryType) Enum() *EntryType {
	p := new(EntryType)
	*p = x
	return p
}
func (x EntryType) String() string {
	return proto.EnumName(EntryType_name, int32(x))
}
func (x *EntryType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EntryType_value, data, "EntryType")
	if err != nil {
		return err
	}
	*x = EntryType(value)
	return nil
}

// * 事件类型 *
type EventType int32

const (
	EventType_INSERT   EventType = 1
	EventType_UPDATE   EventType = 2
	EventType_DELETE   EventType = 3
	EventType_CREATE   EventType = 4
	EventType_ALTER    EventType = 5
	EventType_ERASE    EventType = 6
	EventType_QUERY    EventType = 7
	EventType_TRUNCATE EventType = 8
	EventType_RENAME   EventType = 9
	// *CREATE INDEX*
	EventType_CINDEX EventType = 10
	EventType_DINDEX EventType = 11
)

var EventType_name = map[int32]string{
	1:  "INSERT",
	2:  "UPDATE",
	3:  "DELETE",
	4:  "CREATE",
	5:  "ALTER",
	6:  "ERASE",
	7:  "QUERY",
	8:  "TRUNCATE",
	9:  "RENAME",
	10: "CINDEX",
	11: "DINDEX",
}
var EventType_value = map[string]int32{
	"INSERT":   1,
	"UPDATE":   2,
	"DELETE":   3,
	"CREATE":   4,
	"ALTER":    5,
	"ERASE":    6,
	"QUERY":    7,
	"TRUNCATE": 8,
	"RENAME":   9,
	"CINDEX":   10,
	"DINDEX":   11,
}

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}
func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}
func (x *EventType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EventType_value, data, "EventType")
	if err != nil {
		return err
	}
	*x = EventType(value)
	return nil
}

// *数据库类型*
type Type int32

const (
	Type_ORACLE Type = 1
	Type_MYSQL  Type = 2
	Type_PGSQL  Type = 3
)

var Type_name = map[int32]string{
	1: "ORACLE",
	2: "MYSQL",
	3: "PGSQL",
}
var Type_value = map[string]int32{
	"ORACLE": 1,
	"MYSQL":  2,
	"PGSQL":  3,
}

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}
func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}
func (x *Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Type_value, data, "Type")
	if err != nil {
		return err
	}
	*x = Type(value)
	return nil
}

// ***************************************************************
// message model
// 如果要在Enum中新增类型，确保以前的类型的下标值不变.
// **************************************************************
type Entry struct {
	// *协议头部信息*
	Header *Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// *打散后的事件类型*
	EntryType *EntryType `protobuf:"varint,2,opt,name=entryType,enum=protocol.EntryType,def=2" json:"entryType,omitempty"`
	// *传输的二进制数组*
	StoreValue       []byte `protobuf:"bytes,3,opt,name=storeValue" json:"storeValue,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}

const Default_Entry_EntryType EntryType = EntryType_ROWDATA

func (m *Entry) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Entry) SetHeader(header *Header) {
	m.Header = header
}

func (m *Entry) GetEntryType() EntryType {
	if m != nil && m.EntryType != nil {
		return *m.EntryType
	}
	return Default_Entry_EntryType
}

func (m *Entry) SetEntryType(entryType EntryType) {
	if nil == m.EntryType {
		m.EntryType = new(EntryType)
	}
	*m.EntryType = entryType
}

func (m *Entry) GetStoreValue() []byte {
	if m != nil {
		return m.StoreValue
	}
	return nil
}

func (m *Entry) SetStoreValue(storeValue []byte) {
	m.StoreValue = storeValue
}

// *message Header*
type Header struct {
	// *协议的版本号*
	Version *int32 `protobuf:"varint,1,opt,name=version,def=1" json:"version,omitempty"`
	// *binlog/redolog 文件名*
	LogfileName *string `protobuf:"bytes,2,opt,name=logfileName" json:"logfileName,omitempty"`
	// *binlog/redolog 文件的偏移位置*
	LogfileOffset *int64 `protobuf:"varint,3,opt,name=logfileOffset" json:"logfileOffset,omitempty"`
	// *服务端serverId*
	ServerId *int64 `protobuf:"varint,4,opt,name=serverId" json:"serverId,omitempty"`
	// * 变更数据的编码 *
	ServerenCode *string `protobuf:"bytes,5,opt,name=serverenCode" json:"serverenCode,omitempty"`
	// *变更数据的执行时间 *
	ExecuteTime *int64 `protobuf:"varint,6,opt,name=executeTime" json:"executeTime,omitempty"`
	// * 变更数据的来源*
	SourceType *Type `protobuf:"varint,7,opt,name=sourceType,enum=protocol.Type,def=2" json:"sourceType,omitempty"`
	// * 变更数据的schemaname*
	SchemaName *string `protobuf:"bytes,8,opt,name=schemaName" json:"schemaName,omitempty"`
	// *变更数据的tablename*
	TableName *string `protobuf:"bytes,9,opt,name=tableName" json:"tableName,omitempty"`
	// *每个event的长度*
	EventLength *int64 `protobuf:"varint,10,opt,name=eventLength" json:"eventLength,omitempty"`
	// *数据变更类型*
	EventType *EventType `protobuf:"varint,11,opt,name=eventType,enum=protocol.EventType,def=2" json:"eventType,omitempty"`
	// *预留扩展*
	Props            []*Pair `protobuf:"bytes,12,rep,name=props" json:"props,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Header) Reset() { *m = Header{} }

//func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage() {}

const Default_Header_Version int32 = 1
const Default_Header_SourceType Type = Type_MYSQL
const Default_Header_EventType EventType = EventType_UPDATE

func (m *Header) GetVersion() int32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return Default_Header_Version
}

func (m *Header) SetVersion(version int32) {
	if nil == m.Version {
		m.Version = new(int32)
	}
	*m.Version = version
}

func (m *Header) GetLogfileName() string {
	if m != nil && m.LogfileName != nil {
		return *m.LogfileName
	}
	return ""
}

func (m *Header) SetLogFileName(logfileName string) {
	if nil == m.LogfileName {
		m.LogfileName = new(string)
	}
	*m.LogfileName = logfileName
}

func (m *Header) GetLogfileOffset() int64 {
	if m != nil && m.LogfileOffset != nil {
		return *m.LogfileOffset
	}
	return 0
}

func (m *Header) SetLogfileOffset(logfileOffset int64) {
	if nil == m.LogfileOffset {
		m.LogfileOffset = new(int64)
	}
	*m.LogfileOffset = logfileOffset
}

func (m *Header) GetServerId() int64 {
	if m != nil && m.ServerId != nil {
		return *m.ServerId
	}
	return 0
}

func (m *Header) SetServerId(serverId int64) {
	if nil == m.ServerId {
		m.ServerId = new(int64)
	}
	*m.ServerId = serverId
}

func (m *Header) GetServerenCode() string {
	if m != nil && m.ServerenCode != nil {
		return *m.ServerenCode
	}
	return ""
}

func (m *Header) SetServerCode(serverCode string) {
	if nil == m.ServerenCode {
		m.ServerenCode = new(string)
	}
	*m.ServerenCode = serverCode
}

func (m *Header) GetExecuteTime() int64 {
	if m != nil && m.ExecuteTime != nil {
		return *m.ExecuteTime
	}
	return 0
}

func (m *Header) SetExecuteTime(executeTime int64) {
	if nil == m.ExecuteTime {
		m.ExecuteTime = new(int64)
	}
	*m.ExecuteTime = executeTime
}

func (m *Header) GetSourceType() Type {
	if m != nil && m.SourceType != nil {
		return *m.SourceType
	}
	return Default_Header_SourceType
}

func (m *Header) SetSourceType(sourceType Type) {
	if nil == m.SourceType {
		m.SourceType = new(Type)
	}
	*m.SourceType = sourceType
}

func (m *Header) GetSchemaName() string {
	if m != nil && m.SchemaName != nil {
		return *m.SchemaName
	}
	return ""
}

func (m *Header) SetSchemaName(schemaName string) {
	if nil == m.SchemaName {
		m.SchemaName = new(string)
	}
	*m.SchemaName = schemaName
}

func (m *Header) GetTableName() string {
	if m != nil && m.TableName != nil {
		return *m.TableName
	}
	return ""
}

func (m *Header) SetTableName(tableName string) {
	if nil == m.TableName {
		m.TableName = new(string)
	}
	*m.TableName = tableName
}

func (m *Header) GetEventLength() int64 {
	if m != nil && m.EventLength != nil {
		return *m.EventLength
	}
	return 0
}

func (m *Header) SetEventLength(eventLen int64) {
	if nil == m.EventLength {
		m.EventLength = new(int64)
	}
	*m.EventLength = eventLen
}

func (m *Header) GetEventType() EventType {
	if m != nil && m.EventType != nil {
		return *m.EventType
	}
	return Default_Header_EventType
}

func (m *Header) SetEventType(eventType EventType) {
	if nil == m.EventType {
		m.EventType = new(EventType)
	}
	*m.EventType = eventType
}

func (m *Header) GetProps() []*Pair {
	if m != nil {
		return m.Props
	}
	return nil
}

// *每个字段的数据结构*
type Column struct {
	// *字段下标*
	Index *int32 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	// *字段java中类型*
	SqlType *int32 `protobuf:"varint,2,opt,name=sqlType" json:"sqlType,omitempty"`
	// *字段名称(忽略大小写)，在mysql中是没有的*
	Name *string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// *是否是主键*
	IsKey *bool `protobuf:"varint,4,opt,name=isKey" json:"isKey,omitempty"`
	// *如果EventType=UPDATE,用于标识这个字段值是否有修改*
	Updated *bool `protobuf:"varint,5,opt,name=updated" json:"updated,omitempty"`
	// * 标识是否为空  *
	IsNull *bool `protobuf:"varint,6,opt,name=isNull,def=0" json:"isNull,omitempty"`
	// *预留扩展*
	Props []*Pair `protobuf:"bytes,7,rep,name=props" json:"props,omitempty"`
	// * 字段值,timestamp,Datetime是一个时间格式的文本 *
	Value *string `protobuf:"bytes,8,opt,name=value" json:"value,omitempty"`
	// * 对应数据对象原始长度 *
	Length *int32 `protobuf:"varint,9,opt,name=length" json:"length,omitempty"`
	// *字段mysql类型*
	MysqlType        *string `protobuf:"bytes,10,opt,name=mysqlType" json:"mysqlType,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Column) Reset()         { *m = Column{} }
func (m *Column) String() string { return proto.CompactTextString(m) }
func (*Column) ProtoMessage()    {}

const Default_Column_IsNull bool = false

func (m *Column) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *Column) SetIndex(index int32) {
	if nil == m.Index {
		m.Index = new(int32)
	}
	*m.Index = index
}

func (m *Column) GetSqlType() int32 {
	if m != nil && m.SqlType != nil {
		return *m.SqlType
	}
	return 0
}

func (m *Column) SetSqlType(sqlType int32) {
	if nil == m.SqlType {
		m.SqlType = new(int32)
	}
	*m.SqlType = sqlType
}

func (m *Column) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Column) SetName(name string) {
	if nil == m.Name {
		m.Name = new(string)
	}
	*m.Name = name
}

func (m *Column) GetIsKey() bool {
	if m != nil && m.IsKey != nil {
		return *m.IsKey
	}
	return false
}

func (m *Column) SetIsKey(isKey bool) {
	if nil == m.IsKey {
		m.IsKey = new(bool)
	}
	*m.IsKey = isKey
}

func (m *Column) GetUpdated() bool {
	if m != nil && m.Updated != nil {
		return *m.Updated
	}
	return false
}

func (m *Column) SetUpdated(updated bool) {
	if nil == m.Updated {
		m.Updated = new(bool)
	}
	*m.Updated = updated
}

func (m *Column) GetIsNull() bool {
	if m != nil && m.IsNull != nil {
		return *m.IsNull
	}
	return Default_Column_IsNull
}

func (m *Column) SetIsNull(isNull bool) {
	if nil == m.IsNull {
		m.IsNull = new(bool)
	}
	*m.IsNull = isNull
}

func (m *Column) GetProps() []*Pair {
	if m != nil {
		return m.Props
	}
	return nil
}

func (m *Column) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *Column) SetValue(value string) {
	if "" == value {
		m.Value = nil
		return
	}

	if nil == m.Value {
		m.Value = new(string)
	}
	*m.Value = value
}

func (m *Column) GetLength() int32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *Column) SetLength(length int32) {
	if nil == m.Length {
		m.Length = new(int32)
	}
	*m.Length = length
}

func (m *Column) GetMysqlType() string {
	if m != nil && m.MysqlType != nil {
		return *m.MysqlType
	}
	return ""
}

func (m *Column) SetMysqlType(mysqlType string) {
	if nil == m.MysqlType {
		m.MysqlType = new(string)
	}
	*m.MysqlType = mysqlType
}

type RowData struct {
	// * 字段信息，增量数据(修改前,删除前) *
	BeforeColumns []*Column `protobuf:"bytes,1,rep,name=beforeColumns" json:"beforeColumns,omitempty"`
	// * 字段信息，增量数据(修改后,新增后)  *
	AfterColumns []*Column `protobuf:"bytes,2,rep,name=afterColumns" json:"afterColumns,omitempty"`
	// *预留扩展*
	Props            []*Pair `protobuf:"bytes,3,rep,name=props" json:"props,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RowData) Reset()         { *m = RowData{} }
func (m *RowData) String() string { return proto.CompactTextString(m) }
func (*RowData) ProtoMessage()    {}

func (m *RowData) GetBeforeColumns() []*Column {
	if m != nil {
		return m.BeforeColumns
	}
	return nil
}

func (m *RowData) GetAfterColumns() []*Column {
	if m != nil {
		return m.AfterColumns
	}
	return nil
}

func (m *RowData) GetProps() []*Pair {
	if m != nil {
		return m.Props
	}
	return nil
}

// *message row 每行变更数据的数据结构*
type RowChange struct {
	// *tableId,由数据库产生*
	TableId *int64 `protobuf:"varint,1,opt,name=tableId" json:"tableId,omitempty"`
	// *数据变更类型*
	EventType *EventType `protobuf:"varint,2,opt,name=eventType,enum=protocol.EventType,def=2" json:"eventType,omitempty"`
	// * 标识是否是ddl语句  *
	IsDdl *bool `protobuf:"varint,10,opt,name=isDdl,def=0" json:"isDdl,omitempty"`
	// * ddl/query的sql语句  *
	Sql *string `protobuf:"bytes,11,opt,name=sql" json:"sql,omitempty"`
	// * 一次数据库变更可能存在多行  *
	RowDatas []*RowData `protobuf:"bytes,12,rep,name=rowDatas" json:"rowDatas,omitempty"`
	// *预留扩展*
	Props []*Pair `protobuf:"bytes,13,rep,name=props" json:"props,omitempty"`
	// * ddl/query的schemaName，会存在跨库ddl，需要保留执行ddl的当前schemaName  *
	DdlSchemaName    *string `protobuf:"bytes,14,opt,name=ddlSchemaName" json:"ddlSchemaName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RowChange) Reset()         { *m = RowChange{} }
func (m *RowChange) String() string { return proto.CompactTextString(m) }
func (*RowChange) ProtoMessage()    {}

const Default_RowChange_EventType EventType = EventType_UPDATE
const Default_RowChange_IsDdl bool = false

func (m *RowChange) GetTableId() int64 {
	if m != nil && m.TableId != nil {
		return *m.TableId
	}
	return 0
}

func (m *RowChange) SetTableId(tableId int64) {
	if nil == m.TableId {
		m.TableId = new(int64)
	}
	*m.TableId = tableId
}

func (m *RowChange) GetEventType() EventType {
	if m != nil && m.EventType != nil {
		return *m.EventType
	}
	return Default_RowChange_EventType
}

func (m *RowChange) SetEventType(eventType EventType) {
	if nil == m.EventType {
		m.EventType = new(EventType)
	}
	*m.EventType = eventType
}

func (m *RowChange) GetIsDdl() bool {
	if m != nil && m.IsDdl != nil {
		return *m.IsDdl
	}
	return Default_RowChange_IsDdl
}

func (m *RowChange) SetIsDdl(isDdl bool) {
	if nil == m.IsDdl {
		m.IsDdl = new(bool)
	}
	*m.IsDdl = isDdl
}

func (m *RowChange) GetSql() string {
	if m != nil && m.Sql != nil {
		return *m.Sql
	}
	return ""
}

func (m *RowChange) SetSql(sql string) {
	if nil == m.Sql {
		m.Sql = new(string)
	}
	*m.Sql = sql
}

func (m *RowChange) GetRowDatas() []*RowData {
	if m != nil {
		return m.RowDatas
	}
	return nil
}

func (m *RowChange) SetRowDatas(rowDatas []*RowData) {
	m.RowDatas = rowDatas
}

func (m *RowChange) GetProps() []*Pair {
	if m != nil {
		return m.Props
	}
	return nil
}

func (m *RowChange) GetDdlSchemaName() string {
	if m != nil && m.DdlSchemaName != nil {
		return *m.DdlSchemaName
	}
	return ""
}

// *开始事务的一些信息*
type TransactionBegin struct {
	// *已废弃，请使用header里的executeTime*
	ExecuteTime *int64 `protobuf:"varint,1,opt,name=executeTime" json:"executeTime,omitempty"`
	// *已废弃，Begin里不提供事务id*
	TransactionId *string `protobuf:"bytes,2,opt,name=transactionId" json:"transactionId,omitempty"`
	// *预留扩展*
	Props []*Pair `protobuf:"bytes,3,rep,name=props" json:"props,omitempty"`
	// *执行的thread Id*
	ThreadId         *int64 `protobuf:"varint,4,opt,name=threadId" json:"threadId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TransactionBegin) Reset()         { *m = TransactionBegin{} }
func (m *TransactionBegin) String() string { return proto.CompactTextString(m) }
func (*TransactionBegin) ProtoMessage()    {}

func (m *TransactionBegin) GetExecuteTime() int64 {
	if m != nil && m.ExecuteTime != nil {
		return *m.ExecuteTime
	}
	return 0
}

func (m *TransactionBegin) SetExecuteTime(executeTime int64) {
	if nil == m.ExecuteTime {
		m.ExecuteTime = new(int64)
	}
	*m.ExecuteTime = executeTime
}

func (m *TransactionBegin) GetTransactionId() string {
	if m != nil && m.TransactionId != nil {
		return *m.TransactionId
	}
	return ""
}

func (m *TransactionBegin) SetTransactionId(transactionId string) {
	if nil == m.TransactionId {
		m.TransactionId = new(string)
	}
	*m.TransactionId = transactionId
}

func (m *TransactionBegin) GetProps() []*Pair {
	if m != nil {
		return m.Props
	}
	return nil
}

func (m *TransactionBegin) GetThreadId() int64 {
	if m != nil && m.ThreadId != nil {
		return *m.ThreadId
	}
	return 0
}

func (m *TransactionBegin) SetThreadId(threadId int64) {
	if nil == m.ThreadId {
		m.ThreadId = new(int64)
	}
	*m.ThreadId = threadId
}

// *结束事务的一些信息*
type TransactionEnd struct {
	// *已废弃，请使用header里的executeTime*
	ExecuteTime *int64 `protobuf:"varint,1,opt,name=executeTime" json:"executeTime,omitempty"`
	// *事务号*
	TransactionId *string `protobuf:"bytes,2,opt,name=transactionId" json:"transactionId,omitempty"`
	// *预留扩展*
	Props            []*Pair `protobuf:"bytes,3,rep,name=props" json:"props,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TransactionEnd) Reset()         { *m = TransactionEnd{} }
func (m *TransactionEnd) String() string { return proto.CompactTextString(m) }
func (*TransactionEnd) ProtoMessage()    {}

func (m *TransactionEnd) GetExecuteTime() int64 {
	if m != nil && m.ExecuteTime != nil {
		return *m.ExecuteTime
	}
	return 0
}

func (m *TransactionEnd) SetExecuteTime(executeTime int64) {
	if nil == m.ExecuteTime {
		m.ExecuteTime = new(int64)
	}
	*m.ExecuteTime = executeTime
}

func (m *TransactionEnd) GetTransactionId() string {
	if m != nil && m.TransactionId != nil {
		return *m.TransactionId
	}
	return ""
}

func (m *TransactionEnd) SetTransactionId(transactionId string) {

	if nil == m.TransactionId {
		m.TransactionId = new(string)
	}
	*m.TransactionId = transactionId
}

func (m *TransactionEnd) GetProps() []*Pair {
	if m != nil {
		return m.Props
	}
	return nil
}

// *预留扩展*
type Pair struct {
	Key              *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}

func (m *Pair) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Pair) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
	proto.RegisterEnum("protocol.EntryType", EntryType_name, EntryType_value)
	proto.RegisterEnum("protocol.EventType", EventType_name, EventType_value)
	proto.RegisterEnum("protocol.Type", Type_name, Type_value)
}