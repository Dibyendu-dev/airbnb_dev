-- AlterTable
ALTER TABLE `booking` MODIFY `status` ENUM('pending', 'confirmed', 'cancelled') NOT NULL DEFAULT 'pending';
