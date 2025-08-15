import express from 'express';
import { getAvailableRoomsHandler } from '../../controllers/room.controller';
import { validateRequestBody } from '../../validators';
import { getAvailableRoomsSchema } from '../../validators/room.validator';


const roomRouter = express.Router();

roomRouter.get("/available",validateRequestBody(getAvailableRoomsSchema),getAvailableRoomsHandler);

roomRouter.get("/health",(req,res)=>{
    res.status(200).send("OK")
})

export default roomRouter