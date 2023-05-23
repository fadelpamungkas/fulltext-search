-- -------------------------------------------------------------
-- TablePlus 5.3.5(494)
--
-- https://tableplus.com/
--
-- Database: job
-- Generation Time: 2023-04-16 01:56:48.2810
-- -------------------------------------------------------------


-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS job_id_seq;

-- Table Definition
CREATE TABLE "public"."job" (
    "id" int4 NOT NULL DEFAULT nextval('job_id_seq'::regclass),
    "title" text NOT NULL,
    PRIMARY KEY ("id")
);

INSERT INTO "public"."job" ("id", "title") VALUES
(1, 'Software Engineer'),
(2, 'Database Administrator'),
(4, 'Systems Administrator'),
(5, 'Mobile App Developer'),
(6, 'Business Consultant'),
(7, 'Graphic Designer'),
(8, 'Systems Administrator'),
(9, 'Machine Maintenance Technician'),
(10, 'Building Maintenance Technician');
