package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/urfave/cli/v2"
)

var defaultHttpClient = &http.Client{Timeout: 10 * time.Second}

type Stat struct {
	ID                 int     `json:"id"`
	DownloadCount      float32 `json:"downloadCount"`
	PopularityScore    float32 `json:"popularityScore"`
	GamePopularityRank float32 `json:"gamePopularityRank"`
}

func getJson(url string, target interface{}) error {
	log.Printf("Get response: url=%s", url)
	r, err := defaultHttpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getStat(url string) (*Stat, error) {
	stat := new(Stat)
	err := getJson(url, stat)
	if err != nil {
		return nil, err
	}
	return stat, nil
}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "influx-url",
				EnvVars:  []string{"FORGE_STATS_INFLUX_URL"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "influx-org",
				EnvVars: []string{"FORGE_STATS_INFLUX_ORG"},
				Value:   "default",
			},
			&cli.StringFlag{
				Name:    "influx-bucket",
				EnvVars: []string{"FORGE_STATS_INFLUX_BUCKET"},
				Value:   "default",
			},
			&cli.StringFlag{
				Name:    "status-url",
				EnvVars: []string{"FORGE_STATS_STATUS_URL"},
				Value:   "https://addons-ecs.forgesvc.net/api/v2/addon",
			},
			&cli.IntFlag{
				Name:     "mod-id",
				EnvVars:  []string{"FORGE_STATS_MOD_ID"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "token",
				EnvVars: []string{"FORGE_STATS_INFLUX_TOKEN"},
			},
			&cli.StringFlag{
				Name:    "stat-name",
				EnvVars: []string{"FORGE_STAT_NAME"},
				Value:   "default",
			},
		},
		Action: func(c *cli.Context) error {
			statusURL := c.String("status-url")
			influxURL := c.String("influx-url")
			influxOrg := c.String("influx-org")
			statName := c.String("stat-name")
			influxBucket := c.String("influx-bucket")
			token := c.String("token")
			modID := c.Int("mod-id")

			targetURL := fmt.Sprintf("%s/%d", statusURL, modID)
			stat, err := getStat(targetURL)
			if err != nil {
				log.Fatalf("Something went wrong: err=%s", err.Error())
			}

			influxClient := influxdb2.NewClient(influxURL, token)
			writeAPI := influxClient.WriteAPIBlocking(influxOrg, influxBucket)
			p := influxdb2.NewPoint("stat",
				map[string]string{"name": statName},
				map[string]interface{}{
					"downloadCount":      stat.DownloadCount,
					"popularityScore":    stat.PopularityScore,
					"gamePopularityRank": stat.GamePopularityRank},
				time.Now())
			err = writeAPI.WritePoint(context.Background(), p)
			if err != nil {
				log.Fatal(err.Error())
			}
			defer influxClient.Close()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
