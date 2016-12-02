import { Alarm } from './alarm.model';

export const ALARMS: Alarm[] = [
    {
       id: "id1",
       alarm_time: new Date(86400000),
       created_at: new Date(26400000),
       description: "description1",
       device_id: "deviceid1",
       repeat_weekly: false,
   },
   {
      id: "id2",
      alarm_time: new Date(89400000),
      created_at: new Date(29400000),
      description: "description2",
      device_id: "deviceid2",
      repeat_weekly: true,
  }
];
