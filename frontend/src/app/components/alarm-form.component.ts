import { Component, Input } from '@angular/core';

import { Alarm }    from '../shared/alarm.model';
import { AlarmService } from '../shared/alarm.service';

@Component({
  selector: 'alarm-form',
  templateUrl: 'alarm-form.component.html',
  styleUrls: [ 'alarm-form.component.css' ],
  providers: [AlarmService]
})
export class AlarmFormComponent {
    constructor(private alarmService: AlarmService) {}

    model = new Alarm();
    deviceId: string;

    @Input()
    set setDeviceId(deviceId: string) {
        this.deviceId = deviceId;
        if (this.model) {
            this.model.device_id = deviceId;
        }
    }

    onSubmit() {
        this.model.created_at = new Date();
        this.model.alarm_time = new Date(this.model.alarm_time)
        this.alarmService.postAlarm(this.model);
    }
}
