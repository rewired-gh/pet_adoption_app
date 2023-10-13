set FOREIGN_KEY_CHECKS = 0;

drop trigger if exists review_adoption_action;
drop trigger if exists approve_adoption_action;

drop procedure if exists get_user_privilege_info;

drop view if exists pet_info;

drop table if exists adoption;
drop table if exists blacklist;
drop table if exists category;
drop table if exists contact;
drop table if exists `location`;
drop table if exists pet;
drop table if exists `role`;
drop table if exists user;
drop table if exists user_role;

set FOREIGN_KEY_CHECKS = 1;