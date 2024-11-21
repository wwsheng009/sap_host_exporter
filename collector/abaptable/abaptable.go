package abaptable

import (
	"strconv"
	"strings"

	"github.com/SUSE/sap_host_exporter/collector"
	"github.com/SUSE/sap_host_exporter/lib/sapcontrol"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

func NewCollector(webService sapcontrol.WebService) (*abaptableServerCollector, error) {

	c := &abaptableServerCollector{
		collector.NewDefaultCollector("abap_work"),
		webService,
	}

	c.SetDescriptor("processes", "The abap work processes started by the SAP Start Service", []string{"no", "typ", "status", "reason", "start", "err", "sem", "cpu", "pid", "program", "client", "user", "action", "table", "instance_name", "instance_number", "SID", "instance_hostname"})
	c.SetDescriptor("processes_dia", "The abap dia work processes started by the SAP Start Service", []string{"no", "typ", "status", "reason", "start", "err", "sem", "cpu", "pid", "program", "client", "user", "action", "table", "instance_name", "instance_number", "SID", "instance_hostname"})
	c.SetDescriptor("processes_btc", "The abap btc work processes started by the SAP Start Service", []string{"no", "typ", "status", "reason", "start", "err", "sem", "cpu", "pid", "program", "client", "user", "action", "table", "instance_name", "instance_number", "SID", "instance_hostname"})
	c.SetDescriptor("processes_spo", "The abap spo work processes started by the SAP Start Service", []string{"no", "typ", "status", "reason", "start", "err", "sem", "cpu", "pid", "program", "client", "user", "action", "table", "instance_name", "instance_number", "SID", "instance_hostname"})
	c.SetDescriptor("processes_upd", "The abap upd work processes started by the SAP Start Service", []string{"no", "typ", "status", "reason", "start", "err", "sem", "cpu", "pid", "program", "client", "user", "action", "table", "instance_name", "instance_number", "SID", "instance_hostname"})
	c.SetDescriptor("processes_up2", "The abap up2 work processes started by the SAP Start Service", []string{"no", "typ", "status", "reason", "start", "err", "sem", "cpu", "pid", "program", "client", "user", "action", "table", "instance_name", "instance_number", "SID", "instance_hostname"})
	return c, nil
}

type abaptableServerCollector struct {
	collector.DefaultCollector
	webService sapcontrol.WebService
}

func (c *abaptableServerCollector) Collect(ch chan<- prometheus.Metric) {
	log.Debugln("Collecting ABAPGetWPTable Server metrics")
	err := c.recordAbapTable(ch)
	if err != nil {
		log.Warnf("ABAPGetWPTable Server Collector scrape failed: %s", err)
	}
}

func (c *abaptableServerCollector) recordAbapTable(ch chan<- prometheus.Metric) error {
	abapStatistic, err := c.webService.GetABAPWPTable()
	if err != nil {
		return errors.Wrap(err, "SAPControl web service error")
	}
	currentSapInstance, err := c.webService.GetCurrentInstance()
	if err != nil {
		return errors.Wrap(err, "SAPControl web service error")
	}

	commonLabels := []string{
		currentSapInstance.Name,
		strconv.Itoa(int(currentSapInstance.Number)),
		currentSapInstance.SID,
		currentSapInstance.Hostname,
	}

	for _, process := range abapStatistic.Workprocess.Item {
		labels := append([]string{strconv.Itoa(int(process.No)), process.Typ, process.Status, process.Reason, process.Start, process.Err, process.Sem, timeToSeconds(process.Cpu), strconv.Itoa(int(process.Pid)), process.Program, process.Client, process.User, process.Action, process.Table}, commonLabels...)

		floatValue, err := strconv.ParseFloat(process.Time, 64)
		if err != nil {

			floatValue = 0
		}
		switch process.Typ {
		case "DIA":
			ch <- c.MakeCounterMetric("processes_dia", floatValue, labels...)
		case "BTC":
			ch <- c.MakeCounterMetric("processes_btc", floatValue, labels...)
		case "SPO":
			ch <- c.MakeCounterMetric("processes_spo", floatValue, labels...)
		case "UPD":
			ch <- c.MakeCounterMetric("processes_upd", floatValue, labels...)
		case "UP2":
			ch <- c.MakeCounterMetric("processes_up2", floatValue, labels...)
		default:
			ch <- c.MakeCounterMetric("processes", floatValue, labels...)
		}
	}
	return nil
}

func timeToSeconds(timeStr string) string {
	// Split the string by colon
	parts := strings.Split(timeStr, ":")
	if len(parts) != 3 {
		return timeStr
	}

	// Convert each part to an integer
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return timeStr
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return timeStr
	}

	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		return timeStr
	}

	// Calculate total seconds
	totalSeconds := hours*3600 + minutes*60 + seconds
	return strconv.Itoa(int(totalSeconds))
}
