import { CreationAttributes, Op } from "sequelize";
import BaseRepository from "./base.repository";
import Room from "../db/models/room.model";

export class RoomRepository extends BaseRepository<Room> {
  constructor() {
    super(Room);
  }

  async findByRoomCategoryIdAndDate(
    roomsCategoryId: number,
    currentDate: Date
  ) {
    return await this.model.findOne({
      where: {
        roomsCategoryId,
        dateOfAvailability: currentDate,
        deletedAt: null,
      },
    });
  }

  async bulkCreate(rooms: CreationAttributes<Room>[]) {
    return await this.model.bulkCreate(rooms);
  }

  async findLatestDateByRoomsCategoryId(
    roomsCategoryId: number
  ): Promise<Date | null> {
    const result = await this.model.findOne({
      where: {
        roomsCategoryId,
        deletedAt: null,
      },
      attributes: ["dateOfAvailability"],
      order: [["dateOfAvailability", "DESC"]],
    });

    return result ? result.dateOfAvailability : null;
  }

  async findLatestDatesForAllCategories(): Promise<
    Array<{ roomsCategoryId: number; latestDate: Date }>
  > {
    const results = await this.model.findAll({
      where: {
        deletedAt: null,
      },
      attributes: [
        "roomsCategoryId",
        [
          this.model.sequelize!.fn(
            "MAX",
            this.model.sequelize!.col("date_of_availability")
          ),
          "latestDate",
        ],
      ],
      group: ["roomsCategoryId"],
      raw: true,
    });

    return results.map((result: any) => ({
      roomsCategoryId: result.roomsCategoryId,
      latestDate: new Date(result.latestDate),
    }));
  }

  async findByRoomCategoryIdAndDateRange(
    roomsCategoryId: number,
    checkInDate: Date,
    checkOutDate: Date
  ) {
    return await this.model.findAll({
      where: {
        roomsCategoryId,
        bookingId: null,
        dateOfAvailability: {
        [Op.between]: [checkInDate, checkOutDate],
        },
      },
    });
  }
}
