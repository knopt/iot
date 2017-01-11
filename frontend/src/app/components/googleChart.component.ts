import { Component, OnChanges, OnInit, Input } from '@angular/core';

declare var google:any;

@Component({
  selector: 'chart',
  template: `
  <div class="four wide column center aligned">
      <div id="chart_divEvolution" style="width: 900px; height: 500px;"></div>
  </div>
`
})
export class GoogleChartComponent implements OnInit, OnChanges {
  private static googleLoaded:any;

  private options: any;
  private chart: any;
  private data: any;

  @Input() chartData: any;

  constructor(){
      console.log("Here is GoogleChartComponent")
  }

  getGoogle() {
      return google;
  }

  ngOnInit() {
    if(!GoogleChartComponent.googleLoaded) {
      GoogleChartComponent.googleLoaded = true;
      google.charts.load('current',  {packages: ['corechart', 'bar']});
    }
    google.charts.setOnLoadCallback(() => this.drawGraph());
  }

  ngOnChanges(): void {
      console.log("on changes chart");
  }

  drawGraph(){
    console.log("LineGraph Evolution...");
    this.data = this.createDataTable(this.chartData);

    this.options = {
      title: 'Statistics',
      curveType: 'function',
    };

    this.chart = this.createLineChart(document.getElementById('chart_divEvolution'));
    this.chart.draw(this.data, this.options);
  }

  createLineChart(element:any):any {
      return new google.visualization.LineChart(element);
  }

  createDataTable(array:any[]):any {
      return google.visualization.arrayToDataTable(array);
  }
}
