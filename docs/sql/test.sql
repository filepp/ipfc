
DELETE FROM  miner where id != '';
DELETE FROM  file where cid != '';
DELETE FROM  miner_file where id != '';
DELETE FROM  distribution_log where id > 0;
DELETE FROM  window_post where id > 0;
