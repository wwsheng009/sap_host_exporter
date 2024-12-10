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

type ABAPProcessInfo struct {
	dia_count int
	dia_free  int
	spo_count int
	spo_free  int
	btc_count int
	btc_free  int
	upd_count int
	upd_free  int
	up2_count int
	up2_free  int
}
type abaptableServerCollector struct {
	collector.DefaultCollector
	webService sapcontrol.WebService
}

func NewCollector(webService sapcontrol.WebService) (*abaptableServerCollector, error) {

	c := &abaptableServerCollector{
		collector.NewDefaultCollector("abap_work"),
		webService,
	}

	c.SetDescriptor("processes_dia", "The abap dia work processes", []string{"no", "typ", "status", "reason", "start", "err", "sem", "cpu", "pid", "program", "client", "user", "action", "table", "instance_name", "instance_number", "SID", "instance_hostname"})
	c.SetDescriptor("processes_btc", "The abap btc work processes", []string{"no", "typ", "status", "reason", "start", "err", "sem", "cpu", "pid", "program", "client", "user", "action", "table", "instance_name", "instance_number", "SID", "instance_hostname"})
	c.SetDescriptor("processes_spo", "The abap spo work processes", []string{"no", "typ", "status", "reason", "start", "err", "sem", "cpu", "pid", "program", "client", "user", "action", "table", "instance_name", "instance_number", "SID", "instance_hostname"})
	c.SetDescriptor("processes_upd", "The abap upd work processes", []string{"no", "typ", "status", "reason", "start", "err", "sem", "cpu", "pid", "program", "client", "user", "action", "table", "instance_name", "instance_number", "SID", "instance_hostname"})
	c.SetDescriptor("processes_up2", "The abap up2 work processes", []string{"no", "typ", "status", "reason", "start", "err", "sem", "cpu", "pid", "program", "client", "user", "action", "table", "instance_name", "instance_number", "SID", "instance_hostname"})
	c.SetDescriptor("processes_wait_pct", "The abap dia work processes free percent", []string{"typ", "instance_name", "instance_number", "SID", "instance_hostname"})
	c.SetDescriptor("processes_simple", "The abap work processes", []string{"no", "typ", "pid", "instance_name", "instance_number", "SID", "instance_hostname"})

	return c, nil
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

	info := ABAPProcessInfo{}
	for _, process := range abapStatistic.Workprocess.Item {
		no := strconv.Itoa(int(process.No))
		pid := strconv.Itoa(int(process.Pid))

		labels := append([]string{no, process.Typ, process.Status, process.Reason, process.Start, process.Err, process.Sem,
			timeToSeconds(process.Cpu),
			pid, process.Program, process.Client, process.User, process.Action, process.Table}, commonLabels...)

		// label的变化会导致 指标的变化， 所以需要单独增加一条简化的指标，删除一些会变化的字段
		labels_simple := append([]string{no, process.Typ, pid}, commonLabels...)

		floatValue, err := strconv.ParseFloat(process.Time, 64)
		if err != nil {

			floatValue = 0
		}
		switch process.Typ {
		case "DIA":
			info.dia_count++
			if process.Status == "Wait" {
				info.dia_free++
			}
			ch <- c.MakeCounterMetric("processes_dia", floatValue, labels...)
		case "BTC":
			info.btc_count++
			if process.Status == "Wait" {
				info.btc_free++
			}
			ch <- c.MakeCounterMetric("processes_btc", floatValue, labels...)
		case "SPO":
			info.spo_count++
			if process.Status == "Wait" {
				info.spo_free++
			}
			ch <- c.MakeCounterMetric("processes_spo", floatValue, labels...)
		case "UPD":
			info.upd_count++
			if process.Status == "Wait" {
				info.upd_free++
			}
			ch <- c.MakeCounterMetric("processes_upd", floatValue, labels...)
		case "UP2":
			info.up2_count++
			if process.Status == "Wait" {
				info.up2_free++
			}
			ch <- c.MakeCounterMetric("processes_up2", floatValue, labels...)
		default:
		}
		ch <- c.MakeCounterMetric("processes_simple", floatValue, labels_simple...)

	}
	if info.dia_count > 0 {
		ch <- c.MakeGaugeMetric("processes_wait_pct", float64(info.dia_free)/float64(info.dia_count)*100, append([]string{"ABAP/DIA"}, commonLabels...)...)
	}
	if info.spo_count > 0 {
		ch <- c.MakeGaugeMetric("processes_wait_pct", float64(info.spo_free)/float64(info.spo_count)*100, append([]string{"ABAP/SPO"}, commonLabels...)...)
	}
	if info.btc_count > 0 {
		ch <- c.MakeGaugeMetric("processes_wait_pct", float64(info.btc_free)/float64(info.btc_count)*100, append([]string{"ABAP/BTC"}, commonLabels...)...)
	}
	if info.upd_count > 0 {
		ch <- c.MakeGaugeMetric("processes_wait_pct", float64(info.upd_free)/float64(info.upd_count)*100, append([]string{"ABAP/UPD"}, commonLabels...)...)
	}
	if info.up2_count > 0 {
		ch <- c.MakeGaugeMetric("processes_wait_pct", float64(info.up2_free)/float64(info.up2_count)*100, append([]string{"ABAP/UP2"}, commonLabels...)...)
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
