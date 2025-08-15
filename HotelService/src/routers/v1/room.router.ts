import express from "express";
import {
  getAvailableRoomsHandler,
  updateBookingIdToRoomsHandler,
} from "../../controllers/room.controller";
import { validateRequestBody } from "../../validators";
import {
  getAvailableRoomsSchema,
  updateBookingIdToRoomsSchema,
} from "../../validators/room.validator";

const roomRouter = express.Router();

roomRouter.get(
  "/available",
  validateRequestBody(getAvailableRoomsSchema),
  getAvailableRoomsHandler
);
roomRouter.post(
  "/update-booking-id",
  validateRequestBody(updateBookingIdToRoomsSchema),
  updateBookingIdToRoomsHandler
);

roomRouter.get("/health", (req, res) => {
  res.status(200).send("OK");
});

export default roomRouter;
