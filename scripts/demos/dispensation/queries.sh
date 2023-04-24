#!/usr/bin/env bash
rm -rf all.json pending.json completed.json
grided q dispensation records-by-name ar1 All>> all.json
grided q dispensation records-by-name ar1 Pending >> pending.json
grided q dispensation records-by-name ar1 Completed>> completed.json