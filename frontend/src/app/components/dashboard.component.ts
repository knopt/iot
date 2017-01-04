import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router, Params } from '@angular/router';

import { Device } from '../shared/device.model';

@Component({
  selector: 'my-dashboard',
  templateUrl: 'dashboard.component.html',
  styleUrls: ['dashboard.component.css']
})

export class DashboardComponent {
    selectedTab: string;
    selectedDevice: Device;

    constructor(private router: Router,
                private route: ActivatedRoute
    ) {}

    ngOnInit(): void {
        this.selectedTab = "alarms";
    }
    onSelect(tab: string): void {
        this.selectedTab = tab;
    }
    onDeviceSelected(device: Device) {
        this.selectedDevice = device;
    }
    getSelectedDeviceId() {
        if (this.selectedDevice) {
            return this.selectedDevice.id;
        }
        return "";
    }
}
