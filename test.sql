CREATE TABLE game (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	isfootball integer not null default 0,
	teama varchar not null,
	teamb varchar not null,
	scorea integer null,
	scoreb integer null,
	oddsa real not null,
	oddsb real not null,
	concede real not null,
	scoresum real not null,
	timestarted varchar not null,
	timecreated varchar not null,
	poolwin real null,
	poollose real null,
	pooleven real null,
	poolodd real null,
	poollarge real null,
	poolsmall real null,
	poolsum real null,
	isover integer null
);

CREATE TABLE user (
	id INTEGER PRIMARY KEY AUTOINCREMENT, 
	username varchar unique, 			  
	password varchar not null,
	fundpassword varchar not null,			  	
	email varchar not null unique,
	birth varchar not null,				  
	btcaddress varchar not null unique,
	balance real null default 0,
	profit real null default 0,
	alltimebet real null default 0,
	lastip varchar null,
	referral INTEGER null
);

CREATE TABLE deposit (
	uid INTEGER not null,
	amount real not null,
	time varchar not null
);

CREATE TABLE withdraw (
	uid varchar not null,
	amount real not null,
	address varchar not null,
	time varchar not null
);

CREATE TABLE alluserbet (
	uid INTEGER not null,
	gid INTEGER not null,
	type varchar not null,
	betamount real not null,
	txhash varchar PRIMARY KEY,
	bettime varchar not null,
	profit real not null default 0,
	txhashwin varchar not null
);

CREATE TABLE currentbet (
	txhash varchar PRIMARY KEY,
	gid INTEGER not null,
	uid INTEGER not null,
	type varchar not null,
	bet real not null
);

CREATE TABLE historybet (
	uid INTEGER not null,
	gid INTEGER not null,
	type varchar not null,
	betamount real not null,
	txhash varchar PRIMARY KEY,
	bettime real not null,
	profit real not null,
	txhashwin varchar not null
);