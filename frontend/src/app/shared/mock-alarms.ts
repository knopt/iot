import { Alarm } from './alarm';

export const ALARMS: Alarm[] = [
    {
       id: "id1",
       alarmTime: new Date(86400000),
       createdAt: new Date(26400000),
       description: "description1",
       deviceId: "deviceid1",
       repeatWeekly: false,
   },
   {
      id: "id2",
      alarmTime: new Date(89400000),
      createdAt: new Date(29400000),
      description: "description2",
      deviceId: "deviceid2",
      repeatWeekly: true,
  }
];
