import { NgModule }       from '@angular/core';
import { BrowserModule }  from '@angular/platform-browser';
import { FormsModule }    from '@angular/forms';
import { HttpModule }    from '@angular/http';

import { AppComponent }         from './app.component';
import { DashboardComponent }   from './components/dashboard.component';
import { DevicesComponent }     from './components/devices.component';
import { AlarmsComponent }      from './components/alarms.component';
import { StatsComponent }       from './components/stats.component';
import { AlarmService }         from './shared/alarm.service';
import { DeviceService }        from './shared/device.service';
import { AppRoutingModule }     from './app-routing.module';

@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    AppRoutingModule
  ],
  declarations: [
    AppComponent,
    DashboardComponent,
    AlarmsComponent,
    DevicesComponent,
    StatsComponent,
  ],
  providers: [ AlarmService, DeviceService ],
  bootstrap: [ AppComponent ]
})
export class AppModule { }
