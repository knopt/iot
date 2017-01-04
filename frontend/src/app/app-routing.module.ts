import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AlarmsComponent }      from './components/alarms.component';
import { DashboardComponent }   from './components/dashboard.component';
import { DevicesComponent }     from './components/devices.component';
import { StatsComponent }       from './components/stats.component';

const routes: Routes = [
  { path: '', redirectTo: '/dashboard', pathMatch: 'full' },
  { path: 'dashboard',          component: DashboardComponent },
  { path: 'alarms/:deviceID',   component: AlarmsComponent },
  { path: 'devices',            component: DevicesComponent },
  { path: 'stats/:deviceID',    component: StatsComponent },
];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})

export class AppRoutingModule {}
