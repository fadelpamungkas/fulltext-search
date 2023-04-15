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
    "title_greek" text NOT NULL,
    "title_english" text,
    "keywords" text,
    PRIMARY KEY ("id")
);

INSERT INTO "public"."job" ("id", "title_greek", "title_english", "keywords") VALUES
(2, 'Διαχειριστής Βάσης Δεδομένων', '', 'Oracle, SQL, PL/SQL, Παρακολούθηση και βελτιστοποίηση επιδόσεων'),
(4, 'Διαχειριστής Συστημάτων', '', ''),
(5, 'Προγραμματιστής Κινητών Εφαρμογών', 'Mobile App Developer', 'iOS, Swift, Android, Kotlin, Flutter, Firebase, Mobile App Development'),
(6, 'Σύμβουλος Επιχειρήσεων', 'Business Consultant', ''),
(7, 'Σχεδιαστής Γραφικών', '', 'Adobe Creative Suite, Illustration, Branding, User Experience, Visual Design'),
(8, 'Διαχειριστής Συστημάτων', 'Systems Administrator', 'Linux, Virtualization, Networking, Security'),
(9, 'Συντηρητής Μηχανημάτων', 'Machine Maintenance Technician', 'Mechanical Maintenance, Troubleshooting, Repair, Preventive Maintenance, Equipment Calibration'),
(10, 'Συγγραφέας Περιεχομένου', '', ''),
(11, 'Συντηρητής Κτιρίων', '', 'Electrical Systems, Plumbing, HVAC, Carpentry, Painting');
