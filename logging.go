package go_gcp_logging_easycall

import (
	"context"
	"log"

	"cloud.google.com/go/logging"
	"google.golang.org/genproto/googleapis/api/monitoredres"
)

type Logger struct {
	User         string      `json:"token"`
	Message      string      `json:"message"`
	Reason       string      `json:"reason"`
	RemoteIp     string      `json:"ipaddress"`
	Method       string      `json:"method"`
	Urlpath      string      `json:"route"`
	StatusCode   int         `json:"status_code"`
	RequestData  interface{} `json:"request_data"`
	ResponseData interface{} `json:"response_data"`
	SessionId    string      `json:"sessid"`
}

func Log(project string, loggerName string, message interface{}, severity string, resourceName string, resourceLabels map[string]string) {
	var severityTag logging.Severity

	ctx := context.Background()
	client, err := logging.NewClient(ctx, project)
	if err != nil {
		log.Fatalf("Failed to create Google Cloud Logging client: %v", err)
	}
	defer client.Close()

	switch severity {
	case "alert":
		severityTag = logging.Alert
	case "critical":
		severityTag = logging.Critical
	case "debug":
		severityTag = logging.Debug
	case "emergency":
		severityTag = logging.Emergency
	case "error":
		severityTag = logging.Error
	case "info":
		severityTag = logging.Info
	case "notice":
		severityTag = logging.Notice
	case "warning":
		severityTag = logging.Warning
	default:
		severityTag = logging.Default
	}

	client.Logger(loggerName).Log(logging.Entry{
		Payload:  message,
		Severity: severityTag,
		Resource: &monitoredres.MonitoredResource{
			Type:   resourceName,
			Labels: resourceLabels,
		},
	})
}
