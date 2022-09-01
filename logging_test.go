package go_gcp_logging_easycall

import "testing"

func TestLog(t *testing.T) {
	Log("innovation-dev-321010", "test-api", &Logger{}, "", "api", map[string]string{"session_id": "001"})
}
