import { Component, OnInit, OnChanges, SimpleChange, Input } from '@angular/core';
import { ActivatedRoute, Router, Params } from '@angular/router';

import { Alarm } from '../shared/alarm.model';
import { AlarmService } from '../shared/alarm.service';


@Component({
  selector: 'alarms',
  templateUrl: 'alarms.component.html',
  styleUrls: [ 'alarms.component.css' ],
  providers: [AlarmService]
})

export class AlarmsComponent implements OnInit, OnChanges {
    title = 'Alarms';
    alarms: Alarm[];
    selectedAlarm: Alarm;

    @Input() deviceId: string;

    constructor(private alarmService: AlarmService,
                private router: Router,
                private route: ActivatedRoute,
    ) {}

    getAlarms(): void {
        if (this.deviceId) {
            this.alarmService.getAlarms(this.deviceId).then(alarms => this.alarms = alarms).catch(() => console.log("catch"));
        }
    }

    ngOnInit(): void {
        this.route.params
            .switchMap((params: Params) =>  this.deviceId = params['deviceId']);
        this.route.params
            .subscribe(() => this.getAlarms());
    }

    ngOnChanges(): void {
        this.route.params
            .switchMap((params: Params) =>  this.deviceId = params['deviceId']);
        this.route.params
            .subscribe(() => this.getAlarms());
    }
    onSelect(alarm: Alarm): void {
        this.selectedAlarm = alarm;
    }
}
