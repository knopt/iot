import { Component, OnInit, Input } from '@angular/core';
import { ActivatedRoute, Router, Params } from '@angular/router';

import { Alarm } from '../shared/alarm.model';
import { AlarmService } from '../shared/alarm.service';


@Component({
  selector: 'stats',
  templateUrl: 'stats.component.html',
  styleUrls: [ 'stats.component.css' ],
  providers: [AlarmService]
})

export class StatsComponent implements OnInit {
    title = 'Statistics';

    @Input()
    deviceId: string

    constructor(private alarmService: AlarmService,
                private router: Router,
                private route: ActivatedRoute,
    ) {}


    ngOnInit(): void {
        this.route.params
            .switchMap((params: Params) => this.deviceId = params['deviceId']);
    }
    onSelect(alarm: Alarm): void {
        return;
    }
}
