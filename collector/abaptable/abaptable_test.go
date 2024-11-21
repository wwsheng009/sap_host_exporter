package abaptable

import (
	"strings"
	"testing"

	"github.com/SUSE/sap_host_exporter/lib/sapcontrol"
	"github.com/SUSE/sap_host_exporter/test/mock_sapcontrol"
	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

func TestNewCollector(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWebService := mock_sapcontrol.NewMockWebService(ctrl)

	_, err := NewCollector(mockWebService)

	assert.Nil(t, err)
}

func TestWorkProcessesMetric(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWebService := mock_sapcontrol.NewMockWebService(ctrl)
	mockWebService.EXPECT().GetABAPWPTable().Return(&sapcontrol.GetABAPWPTableResponse{
		Workprocess: &sapcontrol.ArrayOfWorkProcess{
			Item: []*sapcontrol.WorkProcess{
				{
					No:      16,
					Typ:     "DIA",
					Status:  "Run",
					Reason:  "",
					Start:   "yes",
					Err:     "",
					Sem:     "",
					Cpu:     "33891",
					Pid:     18585,
					Program: "Z_ENDLESS_LOOP",
					Client:  "300",
					User:    "WANGWS",
					Action:  "",
					Table:   "",
				},
			},
		},
	}, nil)
	// mockWebService.EXPECT().GetSystemInstanceList().Return(&sapcontrol.GetSystemInstanceListResponse{}, nil)
	mockWebService.EXPECT().GetCurrentInstance().Return(&sapcontrol.CurrentSapInstance{
		SID:      "S4H",
		Number:   0,
		Name:     "D00",
		Hostname: "sap1809demo",
	}, nil).AnyTimes()

	expectedMetrics := `
	# HELP sap_start_service_processes The processes started by the SAP Start Service
	# TYPE sap_start_service_processes gauge
	sap_abap_work_processes_dia{SID="S4H",action="",client="300",cpu="33891",err="",instance_hostname="sap1809demo",instance_name="D00",instance_number="0",no="16",pid="18585",program="Z_ENDLESS_LOOP",reason="",sem="",start="yes",status="Run",table="",typ="DIA",user="WANGWS"} 430
	`

	var err error
	collector, err := NewCollector(mockWebService)
	assert.NoError(t, err)

	err = testutil.CollectAndCompare(collector, strings.NewReader(expectedMetrics), "sap_start_service_processes")
	assert.NoError(t, err)
}
