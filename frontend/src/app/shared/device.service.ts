import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';

import 'rxjs/add/operator/toPromise';

import { Device } from './device.model';
import { Config } from '../app.config';

@Injectable()
export class DeviceService {
    private getDevicesUrl = Config.domainName + '/device/get/all'

    constructor(private http: Http) { }

    getDevices(): Promise<Device[]> {
        return this.http.get(this.getDevicesUrl)
           .toPromise()
           .then(response => response.json() as Device[])
           .catch(this.handleError);
    }

    private handleError(error: any): Promise<any> {
      console.error('An error occurred', error); // for demo purposes only
      return Promise.reject(error.message || error);
    }
}
