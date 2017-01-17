import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';

import 'rxjs/add/operator/toPromise';

import { Alarm } from './alarm.model';
import { Config } from '../app.config';

@Injectable()
export class AlarmService {

    constructor(private http: Http) { }

    getAlarms(deviceID: string): Promise<Alarm[]> {
        return this.http.get(Config.domainName + '/alarm/get/id/' + deviceID + '/all')
           .toPromise()
           .then(response => response.json() as Alarm[])
           .catch(this.handleError);
    }

    postAlarm(alarm: Alarm): Promise<any> {
        return this.http.post(Config.domainName + '/alarm/set', JSON.stringify(alarm))
            .toPromise()
            .catch(this.handleError);
    }

    private handleError(error: any): Promise<any> {
      console.error('An error occurred', error); // for demo purposes only
      return Promise.reject(error.message || error);
    }
}
