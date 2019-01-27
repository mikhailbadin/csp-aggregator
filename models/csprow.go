package models

import (
	"fmt"
	"log"
	"time"

	"github.com/mikhailbadin/csp-aggregator/db"
	"github.com/mikhailbadin/csp-aggregator/forms"
	msgpack "gopkg.in/vmihailenco/msgpack.v2"
)

// CSPRow CSP struct saved in Tarantool
type CSPRow struct {
	Time         uint64
	ReportOnly   bool
	BlockedURI   string
	DocumentURI  string
	SourceFile   string
	LineNumber   int
	ColumnNumber int
	Referer      string
}

// EncodeMsgpack for encode msgpack
func (r *CSPRow) EncodeMsgpack(e *msgpack.Encoder) error {
	if err := e.EncodeSliceLen(9); err != nil {
		return err
	}
	if err := e.EncodeNil(); err != nil {
		return err
	}
	if err := e.EncodeUint64(r.Time); err != nil {
		return err
	}
	if err := e.EncodeBool(r.ReportOnly); err != nil {
		return err
	}
	if err := e.EncodeString(r.BlockedURI); err != nil {
		return err
	}
	if err := e.EncodeString(r.DocumentURI); err != nil {
		return err
	}
	if err := e.EncodeString(r.SourceFile); err != nil {
		return err
	}
	if err := e.EncodeInt(r.LineNumber); err != nil {
		return err
	}
	if err := e.EncodeInt(r.ColumnNumber); err != nil {
		return err
	}
	if err := e.EncodeString(r.Referer); err != nil {
		return err
	}
	return nil
}

// DecodeMsgpack for encode msgpack
func (r *CSPRow) DecodeMsgpack(d *msgpack.Decoder) error {
	return nil
}

// WriteCSPRow write CSP report to Tarantool
func WriteCSPRow(r *forms.CSPReport, h *forms.Headers, reportOnly bool, time time.Time) error {
	row := CSPRow{
		Time:         uint64(time.UnixNano()),
		ReportOnly:   reportOnly,
		BlockedURI:   r.BlockedURI,
		DocumentURI:  r.DocumentURI,
		SourceFile:   r.SourceFile,
		LineNumber:   r.LineNumber,
		ColumnNumber: r.ColumnNumber,
		Referer:      h.Referer,
	}
	directive := r.ViolatedDirective
	switch directive {
	case "script-src":
		return writeScriptSrcRow(&row)
	default:
	}
	return nil
}

func writeScriptSrcRow(row *CSPRow) error {
	c := db.GetTarantoolDB()
	if _, err := c.Insert(db.SpaceScriptSrc, row); err != nil {
		if _, errPing := c.Ping(); errPing != nil {
			log.Fatalln("lost connection to TarantoolDB:", errPing.Error())
		}
		return fmt.Errorf("cannot insert report %+v: %s", row, err.Error())
	}
	return nil
}
