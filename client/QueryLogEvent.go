package client

import (
	"time"

	. "github.com/SDHM/sqlregret/binlogevent"
	"github.com/SDHM/sqlregret/mysql"
)

type QueryLogEvent struct {
	sessionId int64
	execTime  int64
	errCode   int32
	query     string
	dbName    string
}

func ParseQueryLogEvent(logbuf *mysql.LogBuffer,
	descriptionEvent *FormatDescriptionLogEvent) *QueryLogEvent {

	this := new(QueryLogEvent)
	this.sessionId = int64(logbuf.GetUInt32())
	this.execTime = int64(logbuf.GetUInt32())
	schema_length := logbuf.GetUInt8()
	//error_code
	this.errCode = int32(logbuf.GetUInt16())

	postHeaderLen := descriptionEvent.PostHeaderLen[QUERY_EVENT-1]
	if postHeaderLen > QUERY_HEADER_MINIMAL_LEN {
		status_vars_length := logbuf.GetUInt16()
		//status_vars
		logbuf.SkipLen(status_vars_length)
	}

	//schema
	this.dbName = logbuf.GetVarLenString(schema_length)
	logbuf.SkipLen(1)

	this.query = logbuf.GetRestString()

	return this
}

func (this *QueryLogEvent) GetQuery() string {
	return this.query
}

func (this *QueryLogEvent) GetSessionId() int64 {
	return this.sessionId
}

func (this *QueryLogEvent) GetSchema() string {
	return this.dbName
}

func (this *QueryLogEvent) GetTime() string {
	timeSnap := time.Unix(this.execTime, 0)
	return timeSnap.Format("2006-01-02 15:04:05")
}
