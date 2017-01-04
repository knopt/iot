import { NgModule }       from '@angular/core';
import { BrowserModule }  from '@angular/platform-browser';
import { FormsModule }    from '@angular/forms';
import { HttpModule }    from '@angular/http';

import { AppComponent }         from './app.component';
import { DashboardComponent }   from './components/dashboard.component';
import { DevicesComponent }     from './components/devices.component';
import { HeroDetailComponent }  from './components/hero-detail.component';
import { HeroesComponent }      from './components/heroes.component';
import { AlarmsComponent }      from './components/alarms.component';
import { StatsComponent }       from './components/stats.component';
import { AlarmService }         from './shared/alarm.service';
import { DeviceService }        from './shared/device.service';
import { HeroService }          from './shared/hero.service';
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
    HeroDetailComponent,
    HeroesComponent,
    AlarmsComponent,
    DevicesComponent,
    StatsComponent,
  ],
  providers: [ HeroService, AlarmService, DeviceService ],
  bootstrap: [ AppComponent ]
})
export class AppModule { }
