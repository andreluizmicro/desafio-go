-- transfer.users definition

CREATE TABLE `users` (
                         `id` varchar(100) NOT NULL,
                         `user_type_id` int(11) NOT NULL,
                         `name` varchar(200) NOT NULL,
                         `email` varchar(200) NOT NULL,
                         `cpf` varchar(14) NOT NULL,
                         `cnpj` varchar(18),
                         `password` varchar(100) NOT NULL,
                         `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `deleted` tinyint(4) NOT NULL DEFAULT '0',
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;