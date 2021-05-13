package main

import "time"

type WORK_ORDER struct {
	ID                              int
	client_id                       int
	resident_id                     int
	work_order_acceptance_status_id int
	clint_contant_approved_by_id    int
	name                            string
	uuid                            string
	due_date                        time.Time
	is_canceled                     bool
	created_datetime                time.Time
	created_by                      string
	last_modified_datetime          time.Time
	last_modified_by                string
	is_deleted                      bool
}
