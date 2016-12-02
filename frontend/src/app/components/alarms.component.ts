import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { Alarm } from '../shared/alarm.model';
import { AlarmService } from '../shared/alarm.service';


@Component({
  selector: 'alarms',
  templateUrl: 'alarms.component.html',
  styleUrls: [ 'alarms.component.css' ],
  providers: [AlarmService]
})

export class AlarmsComponent implements OnInit {
    title = 'Alarms';
    alarms: Alarm[];
    selectedAlarm: Alarm;
    httpAlarm: Alarm;

    constructor(private alarmService: AlarmService,
                private router: Router
    ) {}

    getAlarms(): void {
        this.alarmService.getAlarms().then(alarms => this.alarms = alarms);
    }
    getHttpAlarm(): void {
        console.log("get alarm!!!");
        this.alarmService.getAlarmHttp().then(alarm => {this.httpAlarm = alarm;
            console.log(alarm);
        });
    }
    ngOnInit(): void {
        this.getHttpAlarm();
        this.getAlarms();
    }
    onSelect(alarm: Alarm): void {
        this.selectedAlarm = alarm
    }
}
