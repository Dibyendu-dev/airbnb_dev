import { serverConfig } from "../config";
import logger from "../config/logger.config";
import transporter from "../config/mailer.config";
import { InternalServerError } from "../utils/errors/app.error";

export async function sendEmail(to:string,subject:string,body:string){
    try {
        await transporter.sendMail({
            from:serverConfig.MAIL_USER,
            to,
            subject,
            html:body,
        })
        logger.info(`email send to ${to} wtth subject ${subject}`)
    } catch (error) {
        throw new InternalServerError(`failed to send email`)
    }
}