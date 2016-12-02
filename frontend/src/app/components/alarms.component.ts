import { Component, OnInit, Input } from '@angular/core';
import { ActivatedRoute, Router, Params } from '@angular/router';

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

    @Input()
    deviceID: string

    constructor(private alarmService: AlarmService,
                private router: Router,
                private route: ActivatedRoute,
    ) {}

    getAlarms(): void {
        this.alarmService.getAlarms(this.deviceID).then(alarms => this.alarms = alarms);
    }
    ngOnInit(): void {
        this.route.params
            .switchMap((params: Params) => this.deviceID = params['deviceID'])
            .subscribe(() => this.getAlarms())
    }
    onSelect(alarm: Alarm): void {
        this.selectedAlarm = alarm
    }
}
