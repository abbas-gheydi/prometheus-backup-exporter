package prommetrics

import "fmt"
//Generate func formats data to prometheus understand it
func Generate(metricName string,metricPrefix string, key string, value string) (help string, metric string) {
        //#TYPE backup gauge
        //backup{backup="_home_ubuntu_bk_tar_backup_tar"} 1.92706e+07
	metric = fmt.Sprint(metricName, "{", metricPrefix, "=", "\"", key, "\"", "}", " ", value,"\n")

	help = fmt.Sprint("#TYPE ", metricName, " ", "gauge","\n")

	return

}
