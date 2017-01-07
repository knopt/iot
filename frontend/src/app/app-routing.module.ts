import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AlarmsComponent }      from './components/alarms.component';
import { DashboardComponent }   from './components/dashboard.component';
import { DevicesComponent }     from './components/devices.component';
import { StatisticsComponent }       from './components/statistics.component';

const routes: Routes = [
  { path: '', redirectTo: '/dashboard', pathMatch: 'full' },
  { path: 'dashboard',              component: DashboardComponent },
  { path: 'alarms/:deviceID',       component: AlarmsComponent },
  { path: 'devices',                component: DevicesComponent },
  { path: 'statistics/:deviceID',   component: StatisticsComponent },
];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})

export class AppRoutingModule {}
