import { Router } from "express";

import {
  getSchedulerHandler,
  manualExtendAvailabilityHandler,
  startSchedulerHandler,
  stopSchedulerHandler,
} from "../../controllers/roomSchedular.controller";

const roomSchedulerRouter = Router();

roomSchedulerRouter.post("/start", startSchedulerHandler);
roomSchedulerRouter.post("/stop", stopSchedulerHandler);
roomSchedulerRouter.get("/status", getSchedulerHandler);
roomSchedulerRouter.post("extend", manualExtendAvailabilityHandler);

export default roomSchedulerRouter;