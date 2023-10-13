/*==============================================================*/
/* DBMS name:      MySQL 5.0                                    */
/* Created on:     2023/10/13 11:56:50                          */
/*==============================================================*/

/*==============================================================*/
/* Table: adoption                                              */
/*==============================================================*/
create table adoption
(
   adoption_id          int not null auto_increment,
   pet_id               int not null,
   uid                  int not null,
   is_reviewed          bool not null,
   is_approved          bool not null,
   primary key (adoption_id)
);

/*==============================================================*/
/* Table: blacklist                                             */
/*==============================================================*/
create table blacklist
(
   blacklist_id         int not null auto_increment,
   uid                  int,
   reason               varchar(1024) not null,
   primary key (blacklist_id)
);

/*==============================================================*/
/* Table: category                                              */
/*==============================================================*/
create table category
(
   category_id          int not null auto_increment,
   species              varchar(32) not null,
   color                varchar(32),
   gender               varchar(32),
   primary key (category_id)
);

/*==============================================================*/
/* Table: contact                                               */
/*==============================================================*/
create table contact
(
   uid                  int not null,
   kind                 varchar(32) not null,
   content              varchar(32) not null
);

/*==============================================================*/
/* Table: location                                              */
/*==============================================================*/
create table location
(
   location_id          int not null auto_increment,
   province             varchar(32) not null,
   city                 varchar(32),
   district             varchar(32) not null,
   primary key (location_id)
);

/*==============================================================*/
/* Table: pet                                                   */
/*==============================================================*/
create table pet
(
   pet_id               int not null auto_increment,
   uid                  int not null,
   category_id          int not null,
   nickname             varchar(32) not null,
   description          varchar(1024) not null,
   birthday             timestamp,
   is_adopted           bool not null,
   publish_date         timestamp not null,
   primary key (pet_id)
);

/*==============================================================*/
/* Index: is_adopted_index                                      */
/*==============================================================*/
create index is_adopted_index on pet
(
   is_adopted
);

/*==============================================================*/
/* Table: role                                                  */
/*==============================================================*/
create table role
(
   role_id              int not null auto_increment,
   rolename             varchar(32) not null unique,
   primary key (role_id)
);

/*==============================================================*/
/* Table: user                                                  */
/*==============================================================*/
create table user
(
   uid                  int not null auto_increment,
   location_id          int not null,
   username             varchar(32) not null unique,
   pword                varchar(1024) not null,
   gender               varchar(32) not null,
   birthday             timestamp not null,
   primary key (uid)
);

/*==============================================================*/
/* Index: username_index                                        */
/*==============================================================*/
create index username_index on user
(
   username
);

/*==============================================================*/
/* Table: user_role                                             */
/*==============================================================*/
create table user_role
(
   uid                  int not null,
   role_id              int not null,
   primary key (uid, role_id)
);

alter table adoption add constraint FK_adopt_pet foreign key (pet_id)
      references pet (pet_id) on delete restrict on update restrict;

alter table adoption add constraint FK_user_adopt foreign key (uid)
      references user (uid) on delete restrict on update restrict;

alter table blacklist add constraint FK_user_blacklist foreign key (uid)
      references user (uid) on delete restrict on update restrict;

alter table contact add constraint FK_user_contact foreign key (uid)
      references user (uid) on delete restrict on update restrict;

alter table pet add constraint FK_category_pet foreign key (category_id)
      references category (category_id) on delete restrict on update restrict;

alter table pet add constraint FK_publish_pet foreign key (uid)
      references user (uid) on delete restrict on update restrict;

alter table user add constraint FK_location_user foreign key (location_id)
      references location (location_id) on delete restrict on update restrict;

alter table user_role add constraint FK_user_role foreign key (role_id)
      references role (role_id) on delete restrict on update restrict;

alter table user_role add constraint FK_user_role2 foreign key (uid)
      references user (uid) on delete restrict on update restrict;
