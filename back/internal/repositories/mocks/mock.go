package mocks

import "loyalty/internal/core/domain"

var MockUsers = []domain.User{
	{
		ID:       0,
		IdPoints: 1,
	},
}

var MockWriterPoints = []domain.Points{
	{
		ID:     0,
		Points: 0,
	},
}

var MockReadPoints = []domain.Points{
	{
		ID:     0,
		Points: 0,
	},
}

var MockProducts = []domain.Product{
	{
		Id:    0,
		Name:  "ball",
		Price: 10_000,
		Image: "https://laika.com.co/cdn-cgi/image/onerror=redirect,format=auto,fit=scale-down,width=600,quality=80/https://laikapp.s3.amazonaws.com/dev_images_products/11145_169454_Argos___Juguete_Surtidos_Mediano_1657722486_500x500.jpg",
	}, {
		Id:    1,
		Name:  "snaks pot",
		Price: 20_000,
		Image: "https://laika.com.co/cdn-cgi/image/fit=scale-down,width=250,format=auto,quality=80,onerror=redirect/https://laikapp.s3.amazonaws.com/images_products/1_4328_SUPER_ORANGE_402X385.jpg",
	}, {
		Id:    2,
		Name:  "rope",
		Price: 30_000,
		Image: "https://laika.com.co/cdn-cgi/image/fit=scale-down,width=250,format=auto,quality=80,onerror=redirect/https://laikapp.s3.amazonaws.com/dev_images_products/74104_Beco_Pets___Aro_De_C____amo_1638208176_0_300x300.png",
	},
	{
		Id:    3,
		Name:  "pillow",
		Price: 40_000,
		Image: "https://laika.com.co/cdn-cgi/image/fit=scale-down,width=250,format=auto,quality=80,onerror=redirect/https://laikapp.s3.amazonaws.com/dev_images_products/23370_136193_Indupet_Peluche_Galleta_1644869825_500x500.jpg",
	},
	{
		Id:    4,
		Name:  "buffalo",
		Price: 50_000,
		Image: "https://laika.com.co/cdn-cgi/image/fit=scale-down,width=250,format=auto,quality=80,onerror=redirect/https://laikapp.s3.amazonaws.com/dev_images_products/15675_168422_Funimals___B__falo_Con_Sonido__1657228119_500x500.png",
	},
	{
		Id:    5,
		Name:  "pineapple",
		Price: 60_000,
		Image: "https://laika.com.co/cdn-cgi/image/fit=scale-down,width=250,format=auto,quality=80,onerror=redirect/https://laikapp.s3.amazonaws.com/dev_images_products/23378_136460_Indupet___Peluche_Pi__a__1644876493_500x500.jpg",
	},
	{
		Id:    6,
		Name:  "koala",
		Price: 70_000,
		Image: "https://laika.com.co/cdn-cgi/image/fit=scale-down,width=250,format=auto,quality=80,onerror=redirect/https://laikapp.s3.amazonaws.com/dev_images_products/67207_kong_peluche_snuzzies_koala__1622842429_0_1000x1000.jpg",
	},
	{
		Id:    7,
		Name:  "steak",
		Price: 80_000,
		Image: "https://laika.com.co/cdn-cgi/image/fit=scale-down,width=250,format=auto,quality=80,onerror=redirect/https://laikapp.s3.amazonaws.com/dev_images_products/23041_136031_Indupet___Peluche_Hueso_Osobuco__1644866425_500x500.jpg",
	},
}
