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
	var confirmed []opts.LineData
	var recovered []opts.LineData
	var deaths []opts.LineData

	for a, r := range statistics {
		axis = append(axis, a)
		confirmed = append(confirmed, opts.LineData{Value: r.Confirmed})
		recovered = append(recovered, opts.LineData{Value: r.Recovered})
		deaths = append(deaths, opts.LineData{Value: r.Deaths})

	}

	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Covid19 Reported Cases in Islamic Republic of IRAN",
			Subtitle: "api.covid19api.com",
		}))

	line.SetXAxis(axis).
		AddSeries("Confirmed Cases", confirmed).
		AddSeries("Recovered Cases", recovered).
		AddSeries("Death Cases", deaths).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func main() {
	fmt.Println("Welcome To Covid19 Report WebApplication")
	fmt.Println("http://localhost:8081")
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)
}
