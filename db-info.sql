sudo apt-get update
sudo apt-get install mysql-server


mysql -u root -p
show databases;
create database demodb;
use demodb;

    CREATE TABLE `users` (
        `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
        `name` varchar(30) NOT NULL,
        `email` varchar(30) NOT NULL,
        `address` varchar(30) NOT NULL,
        PRIMARY KEY (`id`)
      ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;


