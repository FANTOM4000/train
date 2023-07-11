//แสดง ชื่อ,เงินในบัญชี,อายุ ของผู้ใช้ทั้งหมกด้วย for in
//https://www.w3schools.com/js/js_loop_forin.asp
let user = [
    {
      "_id": "64ad59764f7ce5e022a8e67a",
      "index": 0,
      "guid": "126c9fdc-f5f1-4f3e-8b41-9389b75e6db0",
      "isActive": false,
      "balance": "$2,113.36",
      "picture": "http://placehold.it/32x32",
      "age": 20,
      "eyeColor": "blue",
      "name": "Stout Anthony",
      "gender": "male",
      "company": "OPTICOM",
      "email": "stoutanthony@opticom.com",
      "phone": "+1 (846) 594-2669",
      "address": "162 Amherst Street, Catherine, Wyoming, 843",
      "about": "Aute fugiat aliqua adipisicing anim anim deserunt minim eiusmod fugiat in amet ullamco tempor. Sunt pariatur aute nisi incididunt. Commodo laboris nostrud sit elit pariatur et excepteur dolore laborum culpa in adipisicing ea excepteur.\r\n",
      "registered": "2016-12-28T12:12:45 -07:00",
      "latitude": -84.902685,
      "longitude": 21.131895,
      "tags": [
        "fugiat",
        "reprehenderit",
        "cillum",
        "labore",
        "est",
        "laborum",
        "aute"
      ],
      "friends": [
        {
          "id": 0,
          "name": "Herman Hayes"
        },
        {
          "id": 1,
          "name": "Jaime Molina"
        },
        {
          "id": 2,
          "name": "Gilliam Webb"
        }
      ],
      "greeting": "Hello, Stout Anthony! You have 5 unread messages.",
      "favoriteFruit": "banana"
    },
    {
      "_id": "64ad5976297817fd27a18451",
      "index": 1,
      "guid": "14daa391-e8b8-4b4a-991f-5fee1801c3b5",
      "isActive": true,
      "balance": "$2,058.16",
      "picture": "http://placehold.it/32x32",
      "age": 27,
      "eyeColor": "blue",
      "name": "Vaughn Juarez",
      "gender": "male",
      "company": "EQUITOX",
      "email": "vaughnjuarez@equitox.com",
      "phone": "+1 (922) 489-2325",
      "address": "511 Clay Street, Dellview, North Dakota, 7252",
      "about": "Ad qui elit est voluptate aliquip in dolore. Non nostrud irure laborum irure tempor exercitation. Reprehenderit in officia aute eu.\r\n",
      "registered": "2014-05-20T04:57:33 -07:00",
      "latitude": -47.870408,
      "longitude": 87.767182,
      "tags": [
        "qui",
        "consectetur",
        "consectetur",
        "esse",
        "velit",
        "do",
        "duis"
      ],
      "friends": [
        {
          "id": 0,
          "name": "Malone Pacheco"
        },
        {
          "id": 1,
          "name": "Head Fletcher"
        },
        {
          "id": 2,
          "name": "Cheryl Wall"
        }
      ],
      "greeting": "Hello, Vaughn Juarez! You have 10 unread messages.",
      "favoriteFruit": "apple"
    },
    {
      "_id": "64ad59765ce2b1625cb20342",
      "index": 2,
      "guid": "bdd47567-6307-4a61-844e-d73d0624aacc",
      "isActive": false,
      "balance": "$1,336.51",
      "picture": "http://placehold.it/32x32",
      "age": 23,
      "eyeColor": "green",
      "name": "Noble Kerr",
      "gender": "male",
      "company": "ZILLANET",
      "email": "noblekerr@zillanet.com",
      "phone": "+1 (963) 599-2685",
      "address": "892 Bergen Avenue, Tilleda, Maine, 1073",
      "about": "Incididunt et culpa ullamco est esse cillum enim velit. Culpa magna mollit elit aliquip sit officia et magna. In non ad velit qui qui. Proident est nulla eiusmod aliqua tempor nulla Lorem irure ad anim. Esse et deserunt enim elit aute minim sunt magna laborum cupidatat sint cupidatat reprehenderit sit. Irure quis id Lorem eiusmod. Incididunt incididunt excepteur tempor magna sint tempor in.\r\n",
      "registered": "2016-02-16T01:08:44 -06:00",
      "latitude": 40.350037,
      "longitude": 136.466777,
      "tags": [
        "deserunt",
        "fugiat",
        "officia",
        "anim",
        "ea",
        "labore",
        "officia"
      ],
      "friends": [
        {
          "id": 0,
          "name": "Murphy Ward"
        },
        {
          "id": 1,
          "name": "Maxine Harris"
        },
        {
          "id": 2,
          "name": "Charles Abbott"
        }
      ],
      "greeting": "Hello, Noble Kerr! You have 2 unread messages.",
      "favoriteFruit": "strawberry"
    },
    {
      "_id": "64ad59768852ce8fba0ec300",
      "index": 3,
      "guid": "55a26122-e107-4d58-b0b1-9fa344e7be54",
      "isActive": true,
      "balance": "$3,698.24",
      "picture": "http://placehold.it/32x32",
      "age": 22,
      "eyeColor": "green",
      "name": "Jeri Jarvis",
      "gender": "female",
      "company": "QUIZMO",
      "email": "jerijarvis@quizmo.com",
      "phone": "+1 (917) 424-2842",
      "address": "689 Croton Loop, Kiskimere, Louisiana, 5058",
      "about": "Veniam deserunt nostrud occaecat exercitation labore. Irure sunt fugiat cupidatat esse anim qui dolore sint magna veniam reprehenderit dolor proident. Aliqua cupidatat mollit qui laboris voluptate ullamco nostrud aliqua. Nostrud aliqua enim qui proident.\r\n",
      "registered": "2017-03-19T07:56:27 -07:00",
      "latitude": 81.742196,
      "longitude": -41.885602,
      "tags": [
        "labore",
        "consectetur",
        "in",
        "non",
        "sunt",
        "sint",
        "enim"
      ],
      "friends": [
        {
          "id": 0,
          "name": "Joan Allen"
        },
        {
          "id": 1,
          "name": "Roth Guthrie"
        },
        {
          "id": 2,
          "name": "Linda Watkins"
        }
      ],
      "greeting": "Hello, Jeri Jarvis! You have 8 unread messages.",
      "favoriteFruit": "strawberry"
    },
    {
      "_id": "64ad59769dac1247a3da2eb3",
      "index": 4,
      "guid": "973baa8e-596e-486f-86b2-ff8c185e3ed7",
      "isActive": true,
      "balance": "$2,355.08",
      "picture": "http://placehold.it/32x32",
      "age": 22,
      "eyeColor": "brown",
      "name": "Tricia Merrill",
      "gender": "female",
      "company": "GLOBOIL",
      "email": "triciamerrill@globoil.com",
      "phone": "+1 (950) 489-2526",
      "address": "529 Mill Road, Bergoo, Palau, 2456",
      "about": "Duis ex duis sunt nisi non ad culpa mollit culpa sunt sunt. Qui et laborum magna nulla sint. Exercitation elit ea reprehenderit ipsum est sunt proident commodo magna commodo. Ullamco tempor laboris sunt cupidatat amet adipisicing ad est incididunt ipsum elit deserunt ullamco. Ex qui amet officia ea magna irure. Commodo pariatur est laboris proident culpa esse.\r\n",
      "registered": "2017-12-12T04:14:50 -07:00",
      "latitude": -51.680439,
      "longitude": -161.161868,
      "tags": [
        "dolore",
        "duis",
        "culpa",
        "excepteur",
        "consequat",
        "ipsum",
        "aliquip"
      ],
      "friends": [
        {
          "id": 0,
          "name": "Lester Hahn"
        },
        {
          "id": 1,
          "name": "Kerry Bonner"
        },
        {
          "id": 2,
          "name": "Charlotte Pope"
        }
      ],
      "greeting": "Hello, Tricia Merrill! You have 7 unread messages.",
      "favoriteFruit": "banana"
    }
]