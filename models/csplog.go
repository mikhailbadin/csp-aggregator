package models

import (
	"fmt"
	"log"
	"time"

	"github.com/globalsign/mgo"

	"github.com/mikhailbadin/csp-aggregator/db"
	"github.com/mikhailbadin/csp-aggregator/forms"
)

// CSPLogReport CSP report
type CSPLogReport struct {
	/* Basic directives */
	BlockedURI         string `bson:"blocked-uri" json:"blocked-uri"`
	Disposition        string `bson:"disposition" json:"disposition"`
	DocumentURI        string `bson:"document-uri" json:"document-uri"`
	Referrer           string `bson:"referrer" json:"referrer"`
	ViolatedDirective  string `bson:"violated-directive" json:"violated-directive"`
	EffectiveDirective string `bson:"effective-directive" json:"effective-directive"`
	OriginalPolicy     string `bson:"original-policy" json:"original-policy"`
	StatusCode         int    `bson:"status-code" json:"status-code"`
	/* Sctipt directives */
	SourceFile   string `bson:"source-file" json:"source-file"`
	LineNumber   int    `bson:"line-number" json:"line-number"`
	ColumnNumber int    `bson:"column-number" json:"column-number"`
}

// CSPLog format for logging
type CSPLog struct {
	Time       int64         `bson:"report_time" json:"report_time"`
	ReportOnly bool          `bson:"report_only" json:"report_only"`
	Report     *CSPLogReport `bson:"report" json:"report"`
}

// WriteCSPLog write CSP report to MongoDB
func WriteCSPLog(r *forms.CSPReport, h *forms.Headers, reportOnly bool, time time.Time) error {
	l := CSPLog{
		Time:       time.UnixNano(),
		ReportOnly: reportOnly,
		Report: &CSPLogReport{
			BlockedURI:         r.BlockedURI,
			Disposition:        r.Disposition,
			DocumentURI:        r.DocumentURI,
			Referrer:           r.Referrer,
			ViolatedDirective:  r.ViolatedDirective,
			EffectiveDirective: r.EffectiveDirective,
			OriginalPolicy:     r.OriginalPolicy,
			StatusCode:         r.StatusCode,
			SourceFile:         r.SourceFile,
			LineNumber:         r.LineNumber,
			ColumnNumber:       r.ColumnNumber,
		},
	}
	s := db.GetMongoDB()
	c := getMongoCollection(s)
	if err := c.Insert(&l); err != nil {
		if errPing := s.Ping(); errPing != nil {
			log.Fatalln("lost connection to MongoDB:", errPing.Error())
		}
		return fmt.Errorf("cannot insert report %+v: %s", l, err.Error())
	}
	return nil
}

func getMongoCollection(s *mgo.Session) *mgo.Collection {
	return s.DB(db.MongoDBName).C(db.MongoCollectionName)
}
