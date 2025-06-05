import { Job, Worker } from "bullmq";
import { MAILER_QUEUE } from "../queues/mailer.queue";
import { getRedisConnObject } from "../config/redis.config";
import { MAILER_PAYLOAD } from "../producers/email.producers";
import { NotificationDto } from "../dto/notification.dto";
import { renderMailTemplate } from "../templates/template.handler";
import { sendEmail } from "../services/mailer.service";
import logger from "../config/logger.config";


export const setupMailerWorker = ()=>{
    const emailConsumer = new Worker<NotificationDto>(
    MAILER_QUEUE,
    async (job:Job)=>{
        if(job.name !== MAILER_PAYLOAD){
            throw new Error("Invalid job name ")
        }
        const payload =job.data
        console.log(`processing email for: ${JSON.stringify(payload)}`)
        const emailContnt = await renderMailTemplate(payload.templateId,payload.params);
        await sendEmail(payload.to,payload.subject,emailContnt);
        logger.info(`email send to ${payload.to} with subject ${payload.subject}`)
    },
    {
        connection:getRedisConnObject()
    }
)

emailConsumer.on("failed",()=>{
    console.error("email processing failed")
});
emailConsumer.on("completed",()=>{
    console.log("email processing completed successfully");
})
}

