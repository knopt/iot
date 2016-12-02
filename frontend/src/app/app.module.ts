import { NgModule }       from '@angular/core';
import { BrowserModule }  from '@angular/platform-browser';
import { FormsModule }    from '@angular/forms';
import { HttpModule }    from '@angular/http';

import { AppComponent }         from './app.component';
import { DashboardComponent }   from './components/dashboard.component';
import { HeroDetailComponent }  from './components/hero-detail.component';
import { HeroesComponent }      from './components/heroes.component';
import { AlarmsComponent }      from './components/alarms.component';
import { HeroService }          from './shared/hero.service';
import { AlarmService }         from './shared/alarm.service';
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
    AlarmsComponent
  ],
  providers: [ HeroService, AlarmService ],
  bootstrap: [ AppComponent ]
})
export class AppModule { }
