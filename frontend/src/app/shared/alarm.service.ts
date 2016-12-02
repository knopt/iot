import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';

import 'rxjs/add/operator/toPromise';

import { Alarm } from './alarm.model';
import { ALARMS } from './mock-alarms';

@Injectable()
export class AlarmService {
    private getAlarmUrl = 'http://localhost:8080/alarm/get/id/5831fa07d0d41f1364a5eeb0'
    private getAlarmsUrl = 'http://localhost:8080/alarm/get/device/58362c0bd0d41f4f91240f29/all'

    constructor(private http: Http) { }

    getAlarms(): Promise<Alarm[]> {
        return this.http.get(this.getAlarmsUrl)
           .toPromise()
           .then(response => response.json() as Alarm[])
           .catch(this.handleError);
    }

    getAlarm(id: string): Promise<Alarm> {
        return this.getAlarms()
            .then(alarms => alarms.find(alarm => alarm.id === id));
    }

    getAlarmHttp(): Promise<Alarm> {
        return this.http.get(this.getAlarmUrl)
           .toPromise()
           .then(response => response.json() as Alarm)
           .catch(this.handleError);
    }

    private handleError(error: any): Promise<any> {
      console.error('An error occurred', error); // for demo purposes only
      return Promise.reject(error.message || error);
    }
}
