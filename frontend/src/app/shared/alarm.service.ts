import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';

import 'rxjs/add/operator/toPromise';

import { Alarm } from './alarm.model';

@Injectable()
export class AlarmService {
    private getAlarmUrl = 'http://localhost:8080/alarm/get/id/5831fa07d0d41f1364a5eeb0'
    private getAlarmsUrl = 'http://localhost:8080/alarm/get/id/58362c0bd0d41f4f91240f29/all'

    constructor(private http: Http) { }

    getAlarms(deviceID: string): Promise<Alarm[]> {
        return this.http.get('http://localhost:8080/alarm/get/id/' + deviceID + '/all')
           .toPromise()
           .then(response => response.json() as Alarm[])
           .catch(this.handleError);
    }

    // getAlarm(id: string): Promise<Alarm> {
    //     return this.getAlarms()
    //         .then(alarms => alarms.find(alarm => alarm.id === id));
    // }

    private handleError(error: any): Promise<any> {
      console.error('An error occurred', error); // for demo purposes only
      return Promise.reject(error.message || error);
    }
}
