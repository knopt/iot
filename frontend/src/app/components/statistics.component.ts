import { Component, OnChanges, OnInit, Input } from '@angular/core';
import { ActivatedRoute, Router, Params } from '@angular/router';

import { Statistic } from '../shared/statistic.model';
import { StatisticService } from '../shared/statistic.service';


@Component({
  selector: 'statistics',
  templateUrl: 'statistics.component.html',
  styleUrls: [ 'statistics.component.css' ],
  providers: [StatisticService]
})

export class StatisticsComponent implements OnInit, OnChanges {
    title = 'Statistics';
    statistics: Statistic[];
    types: string[];
    selectedType: string;
    selectedStatistic: Statistic;

    @Input()
    deviceId: string

    constructor(private statisticService: StatisticService,
                private router: Router,
                private route: ActivatedRoute,
    ) {}

    getStatistics(): void {
        if (this.deviceId && this.selectedType) {
            this.statisticService.getStatistics(this.deviceId, this.selectedType)
                .then(statistics => this.statistics = statistics)
                .catch(() => console.log("catch"));
        }
    }

    getTypes(): void {
        if (this.deviceId) {
            this.statisticService.getTypes(this.deviceId)
                .then(types => this.types = types)
                .catch(() => console.log("catch"));
        }
    }

    ngOnInit(): void {
        this.route.params
            .switchMap((params: Params) =>  this.deviceId = params['deviceId']);
        this.route.params
            .subscribe(() => this.getStatistics());
        this.route.params
            .subscribe(() => this.getTypes());
    }

    onSelectStatistic(statistic: Statistic): void {
        this.selectedStatistic = statistic;
    }
    
    onSelectType(type: string): void {
        this.selectedType = type;
        this.getStatistics();
    }

    ngOnChanges(): void {
        this.route.params
            .switchMap((params: Params) =>  this.deviceId = params['deviceId']);
        this.route.params
            .subscribe(() => this.getStatistics());
        this.route.params
            .subscribe(() => this.getTypes());
    }
}
