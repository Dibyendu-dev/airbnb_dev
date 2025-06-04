export interface NotificationDto {
    to: string;
    subject:string;
    templete:string;
    params:Record<string,any>;
}