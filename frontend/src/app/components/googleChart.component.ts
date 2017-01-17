import { Component, OnChanges, OnInit, Input } from '@angular/core';

declare var google:any;

@Component({
  selector: 'chart',
  template: `
  <div>
      <div (click)="onRefresh()">
          <h3> Refresh! </h3>
      </div>
      <div class="four wide column center aligned">
          <div id="chart_divEvolution" style="width: 900px; height: 500px;"></div>
      </div>
  </div>
`
})
export class GoogleChartComponent implements OnInit {
  private static googleLoaded:any;

  private options: any;
  private chart: any;
  private data: any;
  private chartData: any;
  private chartType: string;

  @Input('statsData')
  set statsData(data: any) {
      this.chartData = data;
      setTimeout(function() {
          this.drawGraph();
      }, 3000);
  }

  @Input('statsType')
  set statsType(type: string) {
      this.chartType = type;
  }

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

  onRefresh() {
      this.drawGraph();
  }

  drawGraph(){
    console.log("LineGraph Evolution...");

    this.data = this.createDataTable(this.chartData);

    this.options = {
      title: `${this.chartType}`,
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
