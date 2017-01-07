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
const router_1 = require("@angular/router");
const statistic_service_1 = require("../shared/statistic.service");
let StatisticsComponent = class StatisticsComponent {
    constructor(statisticService, router, route) {
        this.statisticService = statisticService;
        this.router = router;
        this.route = route;
        this.title = 'Statistics';
    }
    getStatistics() {
        if (this.deviceId && this.selectedType) {
            this.statisticService.getStatistics(this.deviceId, this.selectedType)
                .then(statistics => this.statistics = statistics)
                .catch(() => console.log("catch"));
        }
    }
    getTypes() {
        if (this.deviceId) {
            this.statisticService.getTypes(this.deviceId)
                .then(types => this.types = types)
                .catch(() => console.log("catch"));
        }
    }
    ngOnInit() {
        this.route.params
            .switchMap((params) => this.deviceId = params['deviceId']);
        this.route.params
            .subscribe(() => this.getStatistics());
        this.route.params
            .subscribe(() => this.getTypes());
    }
    onSelectStatistic(statistic) {
        this.selectedStatistic = statistic;
    }
    onSelectType(type) {
        this.selectedType = type;
        this.getStatistics();
    }
    ngOnChanges() {
        this.route.params
            .switchMap((params) => this.deviceId = params['deviceId']);
        this.route.params
            .subscribe(() => this.getStatistics());
        this.route.params
            .subscribe(() => this.getTypes());
    }
};
__decorate([
    core_1.Input(),
    __metadata("design:type", String)
], StatisticsComponent.prototype, "deviceId", void 0);
StatisticsComponent = __decorate([
    core_1.Component({
        selector: 'statistics',
        templateUrl: 'statistics.component.html',
        styleUrls: ['statistics.component.css'],
        providers: [statistic_service_1.StatisticService]
    }),
    __metadata("design:paramtypes", [statistic_service_1.StatisticService,
        router_1.Router,
        router_1.ActivatedRoute])
], StatisticsComponent);
exports.StatisticsComponent = StatisticsComponent;
//# sourceMappingURL=statistics.component.js.map