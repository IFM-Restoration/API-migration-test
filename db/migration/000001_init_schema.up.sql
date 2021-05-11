/* create new DB tables with relationships, will create 1.up.sql and more with every migrate command */
CREATE TABLE "WORK_ORDER" (
                              "id" bigserial PRIMARY KEY,
                              "client_id" int,
                              "resident_id" int,
                              "work_order_acceptance_status_id" int,
                              "clint_contant_approved_by_id" int DEFAULT 0,
                              "name" VARCHAR,
                              "uuid" VARCHAR,
                              "due_date" TIMESTAMP,
                              "is_canceled" boolean,
                              "created_datetime" TIMESTAMP,
                              "created_by" VARCHAR,
                              "last_modified_datetime" TIMESTAMP,
                              "last_modified_by" VARCHAR,
                              "is_deleted" boolean
);

CREATE TABLE "CLIENT" (
                          "id" bigserial PRIMARY KEY,
                          "name" VARCHAR,
                          "created_datetime" TIMESTAMP,
                          "created_by" VARCHAR,
                          "last_modified_datetime" TIMESTAMP,
                          "last_modified_by" VARCHAR,
                          "is_deleted" boolean
);

CREATE TABLE "WORK_ORDER_ACCEPTANCE_STATUS" (
                                                "id" bigserial PRIMARY KEY,
                                                "acceptance_status" VARCHAR,
                                                "created_datetime" TIMESTAMP,
                                                "created_by" VARCHAR,
                                                "last_modified_datetime" TIMESTAMP,
                                                "last_modified_by" VARCHAR,
                                                "is_deleted" boolean
);

CREATE TABLE "WORK_ORDER_LOCATION" (
                                       "id" bigserial PRIMARY KEY,
                                       "work_order_id" int,
                                       "market_id" int,
                                       "address_1" VARCHAR,
                                       "address_2" VARCHAR,
                                       "city" VARCHAR,
                                       "state" VARCHAR,
                                       "zip" VARCHAR,
                                       "access_information" VARCHAR,
                                       "created_datetime" TIMESTAMP,
                                       "created_by" VARCHAR,
                                       "last_modified_datetime" TIMESTAMP,
                                       "last_modified_by" VARCHAR,
                                       "is_deleted" boolean
);

ALTER TABLE "WORK_ORDER" ADD FOREIGN KEY ("client_id") REFERENCES "CLIENT" ("id");

ALTER TABLE "WORK_ORDER" ADD FOREIGN KEY ("work_order_acceptance_status_id") REFERENCES "WORK_ORDER_ACCEPTANCE_STATUS" ("id");

ALTER TABLE "WORK_ORDER_LOCATION" ADD FOREIGN KEY ("work_order_id") REFERENCES "WORK_ORDER" ("id");