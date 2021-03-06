import { NgModule }       from '@angular/core';
import { BrowserModule }  from '@angular/platform-browser';
import { FormsModule }    from '@angular/forms';
import { HttpModule }    from '@angular/http';

import { AppComponent }         from './app.component';
import { DashboardComponent }   from './components/dashboard.component';
import { DevicesComponent }     from './components/devices.component';
import { AlarmsComponent }      from './components/alarms.component';
import { AlarmFormComponent }   from './components/alarm-form.component';
import { GoogleChartComponent } from './components/googleChart.component';
import { StatisticsComponent }  from './components/statistics.component';
import { AlarmService }         from './shared/alarm.service';
import { DeviceService }        from './shared/device.service';
import { StatisticService }     from './shared/statistic.service';
import { AppRoutingModule }     from './app-routing.module';

@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    AppRoutingModule
  ],
  declarations: [
    AlarmFormComponent,
    AlarmsComponent,
    AppComponent,
    DashboardComponent,
    DevicesComponent,
    GoogleChartComponent,
    StatisticsComponent,
  ],
  providers: [ AlarmService, DeviceService, StatisticService ],
  bootstrap: [ AppComponent ]
})
export class AppModule { }
