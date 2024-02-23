package timeseries

import (
	"context"
	"fmt"
	"sokwva/acfun/billboard/common"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

var tsdbHandle influxdb2.Client

func InitClient() error {
	tsdbHandle = influxdb2.NewClient("http://"+common.ConfHandle.Persist.SvrAddr+":"+common.ConfHandle.Persist.SvrPort, common.ConfHandle.Persist.SvrApiKey)
	var ctx context.Context = context.Background()
	if _, err := tsdbHandle.Health(ctx); err != nil {
		return err
	}
	return nil
}

func SaveTSRecord(tags map[string]string, fields map[string]interface{}) {
	writeAPI := tsdbHandle.WriteAPIBlocking(common.ConfHandle.Persist.OrgName, common.ConfHandle.Persist.StorageBucket)
	point := write.NewPoint("dougaInfo", tags, fields, time.Now())

	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		common.Log.Error(err.Error())
	}
}

func TestWrite() {
	writeAPI := tsdbHandle.WriteAPIBlocking(common.ConfHandle.Persist.OrgName, common.ConfHandle.Persist.StorageBucket)
	for value := 0; value < 5; value++ {
		fmt.Println("field: ", value)
		tags := map[string]string{
			"tagname1": "tagvalue1",
		}
		fields := map[string]interface{}{
			"field1": value,
			"field2": value,
			"field3": value,
		}
		point := write.NewPoint("measurement1", tags, fields, time.Now())
		time.Sleep(1 * time.Second) // separate points by 1 second

		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			common.Log.Error(err.Error())
		}
	}
}

func TestRead() {
	queryAPI := tsdbHandle.QueryAPI(common.ConfHandle.Persist.OrgName)
	query := `from(bucket: "Test0")
            |> range(start: -10m)
            |> filter(fn: (r) => r._measurement == "measurement1")`
	results, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		common.Log.Error(err.Error())
	}
	for results.Next() {
		fmt.Println(results.Record())
	}
	if err := results.Err(); err != nil {
		common.Log.Error(err.Error())
	}
	//Aggregate Query
	query = `from(bucket: "Test0")
	|> range(start: -10m)
	|> filter(fn: (r) => r._measurement == "measurement1")
	|> mean()`
	results, err = queryAPI.Query(context.Background(), query)
	if err != nil {
		common.Log.Error(err.Error())
	}
	for results.Next() {
		fmt.Println(results.Record())
	}
	if err := results.Err(); err != nil {
		common.Log.Error(err.Error())
	}
}
