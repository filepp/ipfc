
DELETE FROM  miner where id != '';
DELETE FROM  file where cid != '';
DELETE FROM  miner_file where id != '';


INSERT INTO miner (`id`, `role`, `state`) VALUES ('12D3KooWP7We1ubNk7U7iQNUihZeDgtcVy8fRQPhVBQs6pjHawm4', 0, 0);
INSERT INTO miner (`id`, `role`, `state`) VALUES ('12D3KooWRee6wn2LtpN1WYsj57avMFDYWbJ2479JnvQG5pm6RztS', 1, 0);