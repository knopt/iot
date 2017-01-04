import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { Router } from '@angular/router';

import { Device } from '../shared/device.model';
import { DeviceService } from '../shared/device.service';


@Component({
  selector: 'devices',
  templateUrl: 'devices.component.html',
  styleUrls: [ 'devices.component.css' ],
  providers: [DeviceService]
})

export class DevicesComponent implements OnInit {
    title = 'Devices';
    devices: Device[];
    selectedDevice: Device;

    @Output() onDeviceSelected = new EventEmitter<Device>();

    constructor(private deviceService: DeviceService,
                private router: Router
    ) {}

    getDevices(): void {
        this.deviceService.getDevices().then(devices => this.devices = devices);
    }
    ngOnInit(): void {
        this.getDevices();
    }
    onSelect(device: Device): void {
        this.selectedDevice = device;
        this.onDeviceSelected.emit(device);
    }
    gotoAlarms(): void {
      this.router.navigate(['/alarms', this.selectedDevice.id]);
    }
}
