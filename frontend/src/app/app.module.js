"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
const core_1 = require("@angular/core");
const platform_browser_1 = require("@angular/platform-browser");
const forms_1 = require("@angular/forms");
const http_1 = require("@angular/http");
const app_component_1 = require("./app.component");
const dashboard_component_1 = require("./components/dashboard.component");
const devices_component_1 = require("./components/devices.component");
const alarms_component_1 = require("./components/alarms.component");
const googleChart_component_1 = require("./components/googleChart.component");
const statistics_component_1 = require("./components/statistics.component");
const alarm_service_1 = require("./shared/alarm.service");
const device_service_1 = require("./shared/device.service");
const statistic_service_1 = require("./shared/statistic.service");
const app_routing_module_1 = require("./app-routing.module");
let AppModule = class AppModule {
};
AppModule = __decorate([
    core_1.NgModule({
        imports: [
            platform_browser_1.BrowserModule,
            forms_1.FormsModule,
            http_1.HttpModule,
            app_routing_module_1.AppRoutingModule
        ],
        declarations: [
            alarms_component_1.AlarmsComponent,
            app_component_1.AppComponent,
            dashboard_component_1.DashboardComponent,
            devices_component_1.DevicesComponent,
            googleChart_component_1.GoogleChartComponent,
            statistics_component_1.StatisticsComponent,
        ],
        providers: [alarm_service_1.AlarmService, device_service_1.DeviceService, statistic_service_1.StatisticService],
        bootstrap: [app_component_1.AppComponent]
    }),
    __metadata("design:paramtypes", [])
], AppModule);
exports.AppModule = AppModule;
//# sourceMappingURL=app.module.js.map