import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';

import 'rxjs/add/operator/toPromise';

import { Statistic } from './statistic.model';

@Injectable()
export class StatisticService {

    constructor(private http: Http) { }

    getStatistics(deviceID: string, type: string): Promise<Statistic[]> {
        return this.http.get('http://localhost:8080/statistics/get/device/' + deviceID + '/date/to//type/' + type)
           .toPromise()
           .then(response => response.json() as Statistic[])
           .catch(this.handleError);
    }

    getTypes(deviceID: string): Promise<string[]> {
        return this.http.get('http://localhost:8080/statistics/get/types/device/' + deviceID)
           .toPromise()
           .then(response => response.json() as string[])
           .catch(this.handleError);
    }

    private handleError(error: any): Promise<any> {
      console.error('An error occurred', error); // for demo purposes only
      return Promise.reject(error.message || error);
    }
}
