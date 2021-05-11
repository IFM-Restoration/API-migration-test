/* down script will revert all changes descending from new DB to 1.up.sql to old DB. You need to add drop to each created table */
DROP TABLE IF EXISTS WORK_ORDER;
DROP TABLE IF EXISTS CLIENT;
DROP TABLE IF EXISTS WORK_ORDER_ACCEPTANCE_STATUS;
DROP TABLE IF EXISTS WORK_ORDER_LOCATION;