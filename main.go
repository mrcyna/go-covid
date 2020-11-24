package main

import (
	"fmt"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/mrcyna/go-covid/data"
)

func httpserver(w http.ResponseWriter, _ *http.Request) {

	statistics, err := data.CovidData("IRAN", 30)
	if err != nil {
		fmt.Printf("Oops! Something goes wrong, %v", err)
	}

	var axis []string
	var cases []opts.LineData

	for a, c := range statistics {
		axis = append(axis, a)
		cases = append(cases, opts.LineData{Value: c})
	}

	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Covid19 Reported Cases in Islamic Republic of IRAN",
			Subtitle: "api.covid19api.com",
		}))

	line.SetXAxis(axis).
		AddSeries("Cases", cases).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func main() {
	fmt.Println("Welcome To Covid19 Report WebApplication")
	fmt.Println("http://localhost:8081")
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)
}
