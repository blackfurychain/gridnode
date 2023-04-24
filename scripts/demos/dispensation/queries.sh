#!/usr/bin/env bash
rm -rf all.json pending.json completed.json
gridnoded q dispensation records-by-name ar1 All>> all.json
gridnoded q dispensation records-by-name ar1 Pending >> pending.json
gridnoded q dispensation records-by-name ar1 Completed>> completed.json